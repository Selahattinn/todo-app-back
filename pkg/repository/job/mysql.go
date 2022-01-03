package job

import (
	"database/sql"
	"fmt"

	"github.com/Selahattinn/todo-app-back/pkg/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

type MySQLRepository struct {
	db *sql.DB
}

const (
	tableName = "jobs"
)

const (
	initTableTemplate = `
	CREATE TABLE IF NOT EXISTS %s (
		id bigint(20) NOT NULL AUTO_INCREMENT PRIMARY KEY,
		body varchar(256) NOT NULL DEFAULT '',
		completed bigint(20) NOT NULL DEFAULT 0,
		UNIQUE KEY id (id)
	  ) ENGINE=MyISAM  DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;	
`
)

func NewMySQLRepository(db *sql.DB) (*MySQLRepository, error) {
	tableInitCmd := fmt.Sprintf(initTableTemplate, tableName)
	_, err := db.Exec(tableInitCmd)

	if err != nil {
		return nil, fmt.Errorf("error init job repository: %v", err)
	}

	return &MySQLRepository{
		db: db,
	}, nil
}

func (r *MySQLRepository) GetJobs(id int64, item_type string) ([]model.Job, error) {
	q := "SELECT id, body, completed FROM " + tableName

	logrus.Debug("QUERY: ", q)
	res, err := r.db.Query(q)
	if err != nil {
		return nil, fmt.Errorf("error init job repository: %v", err)
	}
	var jobs []model.Job
	for res.Next() {
		var job model.Job
		if err := res.Scan(&job.ID, &job.Body, &job.Completed); err != nil {
			return nil, err
		}
		jobs = append(jobs, job)

	}

	return jobs, nil
}
