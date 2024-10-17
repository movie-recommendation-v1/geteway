package services

import (
	"fmt"
	"github.com/movie-recommendation-v1/geteway/api/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type Clients struct {
	MovieClient   *grpc.ClientConn
	AdminClient   *grpc.ClientConn
	UserClient    *grpc.ClientConn
	CommentClient *grpc.ClientConn
}

func ClientConn() *Clients {
	cfg := config.Load()

	Movie, err := grpc.NewClient(fmt.Sprintf("%s:%d", cfg.MOVIESERVICEHOST, cfg.MOVIESEVICEPORT), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while connecting to book", err)
		return nil
	}
	Auth, err := grpc.NewClient("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while connecting to book", err)
		return nil
	}
	return &Clients{
		MovieClient:   Movie,
		AdminClient:   Auth,
		UserClient:    Auth,
		CommentClient: Movie,
	}
}
