package logtb

// ログレベルを保持するEnum
type LogLevelEnum string

const (
	PANIC = "PANIC"
	FATAL = "FATAL"
	ERROR = "ERROR"
	WARN  = "WARN"
	INFO  = "INFO"
	DEBUG = "DEBUG"
	TRACE = "TRACE"
)
