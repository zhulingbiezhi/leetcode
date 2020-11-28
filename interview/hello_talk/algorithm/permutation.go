package algorithm

import (
	"context"
	"time"
)

var TestKey = "f1vn2x1B"

type Algorithm struct {
	//数据流
	permutationsChan chan string
	used             map[int]bool
	count            int64
	QPS              int64
	Ctx              context.Context
}

func (a *Algorithm) Init() {
	a.permutationsChan = make(chan string, a.QPS)
	a.used = make(map[int]bool)
}

func (a *Algorithm) Permutation(s string) {
	a.permutationString(s, "")
	close(a.permutationsChan)
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
		a.count++
		if a.count > 0 && a.count%a.QPS == 0 {
			time.Sleep(time.Second)
		}
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
