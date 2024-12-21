package main

import (
	"encoding/json"
	"fmt"
	"github.com/mazezen/itools"
	"log"
	"net/http"
)

type RequestPrams struct {
	Name string `json:"name" validate:"required"`
	Age  int    `json:"age" validate:"required,number"`
}

func get(w http.ResponseWriter, r *http.Request) {
	var params RequestPrams
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if msg := itools.EnValidateParam(params); msg != "" {
		fmt.Fprintf(w, "msg: %s", msg)
		return
	}

	fmt.Fprintf(w, "name: %s, age: %d", params.Name, params.Age)
}

func get2(w http.ResponseWriter, r *http.Request) {
	var params RequestPrams
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if msg := itools.ZhValidateParam(params); msg != "" {
		fmt.Fprintf(w, "msg: %s", msg)
		return
	}

	fmt.Fprintf(w, "name: %s, age: %d", params.Name, params.Age)
}

func main() {

	http.HandleFunc("/get", get)
	http.HandleFunc("/get2", get2)

	fmt.Println("start server port :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
