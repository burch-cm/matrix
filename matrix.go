package main

import (
	"fmt"
	"math/rand"
)

type Matrix [][]float64

func (m Matrix) T() Matrix {
	mdim := MatrixDim{
		nrow: len(m),
		ncol: len(m[0]),
	}
	out := fMatrix(mdim.ncol, mdim.nrow)
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

func fMatrix(nrow, ncol int) Matrix {
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

func (m Matrix) mult(n Matrix) Matrix {
	mdim := MatrixDim{
		nrow: len(m),
		ncol: len(m[0]),
	}

	var row []float64
	var out Matrix

	for i := 0; i < mdim.nrow; i++ {
		var s []float64
		for j := 0; j < mdim.ncol; j++ {
			s = append(s, m[i][j]*n[j][i])
		}
		row = append(row, vsum(s))
		out = append(out, row)
	}
	return out
}

func (m Matrix) Print() {
	for _, v := range m {
		fmt.Println(v)
	}
}
