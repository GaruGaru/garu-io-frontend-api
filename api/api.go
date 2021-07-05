package api

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
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

func (a Api) tracingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		carrier := opentracing.HTTPHeadersCarrier(r.Header)
		clientContext, err := a.tracer.Extract(opentracing.HTTPHeaders, carrier)

		var methodName = r.URL.Path
		var span opentracing.Span
		if err == nil {
			span = a.tracer.StartSpan(methodName, ext.RPCServerOption(clientContext))
		} else {
			span = a.tracer.StartSpan(methodName)
		}

		defer span.Finish()

		contextWithSpan := context.WithValue(r.Context(), "span", span)
		r.WithContext(contextWithSpan)

		next.ServeHTTP(w, r)

		span.LogFields(log.String("http.method", r.Method))

		// TODO add response writer interceptor
		if r.Response != nil {
			span.LogFields(log.Int("http.status.code", r.Response.StatusCode))
			span.LogFields(log.String("http.status.message", r.Response.Status))
		}

	})
}
