package cutil

import (
	"strings"

	"github.com/gorilla/mux"
)

type RouteDescription struct {
	Name    string `json:"name"`
	Methods string `json:"methods,omitempty"`
	Path    string `json:"path,omitempty"`
}

type RouteDescriptions = []*RouteDescription

func ExtractRoutes(r *mux.Router) RouteDescriptions {
	var ret RouteDescriptions

	_ = r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		methods, _ := route.GetMethods()
		pathTemplate, _ := route.GetPathTemplate()
		name := route.GetName()
		m := strings.Join(methods, ", ")
		ret = append(ret, &RouteDescription{name, m, pathTemplate})
		return nil
	})

	return ret
}
