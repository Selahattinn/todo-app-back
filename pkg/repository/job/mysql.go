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
		body TEXT NOT NULL,
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

func (r *MySQLRepository) GetJobs() ([]model.Job, error) {
	q := "SELECT id, body FROM " + tableName

	logrus.Debug("QUERY: ", q)
	res, err := r.db.Query(q)
	if err != nil {
		return nil, fmt.Errorf("error init job repository: %v", err)
	}
	var jobs []model.Job
	for res.Next() {
		var job model.Job
		if err := res.Scan(&job.ID, &job.Body); err != nil {
			return nil, err
		}
		jobs = append(jobs, job)

	}

	return jobs, nil
}

func (r *MySQLRepository) StoreJob(job model.Job) (int64, error) {
	stmt, err := r.db.Prepare(`INSERT INTO ` + tableName + `(
		body)
		VALUES(
			?)`)
	if err != nil {
		return -1, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(
		job.Body)
	if err != nil {
		return -1, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}
	return id, nil
}
