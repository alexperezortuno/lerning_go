package api

import (
	"../data"
	"./handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"path"
)

var routes = data.Routes{
	data.Route{
		Name:       "Index",
		Method:     "GET",
		Pattern:    "/",
		HandleFunc: handlers.IndexHandler,
	},
	data.Route{
		Name:       "User",
		Method:     "GET",
		Pattern:    "/user",
		HandleFunc: handlers.UserHandler,
	},
}

func NewRoutes() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	base, _ := os.Getwd()
	// Serve static files
	sf := http.FileServer(http.Dir(path.Join(base, "web/static")))
	mf := http.FileServer(http.Dir(path.Join(base, "web/static/media")))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", sf))
	r.PathPrefix("/media/").Handler(http.StripPrefix("/media/", mf))

	for _, route := range routes {
		r.
			Name(route.Name).
			Methods(route.Method).
			Path(route.Pattern).
			Handler(route.HandleFunc)
	}

	return r
}
