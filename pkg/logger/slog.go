package logger

import (
	"log/slog"
	"os"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

var Module = fx.WithLogger(func() fxevent.Logger {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	logger := slog.New(handler)
	return New(logger)
})

// SlogLogger adapta slog a fxevent.Logger.
type SlogLogger struct {
	logger *slog.Logger
}

func New(logger *slog.Logger) fxevent.Logger {
	return &SlogLogger{logger: logger}
}

func (l *SlogLogger) LogEvent(e fxevent.Event) {
	switch evt := e.(type) {
	case *fxevent.OnStartExecuting:
		l.logger.Info("OnStart hook executing", "callee", evt.FunctionName, "caller", evt.CallerName)
	case *fxevent.OnStartExecuted:
		if evt.Err != nil {
			l.logger.Error("OnStart hook failed", "callee", evt.FunctionName, "caller", evt.CallerName, "err", evt.Err)
		} else {
			l.logger.Info("OnStart hook executed", "callee", evt.FunctionName, "caller", evt.CallerName, "runtime", evt.Runtime)
		}
	case *fxevent.OnStopExecuting:
		l.logger.Info("OnStop hook executing", "callee", evt.FunctionName, "caller", evt.CallerName)
	case *fxevent.OnStopExecuted:
		if evt.Err != nil {
			l.logger.Error("OnStop hook failed", "callee", evt.FunctionName, "caller", evt.CallerName, "err", evt.Err)
		} else {
			l.logger.Info("OnStop hook executed", "callee", evt.FunctionName, "caller", evt.CallerName, "runtime", evt.Runtime)
		}
	case *fxevent.Supplied:
		if evt.Err != nil {
			l.logger.Error("Supplied failed", "type", evt.TypeName, "err", evt.Err)
		} else {
			l.logger.Info("Supplied", "type", evt.TypeName)
		}
	case *fxevent.Provided:
		for _, rtype := range evt.OutputTypeNames {
			l.logger.Info("Provided", "constructor", evt.ConstructorName, "type", rtype)
		}
	case *fxevent.Decorated:
		for _, rtype := range evt.OutputTypeNames {
			l.logger.Info("Decorated", "decorator", evt.DecoratorName, "type", rtype)
		}
	case *fxevent.Invoking:
		l.logger.Info("Invoking", "function", evt.FunctionName)
	case *fxevent.Invoked:
		if evt.Err != nil {
			l.logger.Error("Invocation failed", "function", evt.FunctionName, "err", evt.Err)
		}
	case *fxevent.Started:
		if evt.Err != nil {
			l.logger.Error("Start failed", "err", evt.Err)
		} else {
			l.logger.Info("Started")
		}
	case *fxevent.Stopping:
		l.logger.Info("Stopping", "signal", evt.Signal.String())
	case *fxevent.Stopped:
		if evt.Err != nil {
			l.logger.Error("Stop failed", "err", evt.Err)
		} else {
			l.logger.Info("Stopped")
		}
	case *fxevent.RollingBack:
		l.logger.Error("Rolling back", "startErr", evt.StartErr)
	case *fxevent.RolledBack:
		l.logger.Error("Rolled back", "err", evt.Err)
	case *fxevent.LoggerInitialized:
		if evt.Err != nil {
			l.logger.Error("Custom logger initialization failed", "err", evt.Err)
		} else {
			l.logger.Info("Logger initialized", "constructor", evt.ConstructorName)
		}
	}
}
