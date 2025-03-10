package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Welcome to My Gin Website",
	})
}

func about(c *gin.Context) {
	c.HTML(http.StatusOK, "about.html", gin.H{
		"title": "About Us",
	})
}

func login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "Login",
	})
}

func main() {
	// Initialize the Gin router
	r := gin.Default()

	// Load HTML templates from the "templates" folder
	r.LoadHTMLFiles("index.html", "about.html", "login.html")

	// Set up routes
	r.GET("/", home)       // Home route
	r.GET("/about", about) // About route
	r.GET("/login", login) // Login route

	// Run the server
	r.Run(":8080") // Starts the server on port 8080
}
