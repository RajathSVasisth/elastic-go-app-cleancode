package elastic

import (
	"bytes"
	"context"
	"io"

	esv8 "github.com/elastic/go-elasticsearch/v8"
	esv8api "github.com/elastic/go-elasticsearch/v8/esapi"
)

type elasticDB struct {
	client *esv8.Client
}

type ElasticDBMethod interface {
	Index(ctx context.Context, buf bytes.Buffer, indexname string, id string) error
	Delete(ctx context.Context, id string, indexname string) error
	Search(ctx context.Context, buf bytes.Buffer, indexname string) (interface{}, error)
}

func NewElasticDB(client *esv8.Client) ElasticDBMethod {
	return &elasticDB{
		client: client,
	}
}

// Index creates or updates a task in an index.
func (t *elasticDB) Index(ctx context.Context, buf bytes.Buffer, indexname string, id string) error {

	req := esv8api.IndexRequest{
		Index:      indexname,
		Body:       &buf,
		DocumentID: id,
		Refresh:    "true",
	}

	resp, err := req.Do(ctx, t.client)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.IsError() {
		return err
	}

	io.Copy(io.Discard, resp.Body) //nolint: errcheck

	return nil
}

// Delete removes a task from the index.
func (t *elasticDB) Delete(ctx context.Context, id string, indexname string) error {

	req := esv8api.DeleteRequest{
		Index:      indexname,
		DocumentID: id,
	}

	resp, err := req.Do(ctx, t.client)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.IsError() {
		return err
	}

	io.Copy(io.Discard, resp.Body) //nolint: errcheck

	return nil
}

// Search returns tasks matching a query.
//
//nolint:funlen,cyclop
func (t *elasticDB) Search(ctx context.Context, buf bytes.Buffer, indexname string) (interface{}, error) {

	req := esv8api.SearchRequest{
		Index: []string{indexname},
		Body:  &buf,
	}

	resp, err := req.Do(ctx, t.client)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.IsError() {
		return nil, err
	}

	return resp.Body, nil
}
