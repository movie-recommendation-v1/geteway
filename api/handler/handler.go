package handler

import (
	pbMovieService "github.com/movie-recommendation-v1/geteway/genproto/movieservice"
	pbUserService "github.com/movie-recommendation-v1/geteway/genproto/userservice"
)

type Handler struct {
	User    pbUserService.UserServiceClient
	Admin   pbUserService.AdminServiceClient
	Movie   pbMovieService.MovieServiceClient
	Comment pbMovieService.CommentsServiceClient
}

func NewHandler(h *Handler) *Handler {
	return &Handler{
		User:  h.User,
		Movie: h.Movie,
	}
}
