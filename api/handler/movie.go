package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	pbMovie "github.com/movie-recommendation-v1/geteway/genproto/movieservice"
	logger "github.com/movie-recommendation-v1/geteway/pkg/logger"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

// AddMovie godoc
// @Summary Add a new movie
// @Description Adds a new movie with the provided details
// @Tags movie
// @Accept json
// @Produce json
// @Param addMovieReq body movieservice.AddMovieReq true "Movie Add Request"
// @Success 200 {object} movieservice.AddMovieRes "Success Response"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /movie/add [post]
func (h *Handler) AddMovie(c *gin.Context) {
	log, err := logger.NewLogger()
	if err != nil {
		c.JSON(400, gin.H{"error": "Error with logger"})
		return
	}

	req := pbMovie.AddMovieReq{}
	if err := c.BindJSON(&req); err != nil {
		log.Error("BindJSON error", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Movie.AddMovie(context.Background(), &req)
	if err != nil {
		log.Error("AddMovie error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Info("Movie added successfully", zap.String("movie_name", req.MovieName))
	c.JSON(http.StatusOK, res)
}

// GetMovieById godoc
// @Summary Get movie details by ID
// @Description Retrieves details of a movie by its ID
// @Tags movie
// @Accept json
// @Produce json
// @Param id path string true "Movie ID"
// @Success 200 {object} movieservice.GetMovieByIdRes "Success Response"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /movie/{id} [get]
func (h *Handler) GetMovieById(c *gin.Context) {
	log, err := logger.NewLogger()
	if err != nil {
		c.JSON(400, gin.H{"error": "Error with logger"})
		return
	}

	id := c.Param("id")
	req := pbMovie.GetMovieByIdReq{Id: id}

	res, err := h.Movie.GetMovieById(context.Background(), &req)
	if err != nil {
		log.Error("GetMovieById error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Info("Movie retrieved successfully", zap.String("movie_id", id))
	c.JSON(http.StatusOK, res)
}

// UpdateMovie godoc
// @Summary Update movie details
// @Description Updates the details of a movie by its ID
// @Tags movie
// @Accept json
// @Produce json
// @Param updateMovieReq body movieservice.UpdateMovieReq true "Movie Update Request"
// @Success 200 {object} movieservice.UpdateMovieRes "Success Response"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /movie/update [put]
func (h *Handler) UpdateMovie(c *gin.Context) {
	log, err := logger.NewLogger()
	if err != nil {
		c.JSON(400, gin.H{"error": "Error with logger"})
		return
	}

	req := pbMovie.UpdateMovieReq{}
	err = c.BindJSON(&req)
	if err != nil {
		log.Error("BindJSON error", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Movie.UpdateMovie(context.Background(), &req)
	if err != nil {
		log.Error("UpdateMovie error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Info("Movie updated successfully", zap.String("movie_id", req.Id))
	c.JSON(http.StatusOK, res)
}

// DeleteMovie godoc
// @Summary Delete a movie
// @Description Deletes a movie by its ID
// @Tags movie
// @Accept json
// @Produce json
// @Param deleteMovieReq body movieservice.DeleteMovieReq true "Movie Delete Request"
// @Success 200 {object} movieservice.DeleteMovieRes "Success Response"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /movie/delete [delete]
func (h *Handler) DeleteMovie(c *gin.Context) {
	log, err := logger.NewLogger()
	if err != nil {
		c.JSON(400, gin.H{"error": "Error with logger"})
		return
	}

	req := pbMovie.DeleteMovieReq{}
	if err := c.BindJSON(&req); err != nil {
		log.Error("BindJSON error", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Movie.DeleteMovie(context.Background(), &req)
	if err != nil {
		log.Error("DeleteMovie error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Info("Movie deleted successfully", zap.String("movie_id", req.Id))
	c.JSON(http.StatusOK, res)
}

// GetAllMovies godoc
// @Summary Get all movies
// @Description Retrieves a list of all movies with optional filters
// @Tags movie
// @Accept json
// @Produce json
// @Param name query string false "Filter by movie name"
// @Param studio query string false "Filter by studio"
// @Param genre query string false "Filter by genre"
// @Param limit query int false "Limit the number of results"
// @Param offset query int false "Offset for pagination"
// @Success 200 {object} movieservice.GetAllMoviesRes "Successfully retrieved list of movies"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /movie/all [get]
func (h *Handler) GetAllMovies(c *gin.Context) {
	log, err := logger.NewLogger()
	if err != nil {
		c.JSON(400, gin.H{"error": "Error with logger"})
		return
	}

	// Convert query parameters
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		log.Error("Invalid limit parameter", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
		return
	}
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		log.Error("Invalid offset parameter", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset parameter"})
		return
	}

	// Prepare request with optional filters
	req := pbMovie.GetAllMoviesReq{
		MovieName: c.Query("name"),
		Studio:    c.Query("studio"),
		Genres:    []*pbMovie.Genres{{Genre: c.Query("genre")}},
		Limit:     int32(limit),
		Offset:    int32(offset),
	}

	// Call the service
	res, err := h.Movie.GetAllMovies(context.Background(), &req)
	if err != nil {
		log.Error("GetAllMovies error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Info("Movies retrieved successfully", zap.Int("count", len(res.Movies)))
	c.JSON(http.StatusOK, res)
}
