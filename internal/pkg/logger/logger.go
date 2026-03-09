package logger

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

type PrettyHandler struct {
	slog.Handler
}

func (h *PrettyHandler) Handle(ctx context.Context, r slog.Record) error {
	level := r.Level.String()
	switch r.Level {
	case slog.LevelDebug:
		level = colorPurple + level + colorReset
	case slog.LevelInfo:
		level = colorCyan + level + colorReset
	case slog.LevelWarn:
		level = colorYellow + level + colorReset
	case slog.LevelError:
		level = colorRed + level + colorReset
	}

	timeStr := r.Time.Format("2006-01-02 15:04:05")
	msg := r.Message

	// Simple non-JSON output for console
	fmt.Printf("%s [%s] %s", timeStr, level, msg)

	r.Attrs(func(a slog.Attr) bool {
		fmt.Printf(" %s=%v", a.Key, a.Value)
		return true
	})

	fmt.Println()

	return nil
}

func Setup() {
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	// Use our pretty handler for development
	handler := &PrettyHandler{
		Handler: slog.NewJSONHandler(os.Stdout, opts),
	}
	logger := slog.New(handler)

	slog.SetDefault(logger)
}
