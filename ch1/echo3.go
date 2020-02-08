package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args[1:])
	fmt.Println(os.Args[0])
	for i, arg := range os.Args {
		fmt.Printf("index %d， value %s\n", i, arg)
	}
}
