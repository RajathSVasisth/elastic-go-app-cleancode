package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"io"

	"github.com/RajathSVasisth/elasticApp/domain"
	"github.com/RajathSVasisth/elasticApp/elastic"
)

// Task represents the repository used for interacting with Task records.
type taskRepository struct {
	elasticClient elastic.ElasticDBMethod
	index         string
}

func NewTaskRepository(elasticClient elastic.ElasticDBMethod, indexname string) domain.TaskRepository {
	return &taskRepository{
		elasticClient: elasticClient,
		index:         indexname,
	}
}

func (tr *taskRepository) Create(c context.Context, task *domain.Task) error {

	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(task); err != nil {
		return err
	}
	err := tr.elasticClient.Index(c, buf, tr.index, task.ID.Hex())

	return err
}

func (tr *taskRepository) FetchByUserID(c context.Context, userID string) ([]domain.Task, error) {

	should := make([]interface{}, 0, 1)
	should = append(should, map[string]interface{}{
		"match": map[string]interface{}{
			"userid": userID,
		},
	})

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

	// query["from"] = args.From
	// query["size"] = args.Size

	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return []domain.Task{}, err
	}
	res, err := tr.elasticClient.Search(c, buf, tr.index)
	if err != nil {
		return []domain.Task{}, err
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

	if err := json.NewDecoder(res.(io.Reader)).Decode(&hits); err != nil {
		return []domain.Task{}, err
	}

	finalres := make([]domain.Task, len(hits.Hits.Hits))

	for i, hit := range hits.Hits.Hits {
		finalres[i].ID = hit.Source.ID
		finalres[i].Name = hit.Source.Name
		finalres[i].Address = hit.Source.Address
		finalres[i].DOB = hit.Source.DOB
		finalres[i].Country = hit.Source.Country
		finalres[i].Gender = hit.Source.Gender
		finalres[i].UserID = hit.Source.UserID
	}

	return finalres, err
}
