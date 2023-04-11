package renameAll

import "testing"

func TestReName(t *testing.T) {
	src := "/mnt/e/container/Music"
	pattern := "aac;mp3;wma;wav;flac;ogg"
	level := "Debug"
	replace(src, pattern, level, "..", ".")
}
