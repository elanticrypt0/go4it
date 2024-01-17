package go4it

import (
	"log"
	"log/slog"
	"os"
)

func newLog() *log.Logger {
	handler := slog.NewJSONHandler(os.Stdout, nil)
	log := slog.NewLogLogger(handler, slog.LevelDebug)
	return log
}
