package controller

import (
	"net/http"
	"strconv"

	"github.com/RajathSVasisth/elastic-go-app-cleancode/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskController struct {
	TaskUsecase domain.TaskUsecase
}

// Create godoc
// @Summary     Create a new task
// @Description Creates a new task for the authenticated user
// @Tags        Tasks
// @Security    ApiKeyAuth
// @Accept      json
// @Produce     json
// @Param       authorization header   string true "bearer token"
// @Param       task          body     domain.Task true "Task object"
// @Success     200           {object} domain.SuccessResponse
// @Failure     400           {object} domain.ErrorResponse
// @Failure     500           {object} domain.ErrorResponse
// @Router      /tasks [post]
func (tc *TaskController) Create(c *gin.Context) {
	var task domain.Task

	err := c.ShouldBind(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	userID := c.GetString("x-user-id")
	task.ID = primitive.NewObjectID()

	task.UserID, err = primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = tc.TaskUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Task created successfully",
	})
}

// Fetch godoc
// @Summary     Fetch user tasks
// @Description Retrieves tasks for the authenticated user within the specified pagination range
// @Tags        Tasks
// @Security    ApiKeyAuth
// @Accept      json
// @Produce     json
// @Param       authorization header   string  true  "bearer token"
// @Param       from          query    integer false "Pagination start index (default: 0)"
// @Param       to            query    integer false "Pagination end index (default: 10)"
// @Success     200           {object} []domain.Task
// @Failure     400           {object} domain.ErrorResponse
// @Failure     500           {object} domain.ErrorResponse
// @Router      /tasks [get]
func (u *TaskController) Fetch(c *gin.Context) {
	userID := c.GetString("x-user-id")

	var pagination domain.Pagination
	from, err := strconv.Atoi(c.DefaultQuery("from", "0"))
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	to, err := strconv.Atoi(c.DefaultQuery("to", "10"))
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	size := to - from
	pagination.From = from
	pagination.Size = size

	tasks, err := u.TaskUsecase.FetchByUserID(c, userID, pagination)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// Update godoc
// @Summary     Update a task
// @Description Updates an existing task for the authenticated user
// @Tags        Tasks
// @Security    ApiKeyAuth
// @Accept      json
// @Produce     json
// @Param       authorization header   string      true "bearer token"
// @Param       id            path     string      true "Task ID"
// @Param       task          body     domain.Task true "Task object"
// @Success     200           {object} domain.SuccessResponse
// @Failure     400           {object} domain.ErrorResponse
// @Failure     500           {object} domain.ErrorResponse
// @Router      /tasks/{id} [put]
func (u *TaskController) Update(c *gin.Context) {
	var task domain.Task
	var err error

	err = c.ShouldBind(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	task.UserID, err = primitive.ObjectIDFromHex(c.GetString("x-user-id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	task.ID, err = primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = u.TaskUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Task updated successfully",
	})
}

// Delete godoc
// @Summary     Delete a task
// @Description Deletes a task for the authenticated user
// @Tags        Tasks
// @Security    ApiKeyAuth
// @Accept      json
// @Produce     json
// @Param       authorization header   string      true "bearer token"
// @Param       id            path     string true "Task ID"
// @Success     200           {object} domain.SuccessResponse
// @Failure     500           {object} domain.ErrorResponse
// @Router      /tasks/{id} [delete]

func (u *TaskController) Delete(c *gin.Context) {
	id := c.Param("id")

	err := u.TaskUsecase.Delete(c, &id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Task deleted successfully",
	})
}
