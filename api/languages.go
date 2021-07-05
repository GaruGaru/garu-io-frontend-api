package api

import (
	"encoding/json"
	"garu-io-frontend-api/repository"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/sirupsen/logrus"
	"net/http"
)

type LanguagesApi struct {
	repository repository.LanguagesRepository
	tracer     opentracing.Tracer
}

func NewLanguagesApi(repository repository.LanguagesRepository, tracer opentracing.Tracer) LanguagesApi {
	return LanguagesApi{repository: repository, tracer: tracer}
}

func (a LanguagesApi) LanguagesList(w http.ResponseWriter, r *http.Request) {

	span := spanFromRequest(a.tracer, r, "LanguagesList")
	defer span.Finish()

	response, err := a.repository.KnownLanguages(repository.KnownLanguagesRequest{})
	w.Header().Add("Content-Type", "application/json")

	if err != nil {
		logrus.Error("languages api repository error", err)

		span.SetTag("error", true)
		span.LogFields(log.Error(err))

		if _, err := w.Write([]byte(err.Error())); err != nil {
			logrus.Error("unable to write error response", err)
		}
	}

	err = json.NewEncoder(w).Encode(response.Results)

	if err != nil {
		logrus.Error("unable to write api response", err)

		span.SetTag("error", true)
		span.LogFields(log.Error(err))

		if _, err := w.Write([]byte(err.Error())); err != nil {
			logrus.Error("unable to write error response", err)
		}
	}
}
