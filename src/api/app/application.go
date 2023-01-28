package app

import (
	"context"
	"log"
	"net/http"

	"github.com/engmanuelricardoperez/fast-truck-api/src/api/infrastructure/dependencies"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Start() {
	ctx := context.Background()
	println("App Starting :", ctx)
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	log := log.Default()

	logEnviroment(ctx, *log)

	dep := dependencies.Dependencies{}
	handlers := dep.Start()
	ConfigureMappings(router, handlers)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello GolinuxCloud members!"))
	})

	http.ListenAndServe(":8080", router)
}

func logEnviroment(ctx context.Context, log log.Logger) {
	log.Println(ctx, "staring Truck API")
}
