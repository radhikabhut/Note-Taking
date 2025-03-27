# Note-Tacking
his is a simple Go-based web application that allows users to:
Upload Markdown (.md) files.
Render Markdown files to HTML.
List all uploaded files.
The project uses Gin for handling HTTP requests and Goldmark for Markdown rendering.

# Features

File Upload: Upload .md files to the server.
Grammar Check: Checks uploaded Markdown files for grammar errors using an API.
Render Markdown: Converts Markdown files into HTML format.
List Files: Fetch all uploaded Markdown files.

 Installation & Setup

1️. Clone the Repository

git clone https://github.com/radhikabhut/Note-Taking.git
cd Note-Taking

2️. Initialize Go Modules

go mod init Note-Taking
go mod tidy

3️. Install Dependencies

go get github.com/gin-gonic/gin
go get github.com/yuin/goldmark

4️.  Run the Application
go run main.go

# Additional Notes

Ensure the uploads directory exists before running the application.
Only Markdown (.md) files are allowed for upload.