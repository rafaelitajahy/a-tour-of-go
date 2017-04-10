package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.pi) // error - correct is math.Pi
}

/*
		RunFormatReset
		# command-line-arguments
		.\compile6.go:9: cannot refer to unexported name math.pi
		.\compile6.go:9: undefined: math.pi
*/