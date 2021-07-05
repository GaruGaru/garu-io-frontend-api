package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/opentracing/opentracing-go"
	"net/http"
)

type Api struct {
	work     WorkApi
	language LanguagesApi
	projects ProjectsApi
	tracer   opentracing.Tracer
}

func NewApi(workApi WorkApi, languagesApi LanguagesApi, projectsApi ProjectsApi, tracer opentracing.Tracer) Api {
	return Api{
		work:     workApi,
		language: languagesApi,
		projects: projectsApi,
		tracer:   tracer,
	}
}

type ServeOpts struct {
	Addr string
	Port int
}

func (a Api) Serve(opts ServeOpts) error {
	router := mux.NewRouter()

	router.Use(a.tracingMiddleware)

	router.HandleFunc("/", Home).Methods("GET")
	router.HandleFunc("/work/experiences", a.work.WorkList).Methods("GET")
	router.HandleFunc("/personal/projects", a.projects.ProjectsList).Methods("GET")
	router.HandleFunc("/languages", a.language.LanguagesList).Methods("GET")

	return http.ListenAndServe(fmt.Sprintf("%s:%d", opts.Addr, opts.Port), router)
}
