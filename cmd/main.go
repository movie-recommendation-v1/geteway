package main

import (
	"github.com/movie-recommendation-v1/geteway/api/handler"
	pbMovie "github.com/movie-recommendation-v1/geteway/genproto/movieservice"
	pbAdmin "github.com/movie-recommendation-v1/geteway/genproto/userservice"
	"log"

	"github.com/movie-recommendation-v1/geteway/api"
	_ "github.com/movie-recommendation-v1/geteway/api/docs"
	"github.com/movie-recommendation-v1/geteway/api/services"
)

func main() {
	conn := services.ClientConn()

	//movieClient := pbMovie.NewMovieServiceClient(conn.MovieClient)
	UserClient := pbAdmin.NewUserServiceClient(conn.UserClient)
	AdminClient := pbAdmin.NewAdminServiceClient(conn.AdminClient)
	MovieClient := pbMovie.NewMovieServiceClient(conn.MovieClient)
	CommentClient := pbMovie.NewCommentsServiceClient(conn.CommentClient)

	router := api.Router(&handler.Handler{
		User:    UserClient,
		Admin:   AdminClient,
		Movie:   MovieClient,
		Comment: CommentClient,
	})

	log.Printf("Starting server on %s", "localhost:8085")
	err := router.Run(":8085")
	if err != nil {
		log.Fatal(err)
	}
}
