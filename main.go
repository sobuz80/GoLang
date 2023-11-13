// main.go
package main

import (
	"example.com/mod/controllers" // Adjust this path based on your actual module path
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Load HTML templates

	r.LoadHTMLGlob("views\\templates\\*.tmpl")

	// Initialize and set up routes
	controllers.SetupPageRoutes(r)

	// Serve static files
	r.Static("/static", "./static")

	// Run the application
	r.Run(":8080")
}
