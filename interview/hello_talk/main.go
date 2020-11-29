package main

import (
	"net/http"
	"os"

	"hellotalk/apis"
	_ "net/http/pprof"
)

func main() {
	http.HandleFunc("/interview/keys", apis.HandleGetKey)
	http.HandleFunc("/interview/try_keys", apis.HandleTryKeys)
	port, ok := os.LookupEnv("hello_port")
	if !ok {
		port = "8080"
	}
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}

}
