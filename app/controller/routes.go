package controller

import (
	"net/http"

	"github.com/gorilla/handlers"

	"github.com/gorilla/mux"
)

func wrap(f http.HandlerFunc) http.Handler {
	return handlers.CompressHandler(f)
}

func BuildRouter() (*mux.Router, error) {
	r := mux.NewRouter()

	// Home
	r.Path("/").Methods(http.MethodGet).Handler(wrap(Home)).Name("home")

	r.Path("/search").Methods(http.MethodGet).Handler(wrap(Search)).Name("search")
	r.Path("/profile").Methods(http.MethodGet).Handler(wrap(Profile)).Name("profile")
	r.Path("/settings").Methods(http.MethodGet).Handler(wrap(Settings)).Name("settings")
	r.Path("/feedback").Methods(http.MethodGet).Handler(wrap(Feedback)).Name("feedback")
	r.Path("/help").Methods(http.MethodGet).Handler(wrap(Help)).Name("help")

	// Sandbox
	r.Path("/sandbox").Methods(http.MethodGet).Handler(wrap(SandboxList)).Name("sandbox.list")
	r.Path("/sandbox/{key}").Methods(http.MethodGet).Handler(wrap(SandboxRun)).Name("sandbox.run")

	// Source
	r.Path("/source").Methods(http.MethodGet).Handler(wrap(SourceList)).Name("source.list")
	r.Path("/source/{key}").Methods(http.MethodGet).Handler(wrap(SourceDetail)).Name("source.detail")

	// Util
	r.Path("/modules").Methods(http.MethodGet).Handler(wrap(Modules)).Name("modules")
	r.Path("/routes").Methods(http.MethodGet).Handler(wrap(Routes)).Name("routes")

	// Assets
	_ = r.Path("/assets").Subrouter()
	r.Path("/sitemap.xml").Methods(http.MethodGet).Handler(http.HandlerFunc(SitemapXML)).Name("sitemap")
	r.Path("/favicon.ico").Methods(http.MethodGet).Handler(http.HandlerFunc(Favicon)).Name("favicon")
	r.Path("/robots.txt").Methods(http.MethodGet).Handler(http.HandlerFunc(RobotsTxt)).Name("robots")
	r.PathPrefix("/assets").Methods(http.MethodGet).Handler(http.HandlerFunc(Static)).Name("assets")

	r.PathPrefix("").Methods(http.MethodOptions).Handler(wrap(Options))
	r.PathPrefix("").Handler(wrap(NotFound))

	return r, nil
}
