package api

import (
	"encoding/json"
	"garu-io-frontend-api/repository"
	"github.com/sirupsen/logrus"
	"net/http"
)

type ProjectsApi struct {
	repository repository.ProjectsRepository
}

func NewProjectsApi(repository repository.ProjectsRepository) ProjectsApi {
	return ProjectsApi{repository: repository}
}

func (a ProjectsApi) ProjectsList(w http.ResponseWriter, r *http.Request) {
	response, err := a.repository.Projects(repository.ProjectsRequest{
		Count: 5,
	})
	w.Header().Add("Content-Type", "application/json")

	if err != nil {
		logrus.Error("projects api repository error", err)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			logrus.Error("unable to write error response", err)
		}
	}

	err = json.NewEncoder(w).Encode(response.Results)

	if err != nil {
		logrus.Error("unable to write api response", err)

		if _, err := w.Write([]byte(err.Error())); err != nil {
			logrus.Error("unable to write error response", err)
		}
	}
}
