package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

const API_KEY = "sk-prod-1234567890abcdef" // G101: Hardcoded credential
const DB_PASSWORD = "admin123"             // G101: Another hardcoded credential

// G401: Weak cryptographic primitive (MD5)
func hashPassword(password string) string {
	h := md5.New()
	h.Write([]byte(password))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// G104: Unhandled errors in utility function
func readConfig(filename string) map[string]string {
	config := make(map[string]string)
	data, _ := ioutil.ReadFile(filename) // G104: Unhandled error

	// G201: SQL query construction using format string (potential)
	query := fmt.Sprintf("SELECT * FROM config WHERE file = '%s'", filename)
	_ = query // Not actually executed but shows pattern
	_ = data  // Simulate usage to avoid compiler error

	return config
}

// G202: SQL query construction using string concatenation
func buildUserQuery(username string) string {
	return "SELECT * FROM users WHERE name = '" + username + "'"
}

// G301: Poor file permissions
func createTempFile(content string) error {
	tmpFile, _ := ioutil.TempFile("", "temp") // G104: Unhandled error
	defer tmpFile.Close()
	tmpFile.WriteString(content)
	os.Chmod(tmpFile.Name(), 0777) // G302: World writable file
	return nil
}

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
