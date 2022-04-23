package server

import (
	"net/http"

	"github.com/FacePlace/go-seo/seo"
	"github.com/gin-gonic/gin"
)

func setRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("/hello", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"msg": "world"})
		})
		api.GET("/parse", func(ctx *gin.Context) {
			urls := ctx.Query("q")
			ctx.Header("Access-Control-Allow-Origin", "*")
			ctx.Header("Access-Control-Allow-Methods", "GET")
			data, _ := seo.GetSEO([]string{urls})
			ctx.String(200, data)
		})
	}

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{})
	})

	return router
}
