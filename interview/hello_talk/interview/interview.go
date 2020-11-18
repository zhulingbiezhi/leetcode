package interview

import (
	"time"

	"github.com/zhulingbiezhi/leetcode/interview/hello_talk/algorithm"
	"github.com/zhulingbiezhi/leetcode/interview/hello_talk/utils"
)

const MAX_TRY_COUNT = 500

func DistributedTasks(key string, token string) {
	a := &algorithm.Algorithm{}
	a.Permutation(key)
	count := 0
	serverList := getServerList()
	keys := make([]string, 0)

	for permutationKey := range a.ConsumersChan() {
		count++
		//每秒500并发
		if count%500 != 0 {
			keys = append(keys, permutationKey)
			continue
		}
		time.Sleep(time.Second)
		keys = make([]string, 0)
		go utils.HandlePanicGo()
	}
}

func getServerList() []string {
	return []string{
		"45.25.13.64:8080",
		"127.0.0.1:8080",
	}
}
