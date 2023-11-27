# GO Lang
- Go is a statically typed, compiled high-level programming language
- Almost everything in Go is funtions and channels
- Object Oriented Programming - YES and NO [ it has - struct and interface  but no operator overloading ]
-No Try Catch
- lexer does a lot of work(putting semiclon etc)
----------------------------------------------------------
### go help 

- go mod init golang
- go run .
- go env GOPATH : to see the path of the go workspace in my machine
- lexer checks code following the rules of the language
- go case sensitive 
- everything in go is a type [ String , Bool , Integer , Floot , Complex]
- variable declaration
  - var var_name type = value
  - example var name string = "Golang"
  - default value of string is empty string , default value of int is 0 , default value of bool is false etc
- no var style [ most common style]
  - var_name := value
  - example name := "Golang"
  - Its only possible inside function , we cant use it outside function(globally)

- no new variables on left side of `:=`
- keyword with small first letter is private , keyword with capital first letter is public
- memory allocation , deallocation is done by go automatically [ garbage collection]
- nil = NULL
