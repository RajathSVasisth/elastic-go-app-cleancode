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

func NewProfileRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	pc := &controller.ProfileController{
		ProfileUsecase: usecase.NewProfileUsecase(ur, timeout),
	}
	group.GET("/profile", pc.Fetch)
}
