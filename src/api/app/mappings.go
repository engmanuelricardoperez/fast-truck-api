package app

import (
	"github.com/engmanuelricardoperez/fast-truck-api/src/api/infrastructure/dependencies"
	"github.com/go-chi/chi/v5"
)

func ConfigureMappings(route chi.Route, handlers *dependencies.HandlerContainer) {
	configureAPIMappings(route, handlers)
}

func configureAPIMappings(route chi.Route, handlers *dependencies.HandlerContainer) {
	route.Get("/truck", handlers.GetTruck.Handler)
}
