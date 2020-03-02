package main

import (
	"fmt"

	"github.com/gopl.io/ch3/basename"
)

func main() {
	fmt.Println(basename.Basename1("~/go/src/testdata/abc.go"))
	fmt.Println(basename.Basename2("~/go/src/testdata/a.b.c"))
}
