package route

import (
	"time"

	"github.com/RajathSVasisth/elastic-go-app-cleancode/api/middleware"
	"github.com/RajathSVasisth/elastic-go-app-cleancode/bootstrap"
	"github.com/RajathSVasisth/elastic-go-app-cleancode/mongo"
	esv7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine, client *esv7.Client) {
	publicRouter := gin.Group("")
	// All Public APIs
	publicRouter.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	NewSignupRouter(env, timeout, db, publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)
	NewRefreshTokenRouter(env, timeout, db, publicRouter)

	protectedRouter := gin.Group("")
	// Middleware to verify AccessToken
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	// All Private APIs
	NewProfileRouter(env, timeout, db, protectedRouter)
	NewTaskRouter(env, timeout, client, protectedRouter)
}
