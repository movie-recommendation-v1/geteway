package main

import (
	"fmt"
	"github.com/movie-recommendation-v1/geteway/api/config"
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
	cfg := config.Load()
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
	fmt.Println(cfg.ADMINSERVICEHOST)
	log.Printf("%s:%d", cfg.ADMINSERVICEHOST, cfg.ADMINSERVICEPORT)
	err := router.Run(fmt.Sprintf("%s:%d", cfg.ADMINSERVICEHOST, cfg.ADMINSERVICEPORT))
	if err != nil {
		log.Fatal(err)
	}
}
