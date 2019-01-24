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
	fmt.Println("")
	B := Matrix{
		[]float64{2.0, 3.0},
		[]float64{1.0, 9.0},
		[]float64{3.0, 1.0},
	}
	B.Print()
	fmt.Println("")

	C := mult(A, B)
	C.Print()
	fmt.Println("")

	fmt.Println("The determinant of C is", det(C))
}
