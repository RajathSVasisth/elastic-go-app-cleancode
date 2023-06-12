package bootstrap

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/RajathSVasisth/elasticApp/mongo"
	esv8 "github.com/elastic/go-elasticsearch/v8"
)

// NewElasticSearch instantiates the ElasticSearch client using configuration defined in environment variables.
func NewElasticSearch(env *Env) (es *esv8.Client, err error) {
	cfg := esv8.Config{
		Addresses: []string{
			env.ElasticURL,
		},
		Username: env.ElasticUsername,
		Password: env.ElasticPassword,
	}
	es, err = esv8.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	res, err := es.Info()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err = res.Body.Close()
	}()

	return es, nil
}

func NewMongoDatabase(env *Env) mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongodbURI := fmt.Sprintf(env.DBURL)

	client, err := mongo.NewClient(mongodbURI)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func CloseMongoDBConnection(client mongo.Client) {
	if client == nil {
		return
	}

	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to MongoDB closed.")
}
