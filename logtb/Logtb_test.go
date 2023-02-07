package logtb

import (
	"os"
	"path/filepath"
	"testing"
)

// テストで使用するログファイルパスを設定（相対パス）
const testLogFilePath string = "../utlog/ut.log"

func TestSetLogFilePath(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name            string
		args            args
		wantLogFilePath string
	}{
		// 正常系：ログファイルパスを設定する
		{"SetLogFilePath正常系１", args{testLogFilePath}, testLogFilePath},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetLogFilePath(tt.args.filePath)

			if logFilePath != testLogFilePath {
				t.Errorf("createLogFilePath() = %v, want %v", logFilePath, tt.wantLogFilePath)
			}
		})
	}
}

func TestGetLogFilePath(t *testing.T) {
	tests := []struct {
		name            string
		wantLogFilePath string
	}{
		// 正常系：ログファイルパスの取得
		{"GetLogFilePath正常系１", testLogFilePath},
	}

	// ログファイルパスの設定
	logFilePath = testLogFilePath

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLogFilePath(); got != tt.wantLogFilePath {
				t.Errorf("GetLogFilePath() = %v, want %v", got, tt.wantLogFilePath)
			}
		})
	}
}

func TestLogger(t *testing.T) {

	logInfo := NewLogInfo(LogLevelEnum(INFO), "処理を開始します")

	type args struct {
		level     LogInfo
		argsArray []string
	}
	tests := []struct {
		name string
		args args
	}{

		// 正常系：ディレクトリ作成、ログ出力
		{"Logger正常系１", args{logInfo, []string{}}},
	}

	// ログディレクトリ作成処理実施のため、ログディレクトリを削除
	deleteLogDirectory(filepath.Dir(testLogFilePath))

	logFilePath = testLogFilePath

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Logger(tt.args.level, tt.args.argsArray)

			if fi, err := os.Stat(filepath.Dir(testLogFilePath)); os.IsNotExist(err) {
				t.Error("Logger() = ディレクトリ作成失敗")
			} else if fi.Size() == 0 {
				t.Error("Logger() = ログファイル書き込み失敗")
			}
		})
	}

	// 最終テストのため、ログディレクトリを削除
	deleteLogDirectory(filepath.Dir(testLogFilePath))
}

// ログディレクトリを削除
func deleteLogDirectory(logFile string) {

	os.RemoveAll(logFile)

}
