package api

import (
	"encoding/json"
	"garu-io-api/repository"
	"github.com/sirupsen/logrus"
	"net/http"
)

type WorkApi struct {
	repository repository.WorkRepository
}

func NewWorkApi(repository repository.WorkRepository) WorkApi {
	return WorkApi{repository: repository}
}

func (a WorkApi) WorkList(w http.ResponseWriter, r *http.Request) {
	response, err := a.repository.WorkExperiences(repository.WorkExperienceRequest{})

	w.Header().Add("Content-Type", "application/json")

	if err != nil {
		logrus.Error("work api repository error", err)
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
