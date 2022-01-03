package service

import "github.com/Selahattinn/todo-app-back/pkg/service/job"

type Config struct{}

type Service interface {
	GetConfig() *Config
	Shutdown()
	GetJobService() *job.Service
}
