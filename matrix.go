package main

import (
	"errors"
	"fmt"
	"math/rand"
)

type Matrix [][]float64

func (m Matrix) T() Matrix {
	mdim := MatrixDim{
		nrow: len(m),
		ncol: len(m[0]),
	}
	out := makeMatrix(mdim.ncol, mdim.nrow)
	for i := 0; i < mdim.nrow; i++ {
		for j := 0; j < mdim.ncol; j++ {
			out[j][i] = m[i][j]
		}
	}
	return out
}

func vsum(s []float64) float64 {
	var out float64
	for _, v := range s {
		out = out + v
	}
	return out
}

func vmult(a, b []float64) []float64 {
	var out []float64
	for i := range a {
		out = append(out, a[i]*b[i])
	}
	return out
}

func makeMatrix(nrow, ncol int) Matrix {
	var out Matrix
	for i := 0; i < nrow; i++ {
		r := []float64{}
		for j := 0; j < ncol; j++ {
			r = append(r, rand.Float64())
		}
		out = append(out, r)
	}
	return out
}

type MatrixDim struct {
	nrow int
	ncol int
}

func (m Matrix) Print() {
	for _, v := range m {
		fmt.Println(v)
	}
}

func mult(m, n Matrix) Matrix {
	t := n.T()
	mdim := MatrixDim{
		nrow: len(m),
		ncol: len(m[0]),
	}
	var out Matrix
	for i := 0; i < mdim.nrow; i++ {
		outrow := []float64{}
		for j := 0; j < mdim.nrow; j++ {
			row := vmult(m[i], t[j])
			outrow = append(outrow, vsum(row))
		}
		out = append(out, outrow)
	}
	return out
}

func (m *Matrix) isSquare() bool {
	nrow := len(m)
	ncol := len(m[0])
	if nrow != ncol {
		return false
	} else {
		return true
	}
}

/*func det(m Matrix) (float64, error) {
	if isSquare != true {
		return 0, errors.New("Matrix is not square.")
	}
	var r float64
	r = m[0][0]*m[1][1] - m[1][0]*m[0][1]
	return r
}
*/
