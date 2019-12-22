package middleware

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"log"
)

func LoggingMiddleware(next endpoint.Endpoint) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		log.Println("msg", "calling endpoint", ctx.Value(httptransport.ContextKeyRequestPath))
		defer log.Println("msg", "called endpoint",  ctx.Value(httptransport.ContextKeyRequestPath))
		return next(ctx, request)
	}
}
