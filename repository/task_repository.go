package repository

import (
	"context"

	"github.com/RajathSVasisth/elasticApp/domain"
	"github.com/RajathSVasisth/elasticApp/elastic"
)

// Task represents the repository used for interacting with Task records.
type taskRepository struct {
	elasticClient *elastic.ElasticDB
	index         string
}

func NewTaskRepository(elasticClient *elastic.ElasticDB, indexname string) domain.TaskRepository {
	return &taskRepository{
		elasticClient: elasticClient,
		index:         indexname,
	}
}

func (tr *taskRepository) Create(c context.Context, task *domain.Task) error {

	err := tr.elasticClient.Index(c, *task, tr.index)

	return err
}

func (tr *taskRepository) FetchByUserID(c context.Context, userID string) ([]domain.Task, error) {

	res, err := tr.elasticClient.Search(c, domain.SearchParams{UserID: &userID}, tr.index)
	if err != nil {
		return nil, err
	}

	return res.Tasks, err
}
