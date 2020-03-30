package api

import (
	"encoding/json"
	"garu-io-frontend-api/repository"
	"github.com/sirupsen/logrus"
	"net/http"
)

type LanguagesApi struct {
	repository repository.LanguagesRepository
}

func NewLanguagesApi(repository repository.LanguagesRepository) LanguagesApi {
	return LanguagesApi{repository: repository}
}

func (a LanguagesApi) LanguagesList(w http.ResponseWriter, r *http.Request) {
	response, err := a.repository.KnownLanguages(repository.KnownLanguagesRequest{})
	w.Header().Add("Content-Type", "application/json")

	if err != nil {
		logrus.Error("languages api repository error", err)
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
