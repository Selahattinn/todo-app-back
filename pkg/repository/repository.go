package repository

import "github.com/Selahattinn/todo-app-back/pkg/repository/job"

// Repository defines the method for all operations related with repository
// Repository interface is composition of  Repository interfaces of imported packages.
type Repository interface {
	Shutdown()
	GetJobRepository() job.Repository
}
