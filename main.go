package main

import (
	"context"
	"github.com/GabsMeloTI/go-graphql/cmd"
	"github.com/GabsMeloTI/go-graphql/infra"
	_ "github.com/go-chi/chi/middleware"
	"os/signal"
	"syscall"

	_ "github.com/GabsMeloTI/go-graphql/infra"
)

const defaultPort = "8080"

func main() {

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	defer stop()

	loadingEnv := infra.NewConfig()
	container := infra.NewContainerDI(loadingEnv)

	cmd.StartAPI(ctx, container)

	//port := os.Getenv("PORT")
	//if port == "" {
	//	port = defaultPort
	//}
	//
	//router := chi.NewRouter()
	//router.Use(cors.Handler(cors.Options{
	//	AllowedOrigins:   []string{"*"},
	//	AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
	//	AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
	//	AllowCredentials: true,
	//	ExposedHeaders:   []string{"Accept", "Content-Type"},
	//	MaxAge:           300,
	//}))
	//
	//srv := handler.New(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{}}))
	//
	//
	//router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	//router.Handle("/query", srv)
	//
	//log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	//log.Fatal(http.ListenAndServe(":"+port, router))
}
