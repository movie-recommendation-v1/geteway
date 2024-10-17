package api

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/movie-recommendation-v1/geteway/api/docs"
	"github.com/movie-recommendation-v1/geteway/api/handler"
	middleware "github.com/movie-recommendation-v1/geteway/api/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Router @title Gateway Service
// @version 1.0
// @description Betayin Kino API Documentation
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func Router(clients *handler.Handler) *gin.Engine {
	router := gin.Default()

	// Configure CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	// Apply custom middleware
	router.Use(middleware.Middleware(router))

	// Set up routes for movies
	movieGroup := router.Group("/movies")
	{
		movieGroup.POST("/create", clients.AddMovie)
		movieGroup.PUT("/update", clients.UpdateMovie)
		movieGroup.GET("/:id", clients.GetMovieById)
		movieGroup.GET("/all", clients.GetAllMovies)
		movieGroup.DELETE("/:id", clients.DeleteMovie)
	}

	return router
}
