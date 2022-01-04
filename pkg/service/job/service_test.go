package job

import (
	"reflect"
	"testing"

	"github.com/Selahattinn/todo-app-back/pkg/model"
	"github.com/Selahattinn/todo-app-back/pkg/repository"
)

func TestNewService(t *testing.T) {
	type args struct {
		repo repository.Repository
	}
	tests := []struct {
		name    string
		args    args
		want    *Service
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewService(tt.args.repo)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetJobs(t *testing.T) {
	type fields struct {
		repository repository.Repository
	}
	tests := []struct {
		name    string
		fields  fields
		want    []model.Job
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				repository: tt.fields.repository,
			}
			got, err := s.GetJobs()
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetUJobs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetUJobs() = %v, want %v", got, tt.want)
			}
		})
	}
}
