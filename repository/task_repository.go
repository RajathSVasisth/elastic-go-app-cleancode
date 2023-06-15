package repository

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/RajathSVasisth/elastic-go-app-cleancode/domain"
	"github.com/RajathSVasisth/elastic-go-app-cleancode/elastic"
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

func (tr *taskRepository) FetchByUserID(c context.Context, userID string, pagination domain.Pagination) ([]domain.Task, error) {

	should := make([]interface{}, 0, 1)

	should = append(should, map[string]interface{}{
		"match": map[string]interface{}{
			"userid": userID,
		},
	})

	query := map[string]interface{}{
		"query": should[0],
	}

	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return []domain.Task{}, err
	}

	resBody, err := tr.elasticClient.Search(c, buf, tr.index, &pagination.From, &pagination.Size)
	if err != nil {
		return []domain.Task{}, err
	}

	defer resBody.Close()

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

	if err := json.NewDecoder(resBody).Decode(&hits); err != nil {
		return []domain.Task{}, err
	}

	finalres := make([]domain.Task, len(hits.Hits.Hits))

	for i, hit := range hits.Hits.Hits {
		finalres[i].Name = hit.Source.Name
		finalres[i].Address = hit.Source.Address
		finalres[i].DOB = hit.Source.DOB
		finalres[i].Country = hit.Source.Country
		finalres[i].Gender = hit.Source.Gender
		finalres[i].UserID = hit.Source.UserID
	}

	return finalres, err
}

func (tr *taskRepository) Update(c context.Context, task *domain.Task) error {

	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(task); err != nil {
		return err
	}

	err := tr.elasticClient.Index(c, buf, tr.index, task.ID.Hex())

	return err
}

func (tr *taskRepository) Delete(c context.Context, id *string) error {
	err := tr.elasticClient.Delete(c, *id, tr.index)
	return err
}
