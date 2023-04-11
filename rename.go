package renameAll

import (
	"github.com/zhangyiming748/GetAllFolder"
	"github.com/zhangyiming748/GetFileInfo"
	"golang.org/x/exp/slog"
	"os"
	"strings"
)

func replace(src, pattern, level string, older, newer string) {
	folders := GetAllFolder.ListFolders(src, level)
	folders = append(folders, src)
	for _, folder := range folders {
		files := GetFileInfo.GetAllFileInfo(folder, pattern, level)
		for _, file := range files {
			slog.Info("file info", slog.String("FullPath", file.FullPath), slog.String("FullName", file.FullName), slog.String("ExtName", file.ExtName), slog.Int64("FileSize", file.Size))
			if strings.Contains(file.FullName, older) {
				perfix := strings.Trim(file.FullPath, file.FullName) // 带最后一个路径分隔符
				slog.Info("前缀", slog.String("path-name", perfix))
				newName := strings.Replace(file.FullName, older, newer, -1)
				newFull := strings.Join([]string{perfix, newName}, "")
				slog.Info("重命名之前", slog.String("旧全名", file.FullPath), slog.String("新全名", newFull))
				err := os.Rename(file.FullPath, newFull)
				if err != nil {
					slog.Warn("重命名出错")
					return
				}
			}
		}
	}
}
