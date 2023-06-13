package elastic

import (
	"bytes"
	"context"
	"fmt"
	"io"

	esv7 "github.com/elastic/go-elasticsearch/v7"
	esv7api "github.com/elastic/go-elasticsearch/v7/esapi"
)

type elasticDB struct {
	client *esv7.Client
}

type ElasticDBMethod interface {
	Index(ctx context.Context, buf bytes.Buffer, indexname string, id string) error
	Delete(ctx context.Context, id string, indexname string) error
	Search(ctx context.Context, buf bytes.Buffer, indexname string) (io.ReadCloser, error)
}

func NewElasticDB(client *esv7.Client) ElasticDBMethod {
	return &elasticDB{
		client: client,
	}
}

// Index creates or updates a task in an index.
func (t *elasticDB) Index(ctx context.Context, buf bytes.Buffer, indexname string, id string) error {

	req := esv7api.IndexRequest{
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

	req := esv7api.DeleteRequest{
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
func (t *elasticDB) Search(ctx context.Context, buf bytes.Buffer, indexname string) (io.ReadCloser, error) {
	from := 0
	size := 10
	req := esv7api.SearchRequest{
		Index: []string{indexname},
		Body:  &buf,
		From:  &from,
		Size:  &size,
		Sort:  []string{"{_score:{_id:asc}}"},
	}

	resp, err := req.Do(ctx, t.client)
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		fmt.Println(resp.Status(), err)
		return nil, err
	}

	return resp.Body, nil
}
