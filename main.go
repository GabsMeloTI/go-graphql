package main

import (
	"context"
	"github.com/GabsMeloTI/go-graphql/cmd"
	"github.com/GabsMeloTI/go-graphql/infra"
	_ "github.com/GabsMeloTI/go-graphql/infra"
	_ "github.com/go-chi/chi/middleware"
	"os/signal"
	"syscall"
)

func main() {

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	defer stop()
	loadingEnv := infra.NewConfig()
	container := infra.NewContainerDI(loadingEnv)

	cmd.StartAPI(ctx, container)

}
