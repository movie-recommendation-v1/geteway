package main

import (
	pbAdmin "github.com/movie-recommendation-v1/geteway/genproto/userservice"
	"log"

	"github.com/movie-recommendation-v1/geteway/api"
	_ "github.com/movie-recommendation-v1/geteway/api/docs"
	"github.com/movie-recommendation-v1/geteway/api/handler"
	"github.com/movie-recommendation-v1/geteway/api/services"
	pbMovie "github.com/movie-recommendation-v1/geteway/genproto/movieservice"
)

func main() {
	// Load configuration if available
	// cfg := config.Load()  // Uncomment and configure if you have a config package

	// Establish gRPC connections to services
	conn := services.ClientConn()

	// Create a new movie service client
	movieClient := pbMovie.NewMovieServiceClient(conn.MovieClient)
	adminClient := pbAdmin.NewAdminServiceClient(conn.AdminClient)

	// Set up the router with the necessary handlers
	router := api.Router(&handler.Handler{
		Movie: movieClient,
		Admin: adminClient,
	})

	// Start the HTTP server on a configurable port
	// Use cfg.Server.Port if using a configuration file
	port := "localhost:8085"
	log.Printf("Starting server on %s", port)
	if err := router.Run(port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
