package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()

	render(
		c,
		gin.H{
			"title":   "Home Page",
			"payload": articles,
		},
		"index.html",
	)
}

func getArticlePage(c *gin.Context) {
	if id, err := strconv.Atoi(c.Param("article_id")); err == nil {
		if a, err := getArticleByID(id); err == nil {
			render(
				c,
				gin.H{
					"title":   a.Title,
					"payload": a,
				},
				"article.html",
			)

		} else {
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		c.AbortWithError(http.StatusBadRequest,
			errors.New("article_id request parameter not provided"),
		)
	}
}
