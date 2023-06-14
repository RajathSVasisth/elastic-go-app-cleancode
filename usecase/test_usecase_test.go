package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/RajathSVasisth/elasticApp/domain"
	"github.com/RajathSVasisth/elasticApp/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreateTask(t *testing.T) {
	type args struct {
		task *domain.Task
	}

	testCases := []struct {
		name        string
		taskUsecase *taskUsecase
		args        args
		wantErr     error
	}{
		{
			name: "fail - db error for Create",
			taskUsecase: func() *taskUsecase {
				taskCreateRepository := new(mocks.TaskRepository)
				taskCreateRepository.On("Create", mock.Anything, mock.Anything).Return(errors.New("db error"))
				return &taskUsecase{
					taskRepository: taskCreateRepository,
					contextTimeout: 0,
				}
			}(),
			args: args{
				task: &domain.Task{
					ID:      primitive.NewObjectID(),
					UserID:  primitive.NewObjectID(),
					Name:    "test",
					DOB:     "test",
					Address: "test",
					Gender:  "test",
					Country: "test",
				},
			},
			wantErr: errors.New("db error"),
		},
		{name: "success",
			taskUsecase: func() *taskUsecase {
				taskCreateRepository := new(mocks.TaskRepository)
				taskCreateRepository.On("Create", mock.Anything, mock.Anything).Return(nil)
				return &taskUsecase{
					taskRepository: taskCreateRepository,
					contextTimeout: 0,
				}
			}(),
			args: args{
				task: &domain.Task{
					ID:      primitive.NewObjectID(),
					UserID:  primitive.NewObjectID(),
					Name:    "test",
					DOB:     "test",
					Address: "test",
					Gender:  "test",
					Country: "test",
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.taskUsecase.Create(context.Background(), tt.args.task)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestGetTask(t *testing.T) {
	type args struct {
		pagination domain.Pagination
		userID     string
	}
	commonID := primitive.NewObjectID()

	testCases := []struct {
		name        string
		taskUsecase *taskUsecase
		args        args
		want        []domain.Task
		wantErr     error
	}{
		{
			name: "fail - db error for FetchByUserID",
			taskUsecase: func() *taskUsecase {
				taskGetRepository := new(mocks.TaskRepository)
				taskGetRepository.On("FetchByUserID", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("db error"))
				return &taskUsecase{
					taskRepository: taskGetRepository,
					contextTimeout: 0,
				}
			}(),
			args: args{
				pagination: domain.Pagination{
					From: 0,
					Size: 10,
				},
				userID: primitive.NewObjectID().Hex(),
			},
			want:    nil,
			wantErr: errors.New("db error"),
		},
		{
			name: "success",
			taskUsecase: func() *taskUsecase {
				taskGetRepository := new(mocks.TaskRepository)
				taskGetRepository.On("FetchByUserID", mock.Anything, mock.Anything, mock.Anything).Return([]domain.Task{
					{
						ID:      commonID,
						UserID:  commonID,
						Name:    "test",
						DOB:     "test",
						Address: "test",
						Gender:  "test",
						Country: "test",
					},
				}, nil)
				return &taskUsecase{
					taskRepository: taskGetRepository,
					contextTimeout: 0,
				}
			}(),
			args: args{
				pagination: domain.Pagination{
					From: 0,
					Size: 10,
				},
				userID: primitive.NewObjectID().Hex(),
			},
			want: []domain.Task{
				{
					ID:      commonID,
					UserID:  commonID,
					Name:    "test",
					DOB:     "test",
					Address: "test",
					Gender:  "test",
					Country: "test",
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.taskUsecase.FetchByUserID(context.Background(), tt.args.userID, tt.args.pagination)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
