package main

import (
	"garu-io-frontend-api/api"
	"garu-io-frontend-api/repository"
	"github.com/opentracing/opentracing-go"
	"io"
	"log"
	"time"

	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

func main() {
	tracer, clsTracer, err := initTracing()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := clsTracer.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	workApi := api.NewWorkApi(repository.LocalWorkExperiencesRepository{})
	languageApi := api.NewLanguagesApi(repository.LocalLanguagesRepository{}, tracer)
	projectsApi := api.NewProjectsApi(repository.NewProjectsApi("http://garu-io-projects-api.garu-io-projects-api:80", tracer), tracer)
	apiServer := api.NewApi(workApi, languageApi, projectsApi, tracer)

	err = apiServer.Serve(api.ServeOpts{
		Addr: "0.0.0.0",
		Port: 8080,
	})

	if err != nil {
		log.Fatal(err)
	}
}

func initTracing() (opentracing.Tracer, io.Closer, error) {
	conf, err := config.FromEnv()
	if err != nil {
		return nil, nil, err
	}

	conf.ServiceName = "garu-io-frontend-api"

	conf.Sampler.Type = "const"
	conf.Sampler.Param = 1

	conf.Reporter.LogSpans = true
	conf.Reporter.BufferFlushInterval = 10 * time.Second

	return conf.NewTracer(config.Logger(jaeger.StdLogger))
}
