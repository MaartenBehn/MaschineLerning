package main

import "fmt"

type matrix struct {
	row    int
	collum int
	data   []float64
}

func NewMatrix(row int, collum int) *matrix {
	return &matrix{
		row:    row,
		collum: collum,
		data:   make([]float64, row*collum),
	}
}

func (m *matrix) Get(row int, collum int) float64 {
	return m.data[row*m.collum+collum]
}

func (m *matrix) Set(row int, collum int, data float64) {
	m.data[row*m.collum+collum] = data
}

func (m *matrix) Mul(m2 *matrix) (result *matrix) {
	if m.row != m2.collum && m.collum != m2.row {
		panic("Diese Mat Mul gehen nicht.")
	}

	var collum int
	if m2.row == m.collum {
		result = NewMatrix(m.row, m2.collum)
		collum = m.collum

		for i := 0; i < result.row; i++ {
			for j := 0; j < result.collum; j++ {

				var val float64
				for k := 0; k < collum; k++ {
					in1 := m.Get(i, k)
					in2 := m2.Get(k, j)
					val += in1 * in2
				}
				result.Set(i, j, val)
			}
		}
	} else if m.row == m2.collum {
		result = NewMatrix(m2.row, m.collum)
		collum = m2.collum

		for i := 0; i < result.row; i++ {
			for j := 0; j < result.collum; j++ {

				var val float64
				for k := 0; k < collum; k++ {
					in1 := m2.Get(i, k)
					in2 := m.Get(k, j)
					val += in1 * in2
				}
				result.Set(i, j, val)
			}
		}
	}

	return result
}
func (m *matrix) Print() {
	for j := 0; j < m.row; j++ {
		fmt.Print("|")
		for k := 0; k < m.collum; k++ {
			fmt.Printf("% 01.4f  | ", m.Get(j, k))
		}
		fmt.Print("\n")
	}
}
