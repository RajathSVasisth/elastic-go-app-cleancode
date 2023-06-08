package route

import (
	"time"

	"github.com/RajathSVasisth/elasticApp/api/controller"
	"github.com/RajathSVasisth/elasticApp/bootstrap"
	"github.com/RajathSVasisth/elasticApp/domain"
	"github.com/RajathSVasisth/elasticApp/mongo"
	"github.com/RajathSVasisth/elasticApp/repository"
	"github.com/RajathSVasisth/elasticApp/usecase"
	"github.com/gin-gonic/gin"
)

func NewTaskRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	tr := repository.NewTaskRepository(db, domain.CollectionTask)
	tc := &controller.TaskController{
		TaskUsecase: usecase.NewTaskUsecase(tr, timeout),
	}
	group.GET("/task", tc.Fetch)
	group.POST("/task", tc.Create)
}
