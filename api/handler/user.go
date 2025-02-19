package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	token2 "github.com/movie-recommendation-v1/geteway/api/token"
	_ "github.com/movie-recommendation-v1/geteway/genproto/userservice"
	pbUser "github.com/movie-recommendation-v1/geteway/genproto/userservice"
	logger "github.com/movie-recommendation-v1/geteway/pkg/logger"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

// Regist	erUser godoc
// @Summary Register a new user
// @Description Registers a new user with the provided details
// @Tags user
// @Accept json
// @Produce json
// @Param registerUserReq body userservice.RegisterUserReq true "User Register Request"
// @Success 200 {object} userservice.RegisterUserRes "Success Response"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /user/register [post]
func (h *Handler) RegisterUser(c *gin.Context) {
	log, err := logger.NewLogger()
	if err != nil {
		c.JSON(400, gin.H{"error": "Error with logger"})
		return
	}
	fmt.Println(11111)
	req := pbUser.RegisterUserReq{}
	fmt.Println(h.User)
	if err := c.BindJSON(&req); err != nil {
		log.Error("BindJSON error", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.User.RegisterUser(context.Background(), &req)
	if err != nil {
		log.Error("RegisterUser error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(222)

	log.Info("code sanded in email")
	c.JSON(http.StatusOK, res)
}

// VerifyUser godoc
// @Summary Verify a user
// @Description Verifies a user using the verification code
// @Tags user
// @Accept json
// @Produce json
// @Param verifyUserReq body userservice.VerifyUserReq true "User Verification Request"
// @Success 200 {object} userservice.VerifyUserRes "Success Response"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /user/verify-user [post]
func (h *Handler) VerifyUser(c *gin.Context) {
	log, err := logger.NewLogger()
	if err != nil {
		c.JSON(400, gin.H{"error": "Error with logger"})
		return
	}
	req := pbUser.VerifyUserReq{}
	err = c.BindJSON(&req)
	if err != nil {
		log.Error("BindJSON error", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.User.VerifyUser(context.Background(), &req)
	if err != nil {
		log.Error("VerifyUser error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	token := token2.GenereteJWTTokenForUser(&pbUser.LoginRes{UserRes: res.Res})
	c.JSON(http.StatusOK, gin.H{"message": "User successfully verify", "token": token})
}

// LoginUser godoc
// @Summary Login user
// @Description Logins user using email and password
// @Tags user
// @Accept json
// @Produce json
// @Param loginReq body userservice.LoginReq true "Login Request"
// @Success 200 {object} userservice.LoginRes "Success Response"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /user/login [post]
func (h *Handler) LoginUser(c *gin.Context) {
	log, err := logger.NewLogger()

	if err != nil {
		c.JSON(400, gin.H{"error": "Error with logger"})
		return
	}

	req := pbUser.LoginReq{}
	if err := c.BindJSON(&req); err != nil {
		log.Error("BindJSON error", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	println(req.Email)
	println(req.Password)
	res, err := h.User.Login(context.Background(), &req)
	if err != nil {
		log.Error("Login error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	token := token2.GenereteJWTTokenForUser(res)
	log.Info("User logged in successfully", zap.String("user_id", res.UserRes.Id))
	c.JSON(http.StatusOK, gin.H{"User": res.UserRes, "Token": token})
}

// ForgotPassword godoc
// @Summary Forgot password
// @Description Sends a forgot password email
// @Tags user
// @Accept json
// @Produce json
// @Param forgotPasswordReq body userservice.ForgotPasswordReq true "Forgot Password Request"
// @Success 200 {object} userservice.ForgotPasswordRes "Success Response"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /user/forgot_password [post]
func (h *Handler) ForgotPassword(c *gin.Context) {
	log, err := logger.NewLogger()
	if err != nil {
		c.JSON(400, gin.H{"error": "Error with logger"})
		return
	}

	req := pbUser.ForgotPasswordReq{}
	if err := c.BindJSON(&req); err != nil {
		log.Error("BindJSON error", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.User.ForgotPassword(context.Background(), &req)
	if err != nil {
		log.Error("ForgotPassword error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Info("Forgot password email sent successfully", zap.String("message", res.Message))
	c.JSON(http.StatusOK, res)
}

// GetUserByID godoc
// @Summary Get user details by ID
// @Description Retrieves details of a user by their ID
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} userservice.GetUserByIDRes "Success Response"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /user/{id} [get]
func (h *Handler) GetUserByID(c *gin.Context) {
	log, err := logger.NewLogger()
	if err != nil {
		c.JSON(400, gin.H{"error": "Error with logger"})
		return
	}

	id := c.Param("id")
	req := pbUser.GetUserByIDReq{Userid: id}

	res, err := h.User.GetUserByID(context.Background(), &req)
	if err != nil {
		log.Error("GetUserByID error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Info("User retrieved successfully", zap.String("user_id", id))
	c.JSON(http.StatusOK, res)
}

// UpdateUser godoc
// @Summary Update user details
// @Description Updates the details of a user by their ID
// @Tags user
// @Accept json
// @Produce json
// @Param updateUserReq body userservice.UpdateUserReq true "User Update Request"
// @Success 200 {object} userservice.UpdateUserRes "Success Response"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /user/update [put]
func (h *Handler) UpdateUser(c *gin.Context) {
	log, err := logger.NewLogger()
	if err != nil {
		c.JSON(400, gin.H{"error": "Error with logger"})
		return
	}

	req := pbUser.UpdateUserReq{}
	err = c.BindJSON(&req)
	if err != nil {
		log.Error("BindJSON error", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.User.UpdateUser(context.Background(), &req)
	if err != nil {
		log.Error("UpdateUser error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Info("User updated successfully", zap.String("user_id", req.UserReq.Id))
	c.JSON(http.StatusOK, res)
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Deletes a user by their ID
// @Tags user
// @Accept json
// @Produce json
// @Param deleteUserReq body userservice.DeleteUserReq true "User Delete Request"
// @Success 200 {object} userservice.DeleteUserRes "Success Response"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /user/delete [delete]
func (h *Handler) DeleteUser(c *gin.Context) {
	log, err := logger.NewLogger()
	if err != nil {
		c.JSON(400, gin.H{"error": "Error with logger"})
		return
	}

	req := pbUser.DeleteUserReq{}
	if err := c.BindJSON(&req); err != nil {
		log.Error("BindJSON error", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.User.DeleteUser(context.Background(), &req)
	if err != nil {
		log.Error("DeleteUser error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Info("User deleted successfully", zap.String("user_id", req.Id))
	c.JSON(http.StatusOK, res)
}

// GetAllUsers godoc
// @Summary Get all users
// @Description Retrieves a list of all users with optional filters
// @Tags user
// @Accept json
// @Produce json
// @Param limit query int false "Limit the number of results"
// @Param offset query int false "Offset for pagination"
// @Param id query string false "User id"
// @Param name query string false "User-name"
// @Param email query string false "User email"
// Param img_url query string false "User img-url"
// @Success 200 {object} userservice.GetAllUserRes "Successfully retrieved list of users"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /user/all [get]
func (h *Handler) GetAllUsers(c *gin.Context) {
	log, err := logger.NewLogger()
	if err != nil {
		c.JSON(400, gin.H{"error": "Error with logger"})
		return
	}

	req := pbUser.GetAllUserReq{
		UserReq: &pbUser.UserModel{
			Id:     c.Query("id"),
			Name:   c.Query("name"),
			Email:  c.Query("email"),
			ImgUrl: c.Query("img_url"),
		},
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
	res, err := h.User.GetAllUser(context.Background(), &req)
	if err != nil {
		log.Error("GetAllUsers error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Info("Users retrieved successfully", zap.Int("count", len(res.UserRes)))
	c.JSON(http.StatusOK, res)
}
