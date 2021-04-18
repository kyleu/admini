package controller

import (
	"net/http"

	"github.com/gorilla/mux"
)

func BuildRouter() (*mux.Router, error) {
	r := mux.NewRouter()

	// Home
	r.Path("/").Methods(http.MethodGet).Handler(http.HandlerFunc(Home)).Name("home")

	r.Path("/search").Methods(http.MethodGet).Handler(http.HandlerFunc(Search)).Name("search")
	r.Path("/profile").Methods(http.MethodGet).Handler(http.HandlerFunc(Profile)).Name("profile")
	r.Path("/settings").Methods(http.MethodGet).Handler(http.HandlerFunc(Settings)).Name("settings")
	r.Path("/feedback").Methods(http.MethodGet).Handler(http.HandlerFunc(Feedback)).Name("feedback")
	r.Path("/help").Methods(http.MethodGet).Handler(http.HandlerFunc(Help)).Name("help")

	// Sandbox
	r.Path("/sandbox").Methods(http.MethodGet).Handler(http.HandlerFunc(SandboxList)).Name("sandbox.list")
	r.Path("/sandbox/{key}").Methods(http.MethodGet).Handler(http.HandlerFunc(SandboxRun)).Name("sandbox.run")

	// Source
	r.Path("/source").Methods(http.MethodGet).Handler(http.HandlerFunc(SourceList)).Name("source.list")
	r.Path("/source/{key}").Methods(http.MethodGet).Handler(http.HandlerFunc(SourceDetail)).Name("source.detail")

	// Util
	r.Path("/modules").Methods(http.MethodGet).Handler(http.HandlerFunc(Modules)).Name("modules")
	r.Path("/routes").Methods(http.MethodGet).Handler(http.HandlerFunc(Routes)).Name("routes")

	// Assets
	_ = r.Path("/assets").Subrouter()
	r.Path("/sitemap.xml").Methods(http.MethodGet).Handler(http.HandlerFunc(SitemapXML)).Name("sitemap")
	r.Path("/favicon.ico").Methods(http.MethodGet).Handler(http.HandlerFunc(Favicon)).Name("favicon")
	r.Path("/robots.txt").Methods(http.MethodGet).Handler(http.HandlerFunc(RobotsTxt)).Name("robots")
	r.PathPrefix("/assets").Methods(http.MethodGet).Handler(http.HandlerFunc(Static)).Name("assets")

	r.PathPrefix("").Methods(http.MethodOptions).Handler(http.HandlerFunc(Options))
	r.PathPrefix("").Handler(http.HandlerFunc(NotFound))

	return r, nil
}
