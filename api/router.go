package api

import "github.com/gin-gonic/gin"

/**
 * :=  created by:  Shuza
 * :=  create date:  11-Apr-2019
 * :=  (C) CopyRight Shuza
 * :=  www.shuza.ninja
 * :=  shuza.sa@gmail.com
 * :=  Fun  :  Coffee  :  Code
 **/

func GetRoutes() *gin.Engine {
	r := gin.Default()
	r.GET("/", IndexHandler)
	return r
}

func IndexHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Successful",
	})
}
