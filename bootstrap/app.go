package bootstrap

import (
	"log"

	"github.com/RajathSVasisth/elasticApp/mongo"
	esv8 "github.com/elastic/go-elasticsearch/v8"
)

type Application struct {
	Env    *Env
	Client *esv8.Client
	Mongo  mongo.Client
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Mongo = NewMongoDatabase(app.Env)
	var err error
	app.Client, err = NewElasticSearch(app.Env)
	if err != nil {
		log.Fatal(err)
	}
	return *app
}

func (app *Application) CloseDBConnection() {
	CloseMongoDBConnection(app.Mongo)
}
