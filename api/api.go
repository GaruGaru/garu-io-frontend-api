package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Api struct {
	work     WorkApi
	language LanguagesApi
	projects ProjectsApi
}

func NewApi(workApi WorkApi, languagesApi LanguagesApi, projectsApi ProjectsApi) Api {
	return Api{
		work:     workApi,
		language: languagesApi,
		projects: projectsApi,
	}
}

type ServeOpts struct {
	Addr string
	Port int
}

func (a Api) Serve(opts ServeOpts) error {
	router := mux.NewRouter()

	router.HandleFunc("/", Home).Methods("GET")
	router.HandleFunc("/work/experiences", a.work.WorkList).Methods("GET")
	router.HandleFunc("/personal/projects", a.projects.ProjectsList).Methods("GET")
	router.HandleFunc("/languages", a.language.LanguagesList).Methods("GET")

	return http.ListenAndServe(fmt.Sprintf("%s:%d", opts.Addr, opts.Port), router)
}
