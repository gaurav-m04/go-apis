package router

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Init(ctx context.Context) {
	router := setupRouter()

	srv := &http.Server{
		Addr:    "dsds",
		Handler: router,
	}

	go func(ctx context.Context) {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("%s", err)
		}
	}(ctx)

	// Listen signal

}

func setupRouter() *gin.Engine {
	route := gin.New()
	registerRoutes(route)
	return route
}

func registerRoutes(router *gin.Engine) {
	routeGroup := router.Group("/")

	routeGroup.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "hello",
		})
	})
}
