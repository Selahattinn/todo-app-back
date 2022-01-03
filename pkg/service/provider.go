package service

import (
	"github.com/Selahattinn/todo-app-back/pkg/repository"
	"github.com/Selahattinn/todo-app-back/pkg/service/job"
)

type Provider struct {
	cfg         *Config
	repository  repository.Repository
	JobsService *job.Service
}

func NewProvider(cfg *Config, repo repository.Repository) (*Provider, error) {
	JobService, err := job.NewService(repo)
	if err != nil {
		return nil, err
	}
	return &Provider{
		cfg:         cfg,
		repository:  repo,
		JobsService: JobService,
	}, nil
}
func (p *Provider) GetJobService() *job.Service {
	return p.JobsService
}
func (p *Provider) GetConfig() *Config {
	return p.cfg
}

func (p *Provider) Shutdown() {
	p.repository.Shutdown()
}
