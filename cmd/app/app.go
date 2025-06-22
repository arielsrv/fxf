package main

import (
	"github.com/arielsrv/fxf/internal/features/messages/commands"
	"github.com/arielsrv/fxf/internal/features/messages/delivery/http"
	"github.com/arielsrv/fxf/internal/features/messages/queries"
	"github.com/arielsrv/fxf/internal/features/messages/repository"
	"github.com/arielsrv/fxf/internal/features/messages/service"
	"github.com/arielsrv/fxf/pkg/fiber"
	"github.com/arielsrv/fxf/pkg/logger"
	"github.com/arielsrv/fxf/pkg/mediator"
	"github.com/arielsrv/fxf/pkg/telemetry"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		// Pkg Modules
		fiber.Module,
		mediator.Module,
		logger.Module,
		telemetry.Module,

		// Feature Modules
		repository.Module,
		commands.Module,
		queries.Module,
		service.Module,
		http.Module,
	)

	app.Run()
}
