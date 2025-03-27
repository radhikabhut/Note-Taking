package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yuin/goldmark"
)

const grammarCheckAPI = "https://api.languagetool.org/v2/check"
const allowedMimeType = "text/plain; charset=utf-8"

type GrammarResponse struct {
	Matches []struct {
		Message string `json:"message"`
		Offset  int    `json:"offset"`
		Length  int    `json:"length"`
		Context struct {
			Text string `json:"text"`
		} `json:"context"`
	} `json:"matches"`
}

func UploadFile(c *gin.Context) {
	fileh, err := c.FormFile("file")
	if err != nil {
		slog.Error("error while getting file from req.", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid file format",
		})
		return
	}

	file, err := fileh.Open()
	if err != nil {
		slog.Error("error while opening file from req.", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid file format",
		})
		return
	}
	defer file.Close()

	slog.Info("File name received", "filename", fileh.Filename)

	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		slog.Error("error while reading file.", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid file format",
		})
		return
	}

	fileType := http.DetectContentType(buffer)
	slog.Info("Detected file type", "fileType", fileType)

	if fileType != allowedMimeType {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file type. Only Markdown (.md) files are allowed"})
		return
	}

	savePath := "F:\\workspace\\Note-Taking\\uploads\\" + fileh.Filename
	err = c.SaveUploadedFile(fileh, savePath)
	if err != nil {
		slog.Error("error while saving file.", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "cannot save file",
		})
		return
	}

	content, err := os.ReadFile(savePath)
	if err != nil {
		slog.Error("Error while reading saved file.", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Cannot read saved file",
		})
		return
	}
	grammarSuggestions, err := CheckGrammar(string(content))
	if err != nil {
		slog.Error("Error while checking grammar.", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Error checking grammar",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":             true,
		"message":             "File saved successfully",
		"grammar_suggestions": grammarSuggestions,
	})
}
func CheckGrammar(content string) ([]string, error) {
	// Prepare form values
	formData := fmt.Sprintf("text=%s&language=en-US", content)

	// Create a new POST request with form data
	resp, err := http.Post(grammarCheckAPI, "application/x-www-form-urlencoded", strings.NewReader(formData))
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Check if response is valid JSON
	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		body, _ := io.ReadAll(resp.Body)
		slog.Error("Invalid response from API", "body", string(body))
		return nil, fmt.Errorf("invalid response format")
	}

	// Read API response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Check if response status is OK
	if resp.StatusCode != http.StatusOK {
		slog.Error("Error from LanguageTool API", "status", resp.StatusCode, "body", string(body))
		return nil, fmt.Errorf("failed to get valid response from LanguageTool API")
	}

	// Parse API response
	var grammarResponse GrammarResponse
	err = json.Unmarshal(body, &grammarResponse)
	if err != nil {
		slog.Error("Failed to parse API response", "error", err.Error(), "response", string(body))
		return nil, fmt.Errorf("failed to parse API response: %v", err)
	}

	// Extract grammar suggestions
	var suggestions []string
	for _, match := range grammarResponse.Matches {
		suggestion := fmt.Sprintf("Error: %s | Context: %s", match.Message, match.Context.Text)
		suggestions = append(suggestions, suggestion)
	}
	return suggestions, nil
}

func RenderMarkdown(c *gin.Context) {
	filename := c.Query("file")
	if filename == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "File name is required",
		})
		return
	}
	// Read the file from the uploads folder
	filePath := fmt.Sprintf("F:\\workspace\\Note-Taking\\uploads\\%s", filename)
	content, err := os.ReadFile(filePath)
	if err != nil {
		slog.Error("Error reading file", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Unable to read the file",
		})
		return
	}
	var buf bytes.Buffer
	err = goldmark.Convert(content, &buf)
	if err != nil {
		slog.Error("Error converting markdown to HTML", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Error converting file to HTML",
		})
		return
	}

	// Return HTML response
	c.Data(http.StatusOK, "text/html; charset=utf-8", buf.Bytes())
}

func ListFiles(c *gin.Context) {
	uploadDir := "F:\\workspace\\Note-Taking\\uploads\\"

	files, err := os.ReadDir(uploadDir)
	if err != nil {
		slog.Error("Error while reading the directory", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Unable to list files",
		})
		return
	}

	var filelist []string
	for _, file := range files {
		if !file.IsDir() {
			filelist = append(filelist, file.Name())
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"files":   filelist,
	})
}
