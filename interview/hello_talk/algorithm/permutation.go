package algorithm

type Algorithm struct {
	permutationsChan chan string
	used             map[int]bool
}

func (a *Algorithm) Permutation(s string) {
	a.permutationsChan = make(chan string, 100)
	a.permutationString(s, "")
}

//回溯算法
func (a *Algorithm) permutationString(s string, ret string) {
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
