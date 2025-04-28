package controllers

import (
	"net/http"
	"ordent/dto"
	"ordent/middlewares"
	"ordent/models"
	"ordent/repositories"
	"ordent/utils"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// UserController handles user-related requests
// @Description This controller is responsible for user registration, login, and profile fetching
type UserController struct {
	userRepo repositories.UserRepository
}

// NewUserController creates a new instance of UserController
// @Description Create a new UserController with a UserRepository dependency
func NewUserController(userRepo repositories.UserRepository) *UserController {
	return &UserController{
		userRepo: userRepo,
	}
}

// RegisterUser godoc
// @Summary Register a new user
// @Description Create a new user with the provided details.
// @Tags users
// @Accept json
// @Produce json
// @Param user body dto.RegisterBodyRequest true "User registration details"
// @Success 201 {object} map[string]string "User created successfully"
// @Failure 400 {object} utils.APIError "Invalid request body"
// @Failure 500 {object} utils.APIError "Internal server error"
// @Router /api/v1/register [post]
func (uc *UserController) RegisterUser(c echo.Context) error {
	var user dto.RegisterBodyRequest

	if err := c.Bind(&user); err != nil {
		return utils.HandlerError(c, utils.NewBadRequestError("Invalid request body"))
	}

	// Validations
	if user.FullName == "" {
		return utils.HandlerError(c, utils.NewBadRequestError("Name is required"))
	}

	if user.Email == "" {
		return utils.HandlerError(c, utils.NewBadRequestError("Email is required"))
	}

	if user.Password == "" {
		return utils.HandlerError(c, utils.NewBadRequestError("Password is required"))
	}

	if user.Username == "" {
		return utils.HandlerError(c, utils.NewBadRequestError("Username is required"))
	}

	newUser := &models.User{
		FullName: user.FullName,
		Email:    user.Email,
		Username: user.Username,
		Password: user.Password,
		IsAdmin:  user.IsAdmin,
	}

	if err := uc.userRepo.CreateUser(newUser); err != nil {
		return utils.HandlerError(c, utils.NewInternalError("Failed to create user"))
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "User created successfully"})
}

// LoginUser godoc
// @Summary Login a user
// @Description Authenticate a user with email and password, and return a JWT token.
// @Tags users
// @Accept json
// @Produce json
// @Param login body dto.LoginBodyRequest true "Login details"
// @Success 200 {object} map[string]string "JWT token"
// @Failure 400 {object} utils.APIError "Invalid input data"
// @Failure 401 {object} utils.APIError "Invalid email/password"
// @Failure 500 {object} utils.APIError "Internal server error"
// @Router /api/v1/login [post]
func (uc *UserController) LoginUser(c echo.Context) error {
	var loginBody dto.LoginBodyRequest

	if err := c.Bind(&loginBody); err != nil {
		return utils.HandlerError(c, utils.NewBadRequestError("Invalid request body"))
	}

	// Validations
	if loginBody.Email == "" {
		return utils.HandlerError(c, utils.NewBadRequestError("Email is required"))
	}

	if loginBody.Password == "" {
		return utils.HandlerError(c, utils.NewBadRequestError("Password is required"))
	}

	// Find user based on email
	foundUser, err := uc.userRepo.GetUserByEmail(loginBody.Email)
	if err != nil {
		return utils.HandlerError(c, utils.NewInternalError("Failed to fetch user"))
	}

	// Compare the stored hashed password in database with the password body
	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(loginBody.Password))
	if err != nil {
		return utils.HandlerError(c, utils.NewUnauthorizedError("Invalid password/email"))
	}

	// Create JWT Payload for JWT Token
	JWTPayload := dto.JWTPayload{
		UserID:  foundUser.ID,
		IsAdmin: foundUser.IsAdmin,
	}

	// Create JWT Token based on payload
	token, err := middlewares.GenerateJWT(JWTPayload)
	if err != nil {
		return utils.HandlerError(c, utils.NewInternalError("Failed to generate token"))
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}

// MyProfile godoc
// @Summary Get My Profile
// @Description Get user profile. This endpoint can only be accessed by users with `isAdmin = false`.
// @Tags user
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Success 200 {object} dto.GetUserDetailResponse
// @Failure 403 {object} utils.APIError "Forbidden"
// @Failure 401 {object} utils.APIError "Unauthorized"
// @Failure 500 {object} utils.APIError "Internal Server Error"
// @Router /api/v1/myprofiles [get]
func (uc *UserController) MyProfile(c echo.Context) error {
	userPayload := c.Get("userPayload").(*dto.JWTPayload)

	user, err := uc.userRepo.GetUserDetail(userPayload.UserID)
	if err != nil {
		return utils.HandlerError(c, utils.NewInternalError("Failed to fetch user"))
	}

	return c.JSON(http.StatusOK, user)
}
