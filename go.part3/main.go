package main

import (
	"fmt"
	"io"
	"os"
)

func checkNilErr(err error) {
	if err != nil {
		fmt.Println("error:", err)
		panic(err)
	}
}

// func add(num1 int, num2 int) int {
// 	result := num1 + num2
// 	return result
// }

func advAdder(nums ...int) (string, int) {
	total := 0
	for _, num := range nums {
		total += num
	}
	return "sum of all nums is:", total
}

func main() { // entry point of the program
	msg := "Hey User!💩"
	fmt.Println(msg)
	// functions
	num1 := 10
	num2 := 20
	// result := add(num1, num2)
	// fmt.Println("result:", result)

	// // immediately invoked function
	// func(num1 int, num2 int) {
	// 	fmt.Println("result IIF:", num1+num2)
	// }(num1, num2)

	num3 := 30
	num4 := 40
	_, res := advAdder(num1, num2, num3, num4)
	if 1 > 3 {
		fmt.Println("res:", res)
	}

	// u1 := User{
	// 	"John",
	// 	"john@ok.com",
	// }
	// fmt.Println("user:", u1)
	// u1.GetNewEmail()
	// fmt.Println("user:", u1)

	fileMsg := "Hey User! 🫥"
	file, err := os.Create("../go.part3")
	checkNilErr(err)
	length, err := io.WriteString(file, fileMsg)
	checkNilErr(err)
	fmt.Println("length:", length)

	s := "http://lco.dev"
	fmt.Println(s)
	// url parsing
	//	u, err := s.Parse(s)

}

type User struct {
	Name  string
	Email string
}

// method
func (u *User) GetNewEmail() {
	u.Email = "john@go.com"
}
