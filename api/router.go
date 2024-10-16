package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/movie-recommendation-v1/geteway/api/handler"
	middlerware "github.com/movie-recommendation-v1/geteway/api/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"time"
)

// @tite geteway service
// @version 1.0
// @description Betayin Kino
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authourization
func Router(Clients *handler.Handler) *gin.Engine {
	router := gin.Default()
	url := ginSwagger.URL("swagger/doc.json")

	router.GET("api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	router.Use(middlerware.Middleware(router))
	h := handler.NewHandler(Clients)
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, //df
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))
	movieGroup := router.Group("/movies")
	{
		movieGroup.POST("/create", h.AddMovie)
		movieGroup.PUT("/update", h.UpdateMovie)
		movieGroup.GET("/:id", h.GetMovieById)
		movieGroup.GET("/all", h.GetAllMovies)
		movieGroup.DELETE("/:id", h.DeleteMovie)
	}
	return router
}
