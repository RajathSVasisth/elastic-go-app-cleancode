package controller

import (
	"net/http"

	"github.com/RajathSVasisth/elastic-go-app-cleancode/domain"
	"github.com/gin-gonic/gin"
)

type ProfileController struct {
	ProfileUsecase domain.ProfileUsecase
}

// Fetch godoc
// @Summary       Fetch user profile
// @Description   Retrieves the profile of the authenticated user
// @Tags          Profile
// @Security      Bearer Token
// @Authorization Bearer <token>
// @Param         authorization header string true "bearer token"
// @Accept        json
// @Produce       json
// @Success       200 {object} domain.Profile
// @Failure       500 {object} domain.ErrorResponse
// @Router        /profile [get]
func (pc *ProfileController) Fetch(c *gin.Context) {
	userID := c.GetString("x-user-id")

	profile, err := pc.ProfileUsecase.GetProfileByID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, profile)
}
