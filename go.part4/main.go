// // package main

// // import (
// // 	"fmt"
// // )

// // const url string = "https://lco.dev:3456/me?name=hc&course=go"

// // func main() {
// // 	msg := "hey godev!🦥"
// // 	fmt.Println(msg)
// // 	// if 1 > 2 {
// // 	// 	fmt.Println(url)
// // 	// 	response, err := http.Get(url)
// // 	// 	checkNilErr(err)
// // 	// 	fmt.Printf("response: %T\n", response)  // *http.Response type pointer
// // 	// 	fmt.Println("response:", response.Body) //
// // 	// 	defer response.Body.Close()             // close the response body when done using defer
// // 	// 	resData, err := io.ReadAll(response.Body)
// // 	// 	checkNilErr(err)
// // 	// 	fmt.Println("resData:", resData)         // byte slice
// // 	// 	fmt.Println("resData:", string(resData)) // convert byte slice to string
// // 	// }

// // 	// s := "postgres://user:pass@host.com:5432/path?k=v#f"
// // 	// // url parsing
// // 	// u, err := url.Parse(s)
// // 	// if err != nil {
// // 	// 	panic(err)
// // 	// }
// // 	// // Accessing the scheme is straightforward.

// // 	// fmt.Println(u.Scheme)

// // }

// // func checkNilErr(err error) {
// // 	if err != nil {
// // 		fmt.Println("error:", err)
// // 		panic(err)
// // 	}
// // }

// package main

// import (
// 	"fmt"
// 	"net/url"
// )

// const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

// func main() {
// 	fmt.Println("Welcome to handling URLs in golang")
// 	fmt.Println(myurl)

// 	//parsing
// 	result, _ := url.Parse(myurl)

// 	// fmt.Println(result.Scheme)
// 	// fmt.Println(result.Host)
// 	// fmt.Println(result.Path)
// 	// fmt.Println(result.Port())
// 	fmt.Println(result.RawQuery)

// 	qparams := result.Query()
// 	fmt.Printf("The type of query params are: %T\n", qparams)

// 	fmt.Println(qparams["coursename"])

// 	for _, val := range qparams {
// 		fmt.Println("Param is: ", val)
// 	}

// 	partsOfUrl := &url.URL{
// 		Scheme:  "https",
// 		Host:    "lco.dev",
// 		Path:    "/tutcss",
// 		RawPath: "user=hitesh",
// 	}

// 	anotherURL := partsOfUrl.String()
// 	fmt.Println(anotherURL)

// }

package main

import (
	"fmt"
	"net/url"
)

const testURL string = "http://localhost:4000/dc?name=dc&course=go&course=js"

func main() {
	msg := "hello user...)"
	if 1 > 2 {
		fmt.Println(msg, testURL)
	}
	// url parsing
	reqURL, err := url.Parse(testURL)
	checkNilErr(err)
	fmt.Println("req URL host", reqURL.Host)     // localhost:4000
	fmt.Println("req URL scheme", reqURL.Scheme) // http
	fmt.Println("req URL path", reqURL.Path)     // /dc
	fmt.Println("req URL query", reqURL.Query()) // name=dc&course=go -> its a map
	reqQuery := reqURL.Query()
	for key, val := range reqQuery {
		fmt.Println("key:", key, ", val:", val)
	}
	partsOfUrl := &url.URL{
		Scheme:  "http",
		Host:    "localhost:4000",
		Path:    "/dc",
		RawPath: "name=dc&course=go&course=js",
	}
	resultedURL := partsOfUrl.String()
	fmt.Println("partsOfUrl:", partsOfUrl)
	fmt.Println("resultedURL:", resultedURL)
}

func checkNilErr(err error) {
	if err != nil {
		fmt.Println("error:", err)
		panic(err)
	}
}
