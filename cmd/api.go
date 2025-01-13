package cmd

import (
	"context"
	"github.com/99designs/gqlgen/handler"
	"github.com/GabsMeloTI/go-graphql/infra"
	"github.com/GabsMeloTI/go-graphql/internal/generated"
	"github.com/GabsMeloTI/go-graphql/internal/resolvers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"time"
)

func StartAPI(ctx context.Context, container *infra.ContainerDI) {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.OPTIONS},
	}))

	go func() {
		for {
			select {
			case <-ctx.Done():
				if err := e.Shutdown(ctx); err != nil {
					panic(err)
				}
				return
			default:
				time.Sleep(1 * time.Second)
			}
		}
	}()

	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
	e.POST("/query", echo.WrapHandler(handler.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{DB: container.ConnDB}}))))

	e.Logger.Fatal(e.Start(container.Config.ServerPort))
}
