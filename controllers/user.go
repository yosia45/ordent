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

type UserController struct {
	userRepo repositories.UserRepository
}

func NewUserController(userRepo repositories.UserRepository) *UserController {
	return &UserController{
		userRepo: userRepo,
	}
}

func (uc *UserController) RegisterUser(c echo.Context) error {
	var user dto.RegisterBodyRequest

	if err := c.Bind(&user); err != nil {
		return utils.HandlerError(c, utils.NewBadRequestError("Invalid request body"))
	}

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

func (uc *UserController) LoginUser(c echo.Context) error {
	var loginBody dto.LoginBodyRequest

	if err := c.Bind(&loginBody); err != nil {
		return utils.HandlerError(c, utils.NewBadRequestError("Invalid request body"))
	}

	if loginBody.Email == "" {
		return utils.HandlerError(c, utils.NewBadRequestError("Email is required"))
	}

	if loginBody.Password == "" {
		return utils.HandlerError(c, utils.NewBadRequestError("Password is required"))
	}

	foundUser, err := uc.userRepo.GetUserByEmail(loginBody.Email)
	if err != nil {
		return utils.HandlerError(c, utils.NewInternalError("Failed to fetch user"))
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(loginBody.Password))
	if err != nil {
		return utils.HandlerError(c, utils.NewUnauthorizedError("Invalid password/email"))
	}

	JWTPayload := dto.JWTPayload{
		UserID:  foundUser.ID,
		IsAdmin: foundUser.IsAdmin,
	}

	token, err := middlewares.GenerateJWT(JWTPayload)
	if err != nil {
		return utils.HandlerError(c, utils.NewInternalError("Failed to generate token"))
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}

func (uc *UserController) MyProfile(c echo.Context) error {
	userPayload := c.Get("userPayload").(*dto.JWTPayload)

	user, err := uc.userRepo.GetUserDetail(userPayload.UserID)
	if err != nil {
		return utils.HandlerError(c, utils.NewInternalError("Failed to fetch user"))
	}

	return c.JSON(http.StatusOK, user)
}
