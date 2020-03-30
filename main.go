package main

import (
	"garu-io-frontend-api/api"
	"garu-io-frontend-api/repository"
	"log"
)

func main() {
	workApi := api.NewWorkApi(repository.LocalWorkExperiencesRepository{})
	languageApi := api.NewLanguagesApi(repository.LocalLanguagesRepository{})
	apiServer := api.NewApi(workApi, languageApi)

	err := apiServer.Serve(api.ServeOpts{
		Addr: "0.0.0.0",
		Port: 8080,
	})

	if err != nil{
		log.Fatal(err)
	}
}
