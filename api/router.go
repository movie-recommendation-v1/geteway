package api

import (
	"github.com/gin-gonic/gin"
	"github.com/movie-recommendation-v1/geteway/api/handler"
	middlerware "github.com/movie-recommendation-v1/geteway/api/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"time"
)

func Router(Clients *handler.Handler) gin.Engine {
	router := gin.Default()
	url := ginSwagger.URL("swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	router.Use(middlerware.Middleware(router))

	h := handler.NewHandler(clients)

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))
}
