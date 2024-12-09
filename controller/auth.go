package controller

import (
	"database/sql"
	"golang-crud-api/dto"
	"golang-crud-api/service"
	"golang-crud-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service *service.AuthService
}

func NewAuthController(service *service.AuthService) *AuthController {
	return &AuthController{service: service}
}

// @Summary Register User
// @Description Sample Payload: <br> `{ `<br>` "username": "username", `<br>` "email": "email@gmail.com", `<br>` "password": "password" `<br>` }`
// @Tags Auth
// @Accept json
// @Produce json
// @Param task body dto.RegisterRequest true "Register Request Body"
// @Success 201 {object} utils.APIResponse{status=string,message=string,data=string} "User registered successfully"
// @Failure 400 {object} utils.APIResponse{status=string,message=string} "Invalid input"
// @Failure 500 {object} utils.HTTPError "Failed to register user"
// @Router /auth/register [post]
func (c *AuthController) Register(ctx *gin.Context) {
	var params dto.RegisterRequest

	if err := ctx.ShouldBindJSON(&params); err != nil {
		utils.RespondError(ctx, http.StatusBadRequest, utils.ErrInvalidRequest, err)
		return
	}

	err := c.service.Register(params)
	if err != nil {
		utils.RespondError(ctx, http.StatusInternalServerError, utils.ErrInternalServer, err)
		return
	}

	utils.RespondSuccess(ctx, http.StatusOK, utils.MsgOk, "Register Successful")
}

// @Summary Login User
// @Description Sample Payload: <br> `{ `<br>` "username": "username", `<br>` "password": "password" `<br>` }`
// @Tags Auth
// @Accept json
// @Produce json
// @Param task body dto.LoginRequest true "Login Request Body"
// @Success 201 {object} utils.APIResponse{status=string,message=string,data=string} "Login successful"
// @Failure 400 {object} utils.APIResponse{status=string,message=string} "Invalid input"
// @Failure 500 {object} utils.HTTPError "Failed to login"
// @Router /auth/login [post]
func (c *AuthController) Login(ctx *gin.Context) {
	var request dto.LoginRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		utils.RespondError(ctx, http.StatusBadRequest, utils.ErrInvalidRequest, err)
		return
	}

	response, err := c.service.Login(request)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.RespondError(ctx, http.StatusNotFound, utils.ErrInvalidCredentials, err)
			return
		}

		utils.RespondError(ctx, http.StatusUnauthorized, "Login failed", err)
		return
	}

	utils.RespondSuccess(ctx, http.StatusOK, "Login successful", response)
}
