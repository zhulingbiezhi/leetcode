package interview

import (
	"bytes"
	"context"
	"encoding/json"
	"hellotalk/logger"
	"io/ioutil"
	"strings"
	"time"
)

type Server struct {
	httpAddress string
	//数据流
	dataChan   chan string
	resultChan chan<- string
}

func getServerList(qps int, resultChan chan<- string) []*Server {
	serverAddrs := []string{
		//"http://45.25.13.64:8080",
		"http://127.0.0.1:8080",
		"http://52.221.50.69:8080",
	}
	servers := make([]*Server, 0, len(serverAddrs))
	for _, addr := range serverAddrs {
		servers = append(servers, &Server{
			httpAddress: addr,
			dataChan:    make(chan string, qps),
			resultChan:  resultChan,
		})
	}
	return servers
}

func (s *Server) Start(ctx context.Context, i int, token string) {
	logger.Infof("服务器 %d: address: %s, start...", i, s.httpAddress)
	defer func() {
		logger.Infof("服务器 %d:  address: %s, finished", i, s.httpAddress)
	}()
	for {
		keys := make([]string, 0)
		for {
			select {
			case key := <-s.dataChan:
				keys = append(keys, key)
			case <-ctx.Done():
				return
			}
			if len(keys) >= cap(s.dataChan){
				break
			}
		}
		if ret, ok := s.DealWithKeys(keys, token); ok {
			logger.Infof("服务器 %d:  address: %s, find result: %s", i, s.httpAddress, ret)
			s.resultChan <- ret
		}
		time.Sleep(time.Second)
	}
}

type TryKeysRequest struct {
	Keys  string `json:"keys,omitempty"`
	Token string `json:"token,omitempty"`
}

type TryKeysResponse struct {
	Result string `json:"result,omitempty"`
}

func (s *Server) DealWithKeys(keys []string, token string) (string, bool) {
	logger.Info("服务器开始处理数据", s.httpAddress, len(keys))
	urlStr := s.httpAddress + "/interview/try_keys"
	req := &TryKeysRequest{
		Keys:  strings.Join(keys, ","),
		Token: token,
	}
	b, _ := json.Marshal(req)
	resp, err := defaultHttpClient.Post(urlStr, "application/json", bytes.NewReader(b))
	if err != nil {
		logger.Error("Post err", err)
		return "", false
	}
	defer resp.Body.Close()
	bb, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("ioutil.ReadAll err", err)
		return "", false
	}
	ret := &TryKeysResponse{}
	err = json.Unmarshal(bb, ret)
	if err != nil {
		logger.Error("json.Unmarshal err", err)
		return "", false
	}
	if ret.Result != "" {
		logger.Info("DealWithKeys 找到匹配数据", ret.Result)
		return ret.Result, true
	}
	return "", false
}
