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

func NewRefreshTokenRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	rtc := &controller.RefreshTokenController{
		RefreshTokenUsecase: usecase.NewRefreshTokenUsecase(ur, timeout),
		Env:                 env,
	}
	group.POST("/refresh", rtc.RefreshToken)
}
