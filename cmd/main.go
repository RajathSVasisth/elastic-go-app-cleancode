package main

import (
	"time"

	route "github.com/RajathSVasisth/elastic-go-app-cleancode/api/route"
	"github.com/RajathSVasisth/elastic-go-app-cleancode/bootstrap"
	"github.com/gin-gonic/gin"
)

func main() {

	app := bootstrap.App()

	env := app.Env

	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	route.Setup(env, timeout, db, gin, app.Client)

	gin.Run(env.ServerAddress)
}
