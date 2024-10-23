package handler

import (
	pbMovieService "github.com/movie-recommendation-v1/geteway/genproto/movieservice"
	pbUserService "github.com/movie-recommendation-v1/geteway/genproto/userservice"
	minio "github.com/movie-recommendation-v1/geteway/pkg/minIO"
)

type Handler struct {
	User    pbUserService.UserServiceClient
	Admin   pbUserService.AdminServiceClient
	Movie   pbMovieService.MovieServiceClient
	Comment pbMovieService.CommentsServiceClient
	MinIO   minio.MinIO
}

func NewHandler(h *Handler) *Handler {
	return &Handler{
		User:  h.User,
		Admin: h.Admin,
		Movie: h.Movie,
		MinIO: h.MinIO,
	}
}
