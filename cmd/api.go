package cmd

import (
	"context"
	"github.com/GabsMeloTI/go-graphql/infra"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"time"
)

func StartAPI(ctx context.Context, container *infra.ContainerDI) {
	e := echo.New()

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

	e.Logger.Fatal(e.Start(container.Config.ServerPort))
}
