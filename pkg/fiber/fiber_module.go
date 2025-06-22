package fiber

import (
	"context"
	"log"
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

// Module exports the fiber server functionality.
var Module = fx.Options(
	fx.Provide(NewFiberServer),
	fx.Invoke(func(lc fx.Lifecycle, app *fiber.App) {
		lc.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				slog.Info("Starting Fiber server on :3000")
				// The server is started in a goroutine so that it doesn't
				// block the application from starting.
				go func() {
					if err := app.Listen(":3000"); err != nil {
						panic(err)
					}
				}()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				log.Println("Stopping Fiber server")
				return app.Shutdown()
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
	return app
}
