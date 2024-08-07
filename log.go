package auth

type Log interface {
	Info(msg string, data ...any)
	Warn(msg string, data ...any)
	Error(msg string, data ...any)
}
