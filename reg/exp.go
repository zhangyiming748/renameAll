package reg

import (
	"fmt"
	"log/slog"
	"regexp"
)

/*
匹配一个字母和三个数字连续出现的字符串
*/
func F1(str string) (bool, string) {
	re := regexp.MustCompile(`[a-zA-Z]\d{3}`)
	//str := "a123 b456 c789"
	matches := re.FindAllString(str, -1)
	for _, match := range matches {
		fmt.Println(match) // 输出 a123 b456 c789
	}
	if len(matches) == 0 {
		slog.Debug("no match")
		return false, ""
	} else {
		slog.Debug("match")
		return true, matches[0]
	}
}

/*
匹配一个数字和英文句号连续出现的字符串
*/
func F2(str string) (bool, string) {
	re := regexp.MustCompile(`\d{2}\.`)
	//str := "a123 b456 c789"
	matches := re.FindAllString(str, -1)
	for _, match := range matches {
		fmt.Println(match) // 输出 a123 b456 c789
	}
	if len(matches) == 0 {
		slog.Debug("no match")
		return false, ""
	} else {
		slog.Debug("match")
		return true, matches[0]
	}
}
