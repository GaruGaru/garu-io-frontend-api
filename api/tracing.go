package api

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"net/http"
)

const (
	MiddlewareContextSpanKey = "span"
)

func spanFromRequest(tracer opentracing.Tracer, r *http.Request, opName string) opentracing.Span {
	clientContext, isSpan := r.Context().Value(MiddlewareContextSpanKey).(opentracing.SpanContext)
	opts := make([]opentracing.StartSpanOption, 0)
	if isSpan {
		opts = append(opts, ext.RPCServerOption(clientContext))
	}
	return tracer.StartSpan(opName, opts...)
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

		contextWithSpan := context.WithValue(r.Context(), MiddlewareContextSpanKey, span)
		r.WithContext(contextWithSpan)

		next.ServeHTTP(w, r)

		span.SetTag("http.method", r.Method)

		// TODO add response writer interceptor
		if r.Response != nil {
			span.SetTag("http.status_code", r.Response.StatusCode)
		}

	})
}
