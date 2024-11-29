package logger

import (
	"log/slog"
	"os"
)

var Logger *slog.Logger

func LogInit(level string) {
	var lvl slog.Level
	switch level {
	case "levelInfo":
		lvl = slog.LevelInfo
	case "levelWarn":
		lvl = slog.LevelWarn
	case "levelErr":
		lvl = slog.LevelError
	default:
		lvl = slog.LevelDebug
	}

	file, err := os.OpenFile("product-service.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     lvl,
	}

	handler := slog.NewJSONHandler(file, opts)
	Logger = slog.New(handler)
	slog.SetDefault(Logger)
}
