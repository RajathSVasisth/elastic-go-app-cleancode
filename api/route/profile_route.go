package route

import (
	"time"

	"github.com/RajathSVasisth/elastic-go-app-cleancode/api/controller"
	"github.com/RajathSVasisth/elastic-go-app-cleancode/bootstrap"
	"github.com/RajathSVasisth/elastic-go-app-cleancode/domain"
	"github.com/RajathSVasisth/elastic-go-app-cleancode/mongo"
	"github.com/RajathSVasisth/elastic-go-app-cleancode/repository"
	"github.com/RajathSVasisth/elastic-go-app-cleancode/usecase"
	"github.com/gin-gonic/gin"
)

func NewProfileRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	pc := &controller.ProfileController{
		ProfileUsecase: usecase.NewProfileUsecase(ur, timeout),
	}
	group.GET("/profile", pc.Fetch)
}
