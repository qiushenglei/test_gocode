package myhttp

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PostData struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func GetStruct() {

	server := &http.Server{Addr: ":8080"}
	http.HandleFunc("/post", getStruct)
	server.ListenAndServe()
}

func getStruct(m http.ResponseWriter, w *http.Request) {
	var body PostData
	json.NewDecoder(w.Body).Decode(&body)
	//json.Unmarshal(w.Body, body)
	fmt.Println(body)
}
