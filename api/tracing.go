package api

import (
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"net/http"
)

func spanFromRequest(tracer opentracing.Tracer, r *http.Request, opName string) opentracing.Span {
	span := opentracing.SpanFromContext(r.Context())
	return tracer.StartSpan(opName, ext.RPCServerOption(span.Context()))
}

func (a Api) tracingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// skip health check tracing
		if r.URL.Path == "/" {
			next.ServeHTTP(w, r)
			return
		}

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

		ctxWithSpan := opentracing.ContextWithSpan(r.Context(), span)
		next.ServeHTTP(w, r.WithContext(ctxWithSpan))

		span.SetTag("http.method", r.Method)

		// TODO add response writer interceptor
		if r.Response != nil {
			span.SetTag("http.status_code", r.Response.StatusCode)
		}

	})
}
