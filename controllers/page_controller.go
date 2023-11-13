// controllers/page_controller.go
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupPageRoutes initializes routes for the Page controller.
func SetupPageRoutes(router *gin.Engine) {
	router.GET("/", showHomePage)
	router.GET("/about", showAboutPage)
	router.GET("/contact", showContactPage)
	router.GET("/portfolio", showPortfolioPage)
}

func renderPage(c *gin.Context, title string, templateName string, content gin.H) {
	// Merge additional content with the title
	content["Title"] = title

	// Render the page using the layout template
	c.HTML(http.StatusOK, "layout.tmpl", content)
}

func showHomePage(c *gin.Context) {
	data := gin.H{
		"Title":          "Home Page",
		"StaticContent":  "This is the home page content.",
		"StaticContent2": "This is another home page content.",
	}
	renderPage(c, "Home", "home", data)
}

func showAboutPage(c *gin.Context) {
	data := gin.H{
		"Title":         "About",
		"StaticContent": "This is the about page content.",
		// You can add more fields as needed.
	}
	renderPage(c, "About", "about", data)
}

func showContactPage(c *gin.Context) {
	renderPage(c, "Contact", "contact", gin.H{"StaticContent": "This is the contact page content."})
}

func showPortfolioPage(c *gin.Context) {
	data := gin.H{
		"Title":         "Portfolio",
		"StaticContent": "This is the portfolio page content.",
		// You can add more fields as needed.
	}
	renderPage(c, "Portfolio", "portfolio", data)
}
