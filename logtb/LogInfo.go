package logtb

// ログレベルとログメッセージを保持する構造体
type LogInfo struct {
	level   LogLevelEnum
	message string
}

func NewLogInfo(level LogLevelEnum, message string) LogInfo {
	enum := new(LogInfo)
	enum.level = level
	enum.message = message

	return *enum
}
