package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()

	reader(c,
		gin.H{
			"title":   "Home Page",
			"payload": articles,
		},
		"index.html",
	)

}


func getArticle(c *gin.Context){
	if articleID,err := strconv.Atoi(c.Param("article_id"));err == nil {
		if article, err := getArticleByID(articleID); err == nil {
			reader(c,
				gin.H{
					"title":   article.Title,
					"payload": article,
				},
				"article.html",
			)
		} else {
			c.AbortWithError(http.StatusNotFound, err)
		}
	}else {
			c.AbortWithStatus(http.StatusNotFound)

	}
}