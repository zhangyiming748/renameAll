package main

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	SetLogLevel("Debug")
}
func main() {
	root := "/mnt/f/large/GirlFriend4ever"
	replace(root)
	slog.Info("end")
}
func replace(root string) {
	files, err := getFiles(root)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if strings.Contains(file, ".jpg.avif") {
			slog.Debug(fmt.Sprintf("文件信息: %+v", file))
			fresh := strings.Replace(file, ".jpg.avif", ".avif", 1)
			slog.Debug(fmt.Sprintf("新文件名: %+v", fresh))
			err := os.Rename(file, fresh)
			if err != nil {
				slog.Warn("重命名失败")
			} else {
				slog.Debug("重命名成功")
			}
		}
	}
}

func getFiles(folderPath string) ([]string, error) {
	var files []string

	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
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
