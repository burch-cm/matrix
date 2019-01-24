package main

import (
	"fmt"
)

func main() {
	A := Matrix{
		[]float64{1.0, 2.0, 3.0},
		[]float64{4.0, 5.0, 6.0},
	}
	A.Print()
	B := A.T()
	fmt.Println("")
	B.Print()
}
