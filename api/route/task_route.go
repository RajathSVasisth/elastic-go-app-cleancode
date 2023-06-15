package route

import (
	"time"

	"github.com/RajathSVasisth/elastic-go-app-cleancode/api/controller"
	"github.com/RajathSVasisth/elastic-go-app-cleancode/bootstrap"
	"github.com/RajathSVasisth/elastic-go-app-cleancode/domain"
	"github.com/RajathSVasisth/elastic-go-app-cleancode/elastic"
	"github.com/RajathSVasisth/elastic-go-app-cleancode/repository"
	"github.com/RajathSVasisth/elastic-go-app-cleancode/usecase"
	esv7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/gin-gonic/gin"
)

func NewTaskRouter(env *bootstrap.Env, timeout time.Duration, client *esv7.Client, group *gin.RouterGroup) {
	tr := repository.NewTaskRepository(elastic.NewElasticDB(client), domain.Index)
	tc := &controller.TaskController{
		TaskUsecase: usecase.NewTaskUsecase(tr, timeout),
	}
	group.GET("/task", tc.Fetch)
	group.POST("/task", tc.Create)
	group.PUT("/task/:id", tc.Update)
	group.DELETE("/task/:id", tc.Delete)
}
