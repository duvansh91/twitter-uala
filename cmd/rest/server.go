package rest

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type api struct {
	router http.Handler
}

type Server interface {
	Router() http.Handler
}

type Route struct {
	Path    string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}

func (a *api) Router() http.Handler {
	return a.router
}

func NewSever(routes []Route) Server {
	a := &api{}
	r := mux.NewRouter()

	for _, route := range routes {
		r.HandleFunc(route.Path, route.Handler).Methods(route.Method)
	}

	a.router = r
	return a
}

func MergePath(base, route string) string {
	return fmt.Sprintf("%s/%s", base, route)
}
