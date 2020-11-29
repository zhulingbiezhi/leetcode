package interview

import (
	"context"
	"encoding/json"
	"fmt"
	"hellotalk/utils"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"hellotalk/algorithm"
	"hellotalk/logger"
)

const ServerQPS = 500
const MaxGoroutine = 100

var defaultHttpClient = &http.Client{
	Timeout: time.Second * 5,
}

func DistributedTasks(key string, token string) {
	now := time.Now()
	defer func() {
		logger.Infof("DistributedTasks finished, 耗时 %d ms", time.Since(now)/time.Millisecond)
	}()
	logger.Info("DistributedTasks start, key: ", key)

	ctx, cancel := context.WithCancel(context.Background())
	resultChan := make(chan string)
	serverList := getServerList(ServerQPS, resultChan)
	srvLen := len(serverList)
	produceQPS := ServerQPS * srvLen

	a := &algorithm.Algorithm{
		QPS: int64(produceQPS),
		Ctx: ctx,
	}
	a.Init()

	//producer，固定速率每秒500*srvLen
	go utils.HandlePanicGo(func() {
		a.Permutation(key)
	})

	//consumer,均匀分配key到各个server
	go utils.HandlePanicGo(func() {
		keyCount := 0
		moveCount := 0
		tt := time.Now()
		t1 := tt
		t2 := tt

	loop:
		for {
			select {
			case key, ok := <-a.ConsumersChan():
				if !ok {
					break loop
				}
				keyCount++
				select {
				case serverList[keyCount%srvLen].dataChan <- key:
				default: //如果某一个server阻塞了，找下一个server
					//logger.Error("服务器阻塞，http_address:", serverList[keyCount%srvLen].httpAddress)
					moveCount++
				}
				if (keyCount-moveCount)%produceQPS == 0 {
					st := time.Since(t1) / time.Millisecond
					if st > 5 {
						logger.Infof("key数量: %d, 阻塞key转移数量：%d, 耗时: %d ms", keyCount-moveCount, moveCount)
					}
					t1 = time.Now()
				}
			}
		}
		logger.Infof("produce完成 key数量：%d, 阻塞key转移数量：%d，耗时: %d ms", keyCount-moveCount, moveCount, time.Since(t2)/time.Millisecond)
	})
	//start tasks
	wg := &sync.WaitGroup{}
	wg.Add(srvLen)

	for i := 0; i < len(serverList); i++ {
		go func(i int) {
			defer wg.Done()
			server := serverList[i]
			server.Start(ctx, i, token)
		}(i)
	}
	select {
	case result := <-resultChan:
		logger.Info("DistributedTasks find result", result)
		cancel()
	}
	wg.Wait()
}

type taskResponse struct {
	Info string `json:"info,omitempty"`
}

func DealWithTask(keys string, token string) (string, error) {
	logger.Info("DealWithTask start")
	result := ""
	totalCnt := 0
	now := time.Now()

	dataChan := make(chan string, ServerQPS/MaxGoroutine)

	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(MaxGoroutine)
	for i := 0; i < MaxGoroutine; i++ {
		go func() {
			//logger.Info("协程开始运行")
			//t := time.Now()
			//count := 0
			defer wg.Done()
		loop:
			for {
				select {
				case t, ok := <-dataChan:
					if !ok {
						break loop
					}
					totalCnt++
					ret, err := executeTask(t, token)
					if err != nil {
						logger.Error("executeTask err", err)
						dataChan <- t
					}
					if ret != "" {
						logger.Info("协程找到匹配数据", result)
						cancel()
						result = ret
					}
					time.Sleep(time.Microsecond * 20)
				}
			}
			//logger.Infof("协程运行完成，耗时 %d ms, 处理数据: %d", time.Since(t)/time.Millisecond, count)
		}()
	}

	ks := strings.Split(keys, ",")

	//logger.Info("主协程开始发送数据")
	for _, key := range ks {
		select {
		case dataChan <- key:
		case <-ctx.Done():
			break
		}
	}
	close(dataChan)
	wg.Wait()
	logger.Infof("task 处理完成，耗时: %d ms, 处理数据: %d", time.Since(now)/time.Millisecond, totalCnt)
	return result, nil
}

func executeTask(key, token string) (string, error) {
	//todo remove test code
	//if key == algorithm.TestKey {
	//	return algorithm.TestKey, nil
	//} else {
	//	//time.Sleep(time.Millisecond)
	//	//logger.Info("executeTask",key)
	//	return "", nil
	//}
	urlStr := "http://interview.hellotalk8.com"
	vals := url.Values{
		"token": []string{token},
		"key":   []string{key},
	}

	resp, err := defaultHttpClient.PostForm(urlStr, vals)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case http.StatusOK:
		return key, nil
	case http.StatusBadRequest:
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		ret := &taskResponse{}
		if err = json.Unmarshal(b, ret); err != nil {
			return "", err
		}
		if ret.Info != "无效token" {
			logger.Error("发现新错误：", ret.Info)
		}
		return "", nil
	default:
		return "", fmt.Errorf("发现未知code: %d", resp.StatusCode)
	}
}
