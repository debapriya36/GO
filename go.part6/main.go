package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	msg := "welcome user!🦊"
	fmt.Println(msg)
	if 1 > 2 {
		jsonData := []byte(`
		{
			"type":"user",
			"name":"John",
			"age": 30,
			"email":"john@go.dev",
			"tags" : ["golang","coding"]
		}`)

		fmt.Println(jsonData)
		jsonDataMap := make(map[string]interface{})
		json.Unmarshal(jsonData, &jsonDataMap)
		fmt.Println("jsonDataMap:", jsonDataMap)
		for k, v := range jsonDataMap {
			fmt.Printf("key is %v and value is %v and type is %T\n", k, v, v)
		}
	}


	// routing using mux 
	r := mux.NewRouter() // r -> route
	// fmt.Println("r:", r)
	r.HandleFunc("/", serveHomeRoute).Methods("GET")

	log.Fatal(http.ListenAndServe(":3001", r))

}

func serveHomeRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Println("web request at http://localhost:3001/")
	w.Write([]byte("welcome to home route"))
}
