package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse" // this external lib is responsible for reverse any string
)

func main() {
    fmt.Println(reverse.String("Hello world"))
	fmt.Println(reverse.String("Ao contrário igual o CD da Xuxa"))
}

// run on terminal: "go run ."