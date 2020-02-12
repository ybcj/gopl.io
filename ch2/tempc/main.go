package main

import (
	"fmt"

	"github.com/gopl.io/ch2/tempconv"
)

func main() {
	fmt.Println(tempconv.CToF(tempconv.BoilingC))
	fmt.Println(tempconv.FToC(tempconv.CToF(tempconv.AbsoluteZeroC)))
}
