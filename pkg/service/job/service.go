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
func (s *Service) GetUJobs() ([]model.Job, error) {
	jobs, err := s.repository.GetJobRepository().GetJobs()
	if err != nil {
		return nil, nil
	}
	return jobs, nil
}
