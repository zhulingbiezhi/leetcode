package main

import (
	"net/http"

	"github.com/zhulingbiezhi/leetcode/interview/hello_talk/apis"
)

func main() {
	http.HandleFunc("/interview/keys", apis.HandleGetKey)
	http.HandleFunc("/interview/try_keys", apis.HandleTryKeys)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
