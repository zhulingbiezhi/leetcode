package utils

import (
	"runtime/debug"
	"strings"

	"github.com/zhulingbiezhi/leetcode/interview/hello_talk/logger"
)

func HandlePanicGo(f func()) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("[HandlePanicGo] err", err)
			stack := strings.Join(strings.Split(string(debug.Stack()), "\n")[2:], "\n")
			logger.Error("[HandlePanicGo] - stack:", stack)
			return
		}
	}()
	f()
}
