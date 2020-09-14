package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes() {
	router.GET("/", showIndexPage)
}

func showIndexPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title": "Home Page",
		},
	)
}
