package controller

import (
	"net/http"

	"github.com/gorilla/mux"
)

func BuildRouter() (*mux.Router, error) {
	r := mux.NewRouter()

	// Home
	r.Path("/").Methods(http.MethodGet).Handler(http.HandlerFunc(Home)).Name("home")
	r.Path("/settings").Methods(http.MethodGet).Handler(http.HandlerFunc(Settings)).Name("settings")
	r.Path("/feedback").Methods(http.MethodGet).Handler(http.HandlerFunc(Feedback)).Name("feedback")
	r.Path("/help").Methods(http.MethodGet).Handler(http.HandlerFunc(Help)).Name("help")

	r.Path("/sandbox").Methods(http.MethodGet).Handler(http.HandlerFunc(SandboxList)).Name("sandbox.list")
	r.Path("/sandbox/{key}").Methods(http.MethodGet).Handler(http.HandlerFunc(SandboxRun)).Name("sandbox.run")

	// Assets
	_ = r.Path("/assets").Subrouter()
	r.Path("/favicon.ico").Methods(http.MethodGet).Handler(http.HandlerFunc(Favicon)).Name("favicon")
	r.Path("/robots.txt").Methods(http.MethodGet).Handler(http.HandlerFunc(RobotsTxt)).Name("robots")
	r.PathPrefix("/assets").Methods(http.MethodGet).Handler(http.HandlerFunc(Static)).Name("assets")

	return r, nil
}
