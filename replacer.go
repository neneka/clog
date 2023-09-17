package clog

import "log/slog"

// replaceLevel replaces level with severity.
// cf. https://cs.opensource.google/go/go/+/refs/tags/go1.21.1:src/log/slog/handler.go;l=279
func replaceLevel(a slog.Attr) slog.Attr {
	if a.Key == slog.LevelKey {
		l := a.Value.Any().(slog.Level)
		a.Key = "severity"
		a.Value = slog.StringValue(severityString(l))
	}

	return a
}

// replaceMessage replaces message key.
// cf. https://cs.opensource.google/go/go/+/refs/tags/go1.21.1:src/log/slog/handler.go;l=292
func replaceMessage(a slog.Attr) slog.Attr {
	if a.Key == slog.MessageKey {
		a.Key = "message"
	}

	return a
}
