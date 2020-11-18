package apis

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/zhulingbiezhi/leetcode/interview/hello_talk/interview"
	"github.com/zhulingbiezhi/leetcode/interview/hello_talk/logger"
	"github.com/zhulingbiezhi/leetcode/interview/hello_talk/utils"
)

type GetKeyResponse struct {
	Description []string `json:"description,omitempty"`
	Key         string   `json:"key,omitempty"`
	Token       string   `json:"token,omitempty"`
}

func HandleGetKey(w http.ResponseWriter, r *http.Request) {
	logger.Info("[HandleGetKey] start get key")
	urlStr := "http://interview.hellotalk8.com"
	resp, err := http.DefaultClient.Get(urlStr)
	if err != nil {
		logger.Error("[HandleGetKey] request err")
		ResponseServerError(w, "request err")
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("[HandleGetKey] read body err", err)
		ResponseServerError(w, "read body err")
		return
	}
	ret := &GetKeyResponse{}
	err = json.Unmarshal(b, ret)
	if err != nil {
		logger.Error("[HandleGetKey] unmarshal body err", err)
		ResponseServerError(w, "unmarshal body err")
		return
	}
	go utils.HandlePanicGo(func() {
		interview.DistributedTasks()
	})
	logger.Info("[HandleGetKey] get key success, response ", string(b))
	ResponseSuccess(w, "success")
}

type TryKeysRequest struct {
}

func HandleTryKeys(w http.ResponseWriter, r *http.Request) {
	logger.Info("[HandleGetKey] start try key")
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Error("[HandleGetKey] read body err", err)
		ResponseServerError(w, "read body err")
		return
	}
	logger.Info("[HandleGetKey] request body", string(b))
	defer r.Body.Close()
	ret := &TryKeysRequest{}
	err = json.Unmarshal(b, ret)
	if err != nil {
		logger.Error("[HandleGetKey] unmarshal request body err", err)
		ResponseClientError(w, "unmarshal request body err")
		return
	}
	go utils.HandlePanicGo(func() {

	})
}

func ResponseServerError(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusInternalServerError)
	_, err := w.Write([]byte(msg))
	if err != nil {
		logger.Error("response write err ", err)
	}
}

func ResponseClientError(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusBadRequest)
	_, err := w.Write([]byte(msg))
	if err != nil {
		logger.Error("response write err ", err)
	}
}

func ResponseSuccess(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(msg))
	if err != nil {
		logger.Error("response write err ", err)
	}
}
