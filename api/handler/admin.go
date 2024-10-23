package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/movie-recommendation-v1/geteway/genproto/userservice"
	pbAdmin "github.com/movie-recommendation-v1/geteway/genproto/userservice"
	logger "github.com/movie-recommendation-v1/geteway/pkg/logger"
	"go.uber.org/zap"
)

// CreateAdmin godoc
// @Summary Add a new admin
// @Description Adds a new admin with the provided details
// @Tags admin
// @Accept json
// @Produce json
// @Param createAdminReq body userservice.CreateAdminReq true "Admin Add Request"
// @Success 200 {object} userservice.CreateAdminRes "Success Response"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /admin/add [post]
func (h *Handler) CreateAdmin(c *gin.Context) {
	log, err := logger.NewLogger()
	if err != nil {
		c.JSON(400, gin.H{"error": "Error with logger"})
		return
	}

	req := pbAdmin.CreateAdminReq{}
	if err := c.BindJSON(&req); err != nil {
		log.Error("BindJSON error", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Admin.CreateAdmin(context.Background(), &req)
	if err != nil {
		log.Error("CreateAdmin error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Info("Admin added successfully")
	c.JSON(http.StatusOK, res)
}

// UpdateAdmin godoc
// @Summary Update admin details
// @Description Updates the details of an admin by its ID
// @Tags admin
// @Accept json
// @Produce json
// @Param updateAdminReq body userservice.UpdateAdminReq true "Admin Update Request"
// @Success 200 {object} userservice.UpdateAdminRes "Success Response"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /admin/update [put]
func (h *Handler) UpdateAdmin(c *gin.Context) {
	log, err := logger.NewLogger()
	if err != nil {
		c.JSON(400, gin.H{"error": "Error with logger"})
		return
	}

	req := pbAdmin.UpdateAdminReq{}
	if err := c.BindJSON(&req); err != nil {
		log.Error("BindJSON error", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Admin.UpdateAdmin(context.Background(), &req)
	if err != nil {
		log.Error("UpdateAdmin error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Info("Admin updated successfully", zap.String("admin_id", req.AdminReq.Id))
	c.JSON(http.StatusOK, res)
}

// GetAdmin godoc
// @Summary Get admin details by ID
// @Description Retrieves details of an admin by its ID
// @Tags admin
// @Accept json
// @Produce json
// @Param id path string true "Admin ID"
// @Success 200 {object} userservice.GetAdminRes "Success Response"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /admin/{id} [get]
func (h *Handler) GetAdmin(c *gin.Context) {
	log, err := logger.NewLogger()
	if err != nil {
		c.JSON(400, gin.H{"error": "Error with logger"})
		return
	}

	id := c.Param("id")
	req := pbAdmin.GetAdminReq{AdminId: id}

	res, err := h.Admin.GetAdmin(context.Background(), &req)
	if err != nil {
		log.Error("GetAdmin error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Info("Admin retrieved successfully", zap.String("admin_id", id))
	c.JSON(http.StatusOK, res)
}

// ForgetPassword godoc
// @Summary Reset admin password
// @Description Sends a password reset link to the admin's email
// @Tags admin
// @Accept json
// @Produce json
// @Param forgetPasswordReq body userservice.ForgetPasswordReq true "Password Reset Request"
// @Success 200 {object} userservice.ForgetPasswordRes "Success Response"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /admin/forget_password [post]
func (h *Handler) ForgetPassword(c *gin.Context) {
	log, err := logger.NewLogger()
	if err != nil {
		c.JSON(400, gin.H{"error": "Error with logger"})
		return
	}

	req := pbAdmin.ForgetPasswordReq{}
	if err := c.BindJSON(&req); err != nil {
		log.Error("BindJSON error", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Admin.ForgetPassword(context.Background(), &req)
	if err != nil {
		log.Error("ForgetPassword error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Info("Password reset email sent successfully")
	c.JSON(http.StatusOK, res)
}

// GetAllAdmins godoc
// @Summary Get all admins
// @Description Retrieves a list of all admins with optional filters
// @Tags admin
// @Accept json
// @Produce json
// @Param limit query int false "Limit the number of results"
// @Param offset query int false "Offset for pagination"
// @Success 200 {object} userservice.GetAllAdminRes "Successfully retrieved list of admins"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /admin/all [get]
func (h *Handler) GetAllAdmins(c *gin.Context) {
	log, err := logger.NewLogger()
	if err != nil {
		c.JSON(400, gin.H{"error": "Error with logger"})
		return
	}

	req := pbAdmin.GetAllAdminReq{

		Limit: int32(func() int {
			l, _ := strconv.Atoi(c.Query("limit"))
			return l
		}()),
		Offset: int32(func() int {
			o, _ := strconv.Atoi(c.Query("offset"))
			return o
		}()),
	}

	// Call the service
	res, err := h.Admin.GetAllAdmin(context.Background(), &req)
	if err != nil {
		log.Error("GetAllAdmins error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Info("Admins retrieved successfully", zap.Int("count", len(res.AdminRes)))
	c.JSON(http.StatusOK, res)
}

// DeleteAdmin godoc
// @Summary Delete an admin
// @Description Deletes an admin by its ID
// @Tags admin
// @Accept json
// @Produce json
// @Param deleteAdminReq body userservice.DeleteAdminReq true "Admin Delete Request"
// @Success 200 {object} userservice.DeleteAdminRes "Success Response"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /admin/delete [delete]
func (h *Handler) DeleteAdmin(c *gin.Context) {
	log, err := logger.NewLogger()
	if err != nil {
		c.JSON(400, gin.H{"error": "Error with logger"})
		return
	}

	req := pbAdmin.DeleteAdminReq{}
	if err := c.BindJSON(&req); err != nil {
		log.Error("BindJSON error", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Admin.DeleteAdmin(context.Background(), &req)
	if err != nil {
		log.Error("DeleteAdmin error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Info("Admin deleted successfully", zap.String("admin_id", req.AdminId))
	c.JSON(http.StatusOK, res)
}
