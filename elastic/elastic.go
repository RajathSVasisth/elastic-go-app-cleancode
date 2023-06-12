package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/RajathSVasisth/elasticApp/domain"
	esv8 "github.com/elastic/go-elasticsearch/v8"
	esv8api "github.com/elastic/go-elasticsearch/v8/esapi"
)

type ElasticDB struct {
	Client *esv8.Client
}

// Index creates or updates a task in an index.
func (t *ElasticDB) Index(ctx context.Context, task domain.Task, indexname string) error {

	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(task); err != nil {
		return err
	}

	req := esv8api.IndexRequest{
		Index:      indexname,
		Body:       &buf,
		DocumentID: task.ID.Hex(),
		Refresh:    "true",
	}

	resp, err := req.Do(ctx, t.Client)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.IsError() {
		fmt.Println("Error in indexing", err, resp.Status())
		return err
	}
	fmt.Println("Indexing successful")

	io.Copy(io.Discard, resp.Body) //nolint: errcheck

	return nil
}

// Delete removes a task from the index.
func (t *ElasticDB) Delete(ctx context.Context, id string, indexname string) error {

	req := esv8api.DeleteRequest{
		Index:      indexname,
		DocumentID: id,
	}

	resp, err := req.Do(ctx, t.Client)
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
func (t *ElasticDB) Search(ctx context.Context, args domain.SearchParams, indexname string) (domain.SearchResults, error) {

	if args.IsZero() {
		return domain.SearchResults{}, nil
	}

	should := make([]interface{}, 0, 6)

	if args.Name != nil {
		should = append(should, map[string]interface{}{
			"match": map[string]interface{}{
				"name": *args.Name,
			},
		})
	}

	if args.Address != nil {
		should = append(should, map[string]interface{}{
			"match": map[string]interface{}{
				"address": *args.Address,
			},
		})
	}

	if args.Country != nil {
		should = append(should, map[string]interface{}{
			"match": map[string]interface{}{
				"country": *args.Country,
			},
		})
	}

	if args.DOB != nil {
		should = append(should, map[string]interface{}{
			"match": map[string]interface{}{
				"dob": *args.DOB,
			},
		})
	}

	if args.Gender != nil {
		should = append(should, map[string]interface{}{
			"match": map[string]interface{}{
				"gender": *args.Gender,
			},
		})
	}

	if args.UserID != nil {
		should = append(should, map[string]interface{}{
			"match": map[string]interface{}{
				"userid": *args.UserID,
			},
		})
	}

	var query map[string]interface{}

	if len(should) > 1 {
		query = map[string]interface{}{
			"query": map[string]interface{}{
				"bool": map[string]interface{}{
					"should": should,
				},
			},
		}
	} else {
		query = map[string]interface{}{
			"query": should[0],
		}
	}

	query["sort"] = []interface{}{
		"_score",
		map[string]interface{}{"id": "asc"},
	}

	query["from"] = args.From
	query["size"] = args.Size

	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return domain.SearchResults{}, err
	}

	req := esv8api.SearchRequest{
		Index: []string{indexname},
		Body:  &buf,
	}

	resp, err := req.Do(ctx, t.Client)
	if err != nil {
		return domain.SearchResults{}, err
	}
	defer resp.Body.Close()

	if resp.IsError() {
		fmt.Println("Error in searching", err, resp.Status())
		return domain.SearchResults{}, err
	}

	//nolint: tagliatelle
	var hits struct {
		Hits struct {
			Total struct {
				Value int64 `json:"value"`
			} `json:"total"`
			Hits []struct {
				Source domain.Task `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&hits); err != nil {
		return domain.SearchResults{}, err
	}

	res := make([]domain.Task, len(hits.Hits.Hits))

	for i, hit := range hits.Hits.Hits {
		res[i].ID = hit.Source.ID
		res[i].Name = hit.Source.Name
		res[i].Address = hit.Source.Address
		res[i].DOB = hit.Source.DOB
		res[i].Country = hit.Source.Country
		res[i].Gender = hit.Source.Gender
		res[i].UserID = hit.Source.UserID
	}

	return domain.SearchResults{
		Tasks: res,
		Total: hits.Hits.Total.Value,
	}, nil
}
