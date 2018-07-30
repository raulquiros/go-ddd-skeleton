package main

import (
	"github.com/gin-gonic/gin"
	"go-ddd-structure/Commands"
)

func Routes(router *gin.Engine) {

	v1 := router.Group("/api/")
	{
		v1.POST("publish-post", Commands.Post())
	}

}