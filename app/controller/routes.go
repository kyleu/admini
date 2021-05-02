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

	help := r.Path("/help").Subrouter()
	help.Methods(http.MethodGet).Handler(wrap(Help)).Name("help")
	r.Path("/feedback").Methods(http.MethodGet).Handler(wrap(Feedback)).Name("feedback")

	// Source
	source := r.Path("/source").Subrouter()
	source.Methods(http.MethodGet).Handler(wrap(SourceList)).Name("source.list")
	r.Path("/source/{key}").Methods(http.MethodGet).Handler(wrap(SourceDetail)).Name("source.detail")
	r.Path("/source/{key}/refresh").Methods(http.MethodGet).Handler(wrap(SourceRefresh)).Name("source.refresh")

	// Project
	project := r.Path("/project").Subrouter()
	project.Methods(http.MethodGet).Handler(wrap(ProjectList)).Name("project.list")
	r.Path("/project/{key}").Methods(http.MethodGet).Handler(wrap(ProjectDetail)).Name("project.detail")

	// Workspace
	_ = r.PathPrefix("/x/{key}").Handler(wrap(WorkspaceProject)).Name("workspace")
	_ = r.PathPrefix("/s/{key}").Handler(wrap(WorkspaceSource)).Name("workspace.source")

	// Sandbox
	sandbox := r.Path("/sandbox").Subrouter()
	sandbox.Methods(http.MethodGet).Handler(wrap(SandboxList)).Name("sandbox.list")
	r.Path("/sandbox/{key}").Methods(http.MethodGet).Handler(wrap(SandboxRun)).Name("sandbox.run")

	// Util
	_ = r.Path("/util").Subrouter()
	r.Path("/util/modules").Methods(http.MethodGet).Handler(wrap(Modules)).Name("modules")
	r.Path("/util/routes").Methods(http.MethodGet).Handler(wrap(Routes)).Name("routes")

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
