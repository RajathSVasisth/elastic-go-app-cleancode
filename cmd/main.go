package main

import (
	"time"

	route "github.com/RajathSVasisth/elastic-go-app-cleancode/api/route"
	"github.com/RajathSVasisth/elastic-go-app-cleancode/bootstrap"
	_ "github.com/RajathSVasisth/elastic-go-app-cleancode/cmd/docs"
	"github.com/gin-gonic/gin"
)

//	@title			Elastic Go App CRUD API
//	@version		1.0
//	@description	This is a sample server celler server.

//	@contact.name	API Support
//	@contact.email	rajath@ozone.one

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8081
//	@BasePath	/

//	@securityDefinitions.basic	BasicAuth

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
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
