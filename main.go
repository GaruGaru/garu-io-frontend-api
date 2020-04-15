package main

import (
	"garu-io-frontend-api/api"
	"garu-io-frontend-api/repository"
	"log"
)

func main() {
	workApi := api.NewWorkApi(repository.LocalWorkExperiencesRepository{})
	languageApi := api.NewLanguagesApi(repository.LocalLanguagesRepository{})
	projectsApi := api.NewProjectsApi(repository.NewProjectsApi("http://garu-io-projects-api.garu-io-projects-api:80"))

	apiServer := api.NewApi(workApi, languageApi, projectsApi)

	err := apiServer.Serve(api.ServeOpts{
		Addr: "0.0.0.0",
		Port: 8080,
	})

	if err != nil{
		log.Fatal(err)
	}
}
