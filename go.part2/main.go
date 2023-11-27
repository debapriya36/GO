package main

import (
	"fmt"
	"math/rand"
)

func main() {

	msg := "Hey User🐦"
	fmt.Println(msg)
	// array

	var arr [2]int
	arr[0] = 1
	// arr[1] = "ok"
	if 1 > 2 {
		fmt.Println(arr)
	}

	// slice
	var arr1 = []int{}
	arr1 = append(arr1, 1)
	arr1 = append(arr1, 2)
	arr1 = append(arr1, 3)
	if 1 > 2 {
		fmt.Println("arr1:", arr1)
	}
	// // delete from slices
	// arr = append(arr[0:2])
	// fmt.Println("arr:",arr)
	//   str := []string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n"}
	//   fmt.Println("str:",str)
	//   lastIndex := len(str)-1
	//   str = append(str[0:lastIndex])
	//   fmt.Println("str:",str)

	//   arr := make([]int , 2)
	//   arr[0] = 10
	//   arr[1] = 2
	//   fmt.Println("arr:",arr)

	//   // sort the arr
	//   sort.Ints(arr)
	//   fmt.Println("arr:",arr)

	// str := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n"}
	// fmt.Println("str:", str)
	// // delete from slices
	// index := 7
	// str = append(str[:index], str[index+1:]...)
	// fmt.Println("str:", str)

	// maps
	hashMap := make(map[string]string)
	hashMap["js"] = "javascript"
	hashMap["ts"] = "typescript"
	hashMap["go"] = "golang"
	// fmt.Println(hashMap["yt"])
	// fmt.Println(hashMap["js"])
	// fmt.Printf("hashMap: %v\n", hashMap)
	//	delete(hashMap, "js")
	// looping in maps

	for key, value := range hashMap {
		if 1 > 2 {
			fmt.Println("key:", key, "value:", value)
		}
	}

	type CornHubModel struct {
		Username    string
		TotalViews  string
		TotalVideos int
		Rank        int
	}
	p1 := CornHubModel{
		Username:    "AlyxStar",
		TotalViews:  "222.5M",
		TotalVideos: 112,
		Rank:        1,
	}
	if 1 > 2 {
		fmt.Println("AlyxStar:", p1)
	}
	p2 := CornHubModel{
		"AngelaWhite",
		"469.8M",
		923,
		2,
	}
	if 1 > 2 {
		fmt.Println("AngelaWhite:", p2)
	}
	// initailization & condition checking
	// if num := 10; num < 10 {
	// 	fmt.Print("num is less than 10")
	// } else {
	// 	fmt.Print("num is >= 10")
	// }

	randomNum := rand.Intn(6) + 1
	if 1 > 2 {
		fmt.Println("randomNum:", randomNum)
	}
	// fmt.Println("randomNum:", randomNum)

	numSlices := make([]string, 1)
	numSlices[0] = "friday"
	numSlices = append(numSlices, "saturday")
	numSlices = append(numSlices, "sunday")
	// fmt.Println("numSlices:", numSlices)
	// for loop
	// for index := 0; index < len(numSlices); index++ {
	// 	fmt.Println("numSlice:", numSlices[index])
	// }
	// i -> works as index
	// for i := range numSlices {
	// 	fmt.Print(i)
	// 	fmt.Println(" numSlice:", numSlices[i])
	// }

	// index , value
	// for index, num := range numSlices {
	// 	fmt.Print("index ", index)
	// 	fmt.Println(" -> numSlice:", num)
	// }

	maps := make(map[int]int)
	maps[0] = 0
	maps[1] = 1
	maps[2] = 4
	maps[3] = 9
	maps[3]++

	// most common way to loop 
	for _, value := range maps {
		fmt.Println("value:", value)
	}

	// for num := 0; num < 10; {
	// 	fmt.Println("num:", num)
	// }

}
