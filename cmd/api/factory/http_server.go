package factory

import (
	"context"
	"net/http"
)

type HttpServer interface {
	Register(method, path string, callback func(w http.ResponseWriter, r *http.Request))
	Listen(ctx context.Context, port int) error
}
