package api

import (
	"encoding/json"
	"garu-io-frontend-api/repository"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/sirupsen/logrus"
	"net/http"
)

type ProjectsApi struct {
	repository repository.ProjectsRepository
	tracer     opentracing.Tracer
}

func NewProjectsApi(repository repository.ProjectsRepository, tracer opentracing.Tracer) ProjectsApi {
	return ProjectsApi{repository: repository, tracer: tracer}
}

func (a ProjectsApi) ProjectsList(w http.ResponseWriter, r *http.Request) {
	span := spanFromRequest(a.tracer, r, "ProjectList")
	defer span.Finish()

	response, err := a.repository.Projects(repository.ProjectsRequest{
		Count: 5,
	})

	w.Header().Add("Content-Type", "application/json")

	if err != nil {
		logrus.Error("projects api repository error", err)

		span.SetTag("error", true)
		span.LogFields(log.Error(err))
		if _, err := w.Write([]byte(err.Error())); err != nil {
			logrus.Error("unable to write error response", err)
		}
	}

	span.LogFields(log.Int("projects_count", len(response.Results)))

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
