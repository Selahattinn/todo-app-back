package job

import "github.com/Selahattinn/todo-app-back/pkg/model"

type Reader interface {
	GetJobs() ([]model.Job, error)
}

type Writer interface {
	StoreJob(job model.Job) (int64, error)
}

//Repository repository interface
type Repository interface {
	Reader
	Writer
}
