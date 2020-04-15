package repository

import (
	"encoding/json"
	"fmt"
	"garu-io-frontend-api/models"
	"net/http"
	"time"
)

type ProjectsApiResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Language    string `json:"language"`
}

type ProjectsApi struct {
	endpoint string
	client   *http.Client
}

func NewProjectsApi(endpoint string) ProjectsApi {
	return ProjectsApi{
		endpoint: endpoint,
		client:   &http.Client{Timeout: 5 * time.Second},
	}
}

func (p ProjectsApi) Projects(request ProjectsRequest) (ProjectsResponse, error) {
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
