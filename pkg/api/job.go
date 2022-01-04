package api

import (
	"fmt"
	"net/http"

	"github.com/Selahattinn/todo-app-back/pkg/api/response"
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
