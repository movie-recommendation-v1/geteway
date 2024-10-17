package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/movie-recommendation-v1/geteway/genproto/movieservice"
	pbComment "github.com/movie-recommendation-v1/geteway/genproto/movieservice"
	logger "github.com/movie-recommendation-v1/geteway/logger"
	"go.uber.org/zap"
)

// CreateComment godoc
// @Summary Add a new comment
// @Description Adds a new comment with the provided details
// @Tags comment
// @Accept json
// @Produce json
// @Param createCommentReq body movieservice.CreateCommentReq true "Comment Add Request"
// @Success 200 {object} movieservice.CreateCommentRes "Success Response"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /comment/add [post]
func (h *Handler) CreateComment(c *gin.Context) {
	log, err := logger.NewLogger()
	if err != nil {
		c.JSON(400, gin.H{"error": "Error with logger"})
		return
	}

	req := pbComment.CreateCommentReq{}
	if err := c.BindJSON(&req); err != nil {
		log.Error("BindJSON error", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Comment.CreateComment(context.Background(), &req)
	if err != nil {
		log.Error("CreateComment error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Info("Comment added successfully", zap.String("comment_id", res.Comment.Id))
	c.JSON(http.StatusOK, res)
}

// GetComment godoc
// @Summary Get comment details by ID
// @Description Retrieves details of a comment by its ID
// @Tags comment
// @Accept json
// @Produce json
// @Param id path string true "Comment ID"
// @Success 200 {object} movieservice.GetCommentRes "Success Response"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /comment/{id} [get]
func (h *Handler) GetComment(c *gin.Context) {
	log, err := logger.NewLogger()
	if err != nil {
		c.JSON(400, gin.H{"error": "Error with logger"})
		return
	}

	id := c.Param("id")
	req := pbComment.GetCommentReq{Id: id}

	res, err := h.Comment.GetComment(context.Background(), &req)
	if err != nil {
		log.Error("GetComment error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Info("Comment retrieved successfully", zap.String("comment_id", id))
	c.JSON(http.StatusOK, res)
}

// UpdateComment godoc
// @Summary Update comment details
// @Description Updates the details of a comment by its ID
// @Tags comment
// @Accept json
// @Produce json
// @Param updateCommentReq body movieservice.UpdateCommentReq true "Comment Update Request"
// @Success 200 {object} movieservice.UpdateCommentRes "Success Response"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /comment/update [put]
func (h *Handler) UpdateComment(c *gin.Context) {
	log, err := logger.NewLogger()
	if err != nil {
		c.JSON(400, gin.H{"error": "Error with logger"})
		return
	}

	req := pbComment.UpdateCommentReq{}
	err = c.BindJSON(&req)
	if err != nil {
		log.Error("BindJSON error", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Comment.UpdateComment(context.Background(), &req)
	if err != nil {
		log.Error("UpdateComment error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Info("Comment updated successfully", zap.String("comment_id", req.Id))
	c.JSON(http.StatusOK, res)
}

// DeleteComment godoc
// @Summary Delete a comment
// @Description Deletes a comment by its ID
// @Tags comment
// @Accept json
// @Produce json
// @Param deleteCommentReq body movieservice.DeleteCommentReq true "Comment Delete Request"
// @Success 200 {object} movieservice.DeleteCommentRes "Success Response"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /comment/delete [delete]
func (h *Handler) DeleteComment(c *gin.Context) {
	log, err := logger.NewLogger()
	if err != nil {
		c.JSON(400, gin.H{"error": "Error with logger"})
		return
	}

	req := pbComment.DeleteCommentReq{}
	if err := c.BindJSON(&req); err != nil {
		log.Error("BindJSON error", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Comment.DeleteComment(context.Background(), &req)
	if err != nil {
		log.Error("DeleteComment error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Info("Comment deleted successfully", zap.String("comment_id", req.Id))
	c.JSON(http.StatusOK, res)
}

// GetAllComments godoc
// @Summary Get all comments
// @Description Retrieves a list of all comments with optional filters
// @Tags comment
// @Accept json
// @Produce json
// @Param user_id query string false "Filter by user ID"
// @Param movie_id query string false "Filter by movie ID"
// @Param rate query int false "Filter by rating"
// @Param limit query int false "Limit the number of results"
// @Param offset query int false "Offset for pagination"
// @Success 200 {object} movieservice.GetAllCommentsRes "Successfully retrieved list of comments"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /comment/all [get]
func (h *Handler) GetAllComments(c *gin.Context) {
	log, err := logger.NewLogger()
	if err != nil {
		c.JSON(400, gin.H{"error": "Error with logger"})
		return
	}

	req := pbComment.GetAllCommentsReq{
		UserId:  c.Query("user_id"),
		MovieId: c.Query("movie_id"),
		Rate: int32(func() int {
			r, _ := strconv.Atoi(c.Query("rate"))
			return r
		}()),
	}

	// Call the service
	res, err := h.Comment.GetAllComments(context.Background(), &req)
	if err != nil {
		log.Error("GetAllComments error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Info("Comments retrieved successfully", zap.Int("count", len(res.Comments)))
	c.JSON(http.StatusOK, res)
}
