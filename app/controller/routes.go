package controller

import (
	"net/http"

	"github.com/gorilla/mux"
)

func BuildRouter() (*mux.Router, error) {
	r := mux.NewRouter()

	// Home
	r.Path("/").Methods(http.MethodGet).Handler(http.HandlerFunc(Home)).Name("home")

	// Assets
	_ = r.Path("/assets").Subrouter()
	r.Path("/favicon.ico").Methods(http.MethodGet).Handler(http.HandlerFunc(Favicon)).Name("favicon")
	r.Path("/robots.txt").Methods(http.MethodGet).Handler(http.HandlerFunc(RobotsTxt)).Name("robots")
	r.PathPrefix("/assets").Methods(http.MethodGet).Handler(http.HandlerFunc(Static)).Name("assets")

	return r, nil
}
