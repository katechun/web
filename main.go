package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var router *gin.Engine

func main(){
	router = gin.Default()
	router.LoadHTMLGlob("templates/*")

	//router.GET("/",func(c *gin.Context){
	//	c.HTML(
	//		http.StatusOK,
	//		"index.html",
	//		gin.H{
	//			"title":"Home Page",
	//		},
	//		)
	//})

	initializeRoutes()
	router.Run()

}


func reader(c *gin.Context,data gin.H,templateName string){
	switch  c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK,data["payload"])
	case "application/xml":
		c.XML(http.StatusOK,data["payload"])
	default:
		c.HTML(http.StatusOK,templateName,data)
	}
}