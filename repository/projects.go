package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"garu-io-frontend-api/models"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"net/http"
	"time"
)

type ProjectsApiResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Language    string `json:"language"`
	Uri         string `json:"uri"`
}

type ProjectsApi struct {
	endpoint string
	client   *http.Client
	tracer   opentracing.Tracer
}

func NewProjectsApi(endpoint string, tracer opentracing.Tracer) ProjectsApi {
	return ProjectsApi{
		endpoint: endpoint,
		client:   &http.Client{Timeout: 5 * time.Second},
		tracer:   tracer,
	}
}

func (p ProjectsApi) Projects(ctx context.Context, request ProjectsRequest) (ProjectsResponse, error) {
	parentSpan := opentracing.SpanFromContext(ctx)
	span := p.tracer.StartSpan("repository.Projects", ext.RPCServerOption(parentSpan.Context()))
	defer span.Finish()

	response, err := p.client.Get(fmt.Sprintf("%s/projects", p.endpoint))
	if err != nil {
		return ProjectsResponse{}, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return ProjectsResponse{}, fmt.Errorf("unexpected status code %d != 200", response.StatusCode)
	}

	decoder := json.NewDecoder(response.Body)

	var apiResponse []ProjectsApiResponse
	if err := decoder.Decode(&apiResponse); err != nil {
		return ProjectsResponse{}, err
	}

	itemsCount := min(request.Count, len(apiResponse))
	apiResponse = apiResponse[:itemsCount]

	projects := make([]models.Project, len(apiResponse))
	for i, r := range apiResponse {
		projects[i] = models.Project{
			Name:        r.Name,
			Description: r.Description,
			Language:    r.Language,
			Uri:         r.Uri,
		}
	}

	return ProjectsResponse{Results: projects}, nil
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
