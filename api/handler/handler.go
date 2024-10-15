package handler

import (
	pbMovieService "github.com/movie-recommendation-v1/geteway/genproto/movieservice"
	pbUserService "github.com/movie-recommendation-v1/geteway/genproto/userservice"
)

type Handler struct {
	user  pbUserService.UserServiceClient
	movie pbMovieService.MovieServiceClient
}

func NewHandler(h *Handler) *Handler {
	return h
}
