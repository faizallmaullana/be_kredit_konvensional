package database

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DownloadDatabase(c *gin.Context) {
	// Read the content of main.go file
	content, err := ioutil.ReadFile("database.db")
	if err != nil {
		c.String(http.StatusInternalServerError, "Error reading file")
		return
	}

	// Set the response headers
	c.Header("Content-Disposition", "attachment; filename=database.db")
	c.Header("Content-Type", "application/octet-stream")

	// Write the file content to response body
	c.Writer.Write(content)
}
