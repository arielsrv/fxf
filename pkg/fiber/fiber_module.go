package fiber

import (
	"context"
	"log/slog"

	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/contrib/otelfiber"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

// Module exports the fiber server functionality.
var Module = fx.Options(
	fx.Provide(NewFiberServer),
	fx.Invoke(func(lc fx.Lifecycle, app *fiber.App) {
		lc.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				slog.InfoContext(ctx, "Starting Fiber server on :3000")
				// The server is started in a goroutine so that it doesn't
				// block the application from starting.
				go func() {
					if err := app.Listen(":3000"); err != nil {
						slog.ErrorContext(ctx, err.Error())
					}
				}()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				slog.InfoContext(ctx, "Stopping Fiber server")
				return app.ShutdownWithContext(ctx)
			},
		})
	}),
)

// NewFiberServer creates a new Fiber server instance.
func NewFiberServer() *fiber.App {
	app := fiber.New(fiber.Config{
		EnablePrintRoutes:     true,
		DisableStartupMessage: true,
	})

	app.Use(otelfiber.Middleware())

	prometheus := fiberprometheus.NewWithDefaultRegistry("fxf")
	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Middleware)

	return app
}
