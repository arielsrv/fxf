package main

import (
	"github.com/arielsrv/fxf/internal/features/messages/commands"
	"github.com/arielsrv/fxf/internal/features/messages/delivery/http"
	"github.com/arielsrv/fxf/internal/features/messages/queries"
	"github.com/arielsrv/fxf/internal/features/messages/repository"
	"github.com/arielsrv/fxf/pkg/fiber"
	"github.com/arielsrv/fxf/pkg/mediator"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		// Pkg Modules
		fiber.Module,
		mediator.Module,

		// Feature Modules
		repository.Module,
		commands.Module,
		queries.Module,
		http.Module,
	)

	app.Run()
}
