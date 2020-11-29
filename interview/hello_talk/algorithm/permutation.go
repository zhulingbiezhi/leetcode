package algorithm

import (
	"context"
	"hellotalk/logger"
)

var TestKey = "f1vn2x1B"

type Algorithm struct {
	//数据流
	permutationsChan chan string
	used             map[int]bool
	Ctx              context.Context
}

func (a *Algorithm) Init() {
	a.permutationsChan = make(chan string)
	a.used = make(map[int]bool)
}

func (a *Algorithm) Permutation(s string) {
	logger.Info("Permutation key")
	a.permutationString(s, "")
	close(a.permutationsChan)
	logger.Info("Permutation finish")
}

//回溯算法
func (a *Algorithm) permutationString(s string, ret string) {
	select {
	case <-a.Ctx.Done():
		return
	default:

	}
	if len(ret) == len(s) {
		a.permutationsChan <- ret
		return
	}
	m := make(map[uint8]bool)
	for i := 0; i < len(s); i++ {
		if a.used[i] {
			continue
		}
		//同一平行层，去除重复元素
		if m[s[i]] {
			continue
		}
		m[s[i]] = true

		a.used[i] = true
		a.permutationString(s, ret+string(s[i]))
		a.used[i] = false
	}
}

func (a *Algorithm) ConsumersChan() <-chan string {
	return a.permutationsChan
}
