package job

import (
	"github.com/Selahattinn/todo-app-back/pkg/model"
	"github.com/Selahattinn/todo-app-back/pkg/repository"
)

type Service struct {
	repository repository.Repository
}

func NewService(repo repository.Repository) (*Service, error) {
	return &Service{
		repository: repo,
	}, nil
}
func (s *Service) GetJobs() ([]model.Job, error) {
	jobs, err := s.repository.GetJobRepository().GetJobs()
	if err != nil {
		return nil, nil
	}
	return jobs, nil
}

func (s *Service) StoreJob(job model.Job) (int64, error) {
	id, err := s.repository.GetJobRepository().StoreJob(job)
	if err != nil {
		return id, nil
	}
	return id, nil
}
