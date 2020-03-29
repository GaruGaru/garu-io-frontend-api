package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Api struct {
	work     WorkApi
	language LanguagesApi
}

func NewApi(workApi WorkApi, languagesApi LanguagesApi) Api {
	return Api{
		work:     workApi,
		language: languagesApi,
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
	router.HandleFunc("/languages", a.language.LanguagesList).Methods("GET")

	return http.ListenAndServe(fmt.Sprintf("%s:%d", opts.Addr, opts.Port), router)
}
