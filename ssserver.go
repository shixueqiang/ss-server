package main

//简单的JSON Restful API演示(服务端)
//author: Xiong Chuan Liang
//date: 2015-2-28

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Item struct {
	Seq    int
	Result map[string]int
}

type Message struct {
	Dept    string
	Subject string
	Time    int64
	Detail  []Item
}

func getJson() ([]byte, error) {
	pass := make(map[string]int)
	pass["x"] = 50
	pass["c"] = 60
	item1 := Item{100, pass}

	reject := make(map[string]int)
	reject["l"] = 11
	reject["d"] = 20
	item2 := Item{200, reject}

	detail := []Item{item1, item2}
	m := Message{"IT", "KPI", time.Now().Unix(), detail}
	return json.MarshalIndent(m, "", "")
}

func handler(w http.ResponseWriter, r *http.Request) {
	resp, err := getJson()
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, string(resp))
}

// func main() {
// 	http.HandleFunc("/", handler)
// 	http.ListenAndServe("localhost:8085", nil)
// }
