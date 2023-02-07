package logtb

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/buildkite/interpolate"
)

// ログファイルに実行ログを出力する
type Logtb string

var logFilePath string

// ログファイルパスを設定する
func SetLogFilePath(filePath string) {

	logFilePath = filePath

}

// ログファイルパスを取得する
func GetLogFilePath() string {

	return logFilePath

}

// ログ出力処理
func Logger(logInfo LogInfo, argsArray []string) {

	const MESSAGE_FORMAT string = "${timestamp} [${level}] "
	const TIMESTAMP_FORMAT string = "2006-01-02 15:04:05.000"

	// ログディレクトリがアプリケーション起動中に削除された時の対処
	if fi, err := os.Stat(filepath.Dir(logFilePath)); os.IsNotExist(err) || fi != nil && !fi.IsDir() {

		os.MkdirAll(filepath.Dir(logFilePath), 0755)

	}

	file, _ := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer file.Close()

	argsArray = append(argsArray, "level="+string(logInfo.level), "timestamp="+time.Now().Format(TIMESTAMP_FORMAT))

	env := interpolate.NewSliceEnv(argsArray)

	fixedMessage, _ := interpolate.Interpolate(env, MESSAGE_FORMAT+string(logInfo.message))

	// ログをファイルへ出力
	fmt.Fprintln(file, fixedMessage)

}
