package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Selahattinn/todo-app-back/pkg/api/response"
	"github.com/Selahattinn/todo-app-back/pkg/model"
)

//endpoint for get userInformation
func (a *API) GetJobs(w http.ResponseWriter, r *http.Request) {
	jobs, err := a.service.GetJobService().GetJobs()
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting storeUserInformation info: %v", err), http.StatusBadRequest, err.Error())
		return
	}
	response.Write(w, r, jobs)
}

//endpoint for get userInformation
func (a *API) StoreJobs(w http.ResponseWriter, r *http.Request) {
	var payload model.Job
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting store job info: %v", err), http.StatusBadRequest, err.Error())
		return
	}
	if payload.Body == "" {
		response.Errorf(w, r, fmt.Errorf("error getting store job info: %v", err), http.StatusBadRequest, err.Error())
		return
	}
	id, err := a.service.GetJobService().StoreJob(payload)
	if err != nil {
		fmt.Println(err)
		response.Errorf(w, r, fmt.Errorf("error getting storeUserInformation info: %v", err), http.StatusBadRequest, err.Error())
		return
	}
	var job model.Job
	job.ID = id
	response.Write(w, r, job)
}
