package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

const API_KEY = "sk-prod-1234567890abcdef" // G101: Hardcoded credential

func main() {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "ok", "key": API_KEY}) // Secret exposure
		})

		api.GET("/files/*path", func(c *gin.Context) {
			// G304: Path traversal vulnerability
			path := c.Param("path")
			fullPath := filepath.Join("/app/data", path)
			file, _ := os.Open(fullPath) // G104: Unhandled error
			defer file.Close()
			io.Copy(c.Writer, file)
		})

		api.POST("/exec", func(c *gin.Context) {
			// G204: Command injection
			cmd := c.PostForm("command")
			output, _ := exec.Command("sh", "-c", cmd).Output() // G104: Unhandled error
			c.String(200, string(output))
		})

		api.POST("/upload", func(c *gin.Context) {
			file, header, _ := c.Request.FormFile("file") // G104: Unhandled error
			defer file.Close()
			// G302: Insecure file permissions
			dst, _ := os.Create("/tmp/" + header.Filename) // G104 + Path traversal
			defer dst.Close()
			io.Copy(dst, file)
			c.JSON(200, gin.H{"uploaded": header.Filename})
		})
	}

	port := "8080"
	fmt.Printf("API server on :%s (key: %s)\n", port, API_KEY) // Secret in logs
	http.ListenAndServe(":"+port, r)                           // G104: Unhandled error
}
