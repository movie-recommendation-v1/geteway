package main

import (
	"github.com/movie-recommendation-v1/geteway/api"
	_ "github.com/movie-recommendation-v1/geteway/api/docs"
	"github.com/movie-recommendation-v1/geteway/api/handler"
	"github.com/movie-recommendation-v1/geteway/api/services"
	pbMovie "github.com/movie-recommendation-v1/geteway/genproto/movieservice"
	"log"
)

func main() {
	//cfg := config.Load()
	conn := services.ClientConn()

	movie := pbMovie.NewMovieServiceClient(conn.MovieClient)

	router := api.Router(&handler.Handler{
		Movie: movie,
	})

	err := router.Run("localhost:8085")
	if err != nil {
		log.Fatal(err)
	}
}
