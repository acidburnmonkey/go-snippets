// Usage
//
// logger := slog.New(&pkg.CustomHandler{Level: slog.LevelDebug})
// slog.SetDefault(logger)
// Output : 2026/06/06 16:42:42 DEBUG | Message | main()
package main

import (
	"context"
	"fmt"
	"log/slog"
	"runtime"
	"strings"
)

type CustomHandler struct {
	level slog.Level
}

func (h *CustomHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.level
}

func (h *CustomHandler) Handle(_ context.Context, r slog.Record) error {
	fn := "unknown"
	pc := make([]uintptr, 10)
	n := runtime.Callers(5, pc)
	frames := runtime.CallersFrames(pc[:n])
	for {
		frame, more := frames.Next()
		if !strings.Contains(frame.Function, "slog") {
			parts := strings.Split(frame.Function, ".")
			fn = parts[len(parts)-1]
			break
		}
		if !more {
			break
		}
	}

	t := r.Time.Format("2006/01/02 15:04:05")
	fmt.Printf("%s %s | %s | %s()\n", t, r.Level, r.Message, fn)
	return nil
}

func (h *CustomHandler) WithAttrs(attrs []slog.Attr) slog.Handler { return h }
func (h *CustomHandler) WithGroup(name string) slog.Handler       { return h }
