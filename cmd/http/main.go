package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	http.HandleFunc("/beat", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello world1"))
	})
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello world2"))
	})
	http.HandleFunc("/123", func(writer http.ResponseWriter, request *http.Request) {
		i := 0
		for {
			i++
			if i > 10 {
				break
			}
			time.Sleep(1 * time.Second)
		}
		fmt.Println("finish")
		writer.Write([]byte("hello world2 finish"))
	})
	http.ListenAndServe(":10012", nil)
}
