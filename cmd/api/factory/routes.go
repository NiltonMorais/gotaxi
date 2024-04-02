package factory

import "net/http"

type Routes struct {
	HttpServer HttpServer
}

func NewRoutes(httpServer HttpServer) *Routes {
	return &Routes{
		HttpServer: httpServer,
	}
}

func (r *Routes) RegisterRoutes() {
	r.HttpServer.Register("POST", "/ride/request", r.requestRide)
}

func (route *Routes) requestRide(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Request Ride"))
}
