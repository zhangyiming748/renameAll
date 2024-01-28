package main

import (
	"fmt"
	"github.com/zhangyiming748/GetFileInfo"
	"github.com/zhangyiming748/renameAll/reg"
	"io"
	"log/slog"
	"os"
	"strings"
)

func init() {
	SetLogLevel("Info")
}
func main() {
	root := "/Users/zen/Documents/柒柒音声/music"
	pattern := "mp3"
	replace(root, pattern)
	slog.Info("end")
}
func replace(root, pattern string) {
	files := GetFileInfo.GetAllFileInfo(root, pattern)
	for _, file := range files {
		slog.Debug(fmt.Sprintf("文件信息: %+v", file))
		if has, substring := reg.F1(file.PurgeName); has {
			slog.Info(fmt.Sprintf("前缀名 : %s", substring))
			nName := strings.Replace(file.PurgeName, substring, "", 1)
			oldFull := strings.Join([]string{file.PurgePath, file.PurgeName, ".", file.PurgeExt}, "")
			newFull := strings.Join([]string{file.PurgePath, nName, ".", file.PurgeExt}, "")
			slog.Info("summarize", slog.String("旧文件名", oldFull), slog.String("新文件名", newFull))
			err := os.Rename(oldFull, newFull)
			if err != nil {
				continue
			}
		}
	}

}

/*
设置程序运行的日志等级
*/
func SetLogLevel(s string) {
	file := "rename.log"
	logf, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		panic(err)
	}
	var opt slog.HandlerOptions
	switch s {
	case "Debug", "debug":
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelDebug, // slog 默认日志级别是 info
		}
	case "Info", "info":
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: false,
			Level:     slog.LevelInfo, // slog 默认日志级别是 info
		}
	case "Warn", "warn":
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelWarn, // slog 默认日志级别是 info
		}
	case "Err", "err":
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelError, // slog 默认日志级别是 info
		}
	default:
		slog.Warn("需要正确设置环境变量 Debug,Info,Warn or Err")
		slog.Debug("默认使用Debug等级")
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelDebug, // slog 默认日志级别是 info
		}
	}
	logger := slog.New(slog.NewJSONHandler(io.MultiWriter(logf, os.Stdout), &opt))
	slog.SetDefault(logger)
}
