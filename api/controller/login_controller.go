package controller

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/RajathSVasisth/elastic-go-app-cleancode/bootstrap"
	"github.com/RajathSVasisth/elastic-go-app-cleancode/domain"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	LoginUsecase domain.LoginUsecase
	Env          *bootstrap.Env
}

// Login godoc
//
// @Summary     Login user
// @Description Logs in a user and returns access and refresh tokens
// @Tags        Login
// @Accept      json
// @Produce     json
// @Param       request body     domain.LoginRequest true "Login request body"
// @Success     200     {object} domain.LoginResponse
// @Failure     400     {object} domain.ErrorResponse
// @Failure     401     {object} domain.ErrorResponse
// @Failure     404     {object} domain.ErrorResponse
// @Failure     500     {object} domain.ErrorResponse
// @Router      /login [post]
func (lc *LoginController) Login(c *gin.Context) {
	var request domain.LoginRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	user, err := lc.LoginUsecase.GetUserByEmail(c, request.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "User not found with the given email"})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Invalid credentials"})
		return
	}

	accessToken, err := lc.LoginUsecase.CreateAccessToken(&user, lc.Env.AccessTokenSecret, lc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := lc.LoginUsecase.CreateRefreshToken(&user, lc.Env.RefreshTokenSecret, lc.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	loginResponse := domain.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, loginResponse)
}
