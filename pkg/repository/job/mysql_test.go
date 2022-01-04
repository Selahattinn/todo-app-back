package job

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/Selahattinn/todo-app-back/pkg/model"
	_ "github.com/go-sql-driver/mysql"
)

func TestMySQLRepository_GetJobs(t *testing.T) {
	type fields struct {
		db *sql.DB
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
			r := &MySQLRepository{
				db: tt.fields.db,
			}
			got, err := r.GetJobs()
			if (err != nil) != tt.wantErr {
				t.Errorf("MySQLRepository.GetJobs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MySQLRepository.GetJobs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMySQLRepository(t *testing.T) {
	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name    string
		args    args
		want    *MySQLRepository
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewMySQLRepository(tt.args.db)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMySQLRepository() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMySQLRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}
