package main

import (
	"fmt"
	"io"
	"net/http"
)

const url string = "http://localhost:3000/me"

func HandleGetReq(url string) {
	resp, err := http.Get(url)
	CheckNilErr(err)
	defer resp.Body.Close()
	contest, err := io.ReadAll(resp.Body)
	CheckNilErr(err)
	finalResponse := string(contest)
	fmt.Println("finalResponse:", finalResponse)

}

type User struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Password string `json:"-"` // ignore the password field
}

func main() {
	fmt.Println("Welcome to handling URLs in golang")
	// handleGetReq()
	// handlePostReq()
	// HandleGetReq(url)
	// http.HandleFunc("/me", HandleMal)
// 	fmt.Println("url", url)
	
}

func HandleMal(w http.ResponseWriter, r *http.Request) {

}

// func handleGetReq() {
// 	result, err := http.Get(url)
// 	CheckNilErr(err)
// 	defer result.Body.Close() // close the response body when done using defer
// 	// fmt.Println("result:", result)
// 	// StatusCode -> 200
// 	// status -> 202 Accepted
// 	var status int = result.StatusCode
// 	if status == 200 || status == 201 {
// 		fmt.Println("status:", status)
// 		contest, err := io.ReadAll(result.Body)
// 		CheckNilErr(err)
// 		fmt.Println("contest:", string(contest))

// 	} else {
// 		// fmt.Println("status:", status)
// 		panic("status not < 202")
// 	}
// }

func CheckNilErr(err error) {
	if err != nil {
		fmt.Println("error:", err)
		panic(err)
	}
}

// func handlePostReq() {
// 	payload := strings.NewReader(`
// 	{
// 		"coursename":"Let's go with golang",
// 		"price": 0,
// 		"platform":"learnCodeOnline.in"
// 	}
// `)

// 	resp, err := http.Post(url, "application/json", payload)
// 	CheckNilErr(err)
// 	defer resp.Body.Close()
// 	fmt.Println("response status", resp.Status)
// 	content,err := io.ReadAll(resp.Body)
// 	CheckNilErr(err)
// 	fmt.Println("content:", string(content))

// }
