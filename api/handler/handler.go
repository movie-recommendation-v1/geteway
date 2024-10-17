package handler

import (
	pbMovieService "github.com/movie-recommendation-v1/geteway/genproto/movieservice"
	pbUserService "github.com/movie-recommendation-v1/geteway/genproto/userservice"
)

type Handler struct {
	User    pbUserService.UserServiceClient
	Movie   pbMovieService.MovieServiceClient
	Comment pbMovieService.CommentsServiceClient
	pbMovieService.UnimplementedMovieServiceServer
	pbMovieService.UnimplementedCommentsServiceServer
	pbUserService.UnimplementedUserServiceServer
	pbUserService.UnimplementedAdminServiceServer
}

func NewHandler(h *Handler) *Handler {
	return &Handler{
		User:  h.User,
		Movie:   h.Movie,
		Comment: h.Comment,
	}
}
