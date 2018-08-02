package main

import (
	"github.com/gin-gonic/gin"
	"go-ddd-structure/Controllers"
)

func Routes(router *gin.Engine) {

	v1 := router.Group("/api/")
	{
		v1.POST("publish-post", Controllers.CreatePost)
	}

}
