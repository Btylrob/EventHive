package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Single route for the home page
	router.GET("/", func(c *gin.Context) {
		// Get name from query parameter if provided (e.g., ?name=John)
		name := c.DefaultQuery("name", "Visitor") // If no 'name' query, default to "Visitor"

		// HTML content
		htmlContent := `
        <!DOCTYPE html>
        <html>
        <head>
            <title>Welcome</title>
            <style>
                body {
                    font-family: Arial, sans-serif;
                    margin: 40px;
                    text-align: center;
                    background-color: #f0f0f0;
                }
                .container {
                    max-width: 800px;
                    margin: 0 auto;
                    padding: 20px;
                    background-color: white;
                    border-radius: 5px;
                    box-shadow: 0 2px 5px rgba(0,0,0,0.1);
                }
                h1 {
                    color: #333;
                }
            </style>
        </head>
        <body>
            <div class="container">
                <h1>Hello from Go and Gin!</h1>
                <p>Welcome, ` + name + `!</p>
                <p>This is a simple website built with Go and Gin.</p>
            </div>
        </body>
        </html>`

		// Send the HTML content to the browser
		c.Header("Content-Type", "text/html")
		c.String(http.StatusOK, htmlContent)
	})

	// Start the server
	router.Run(":8080")
}
