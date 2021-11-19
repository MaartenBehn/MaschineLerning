package main

import (
	"testing"
)

func TestMatrix(t *testing.T) {
	mat := NewMatrix(2, 2)
	mat.Set(0, 0, 1)
	if mat.Get(0, 0) != 1 {
		t.Error("Get Set wrong")
	}
}

func TestMatrixMul(t *testing.T) {
	mat := NewMatrix(2, 3)
	mat.Set(0, 0, 3)
	mat.Set(0, 1, 2)
	mat.Set(0, 2, 1)
	mat.Set(1, 0, 1)
	mat.Set(1, 1, 0)
	mat.Set(1, 2, 2)

	mat1 := NewMatrix(3, 2)
	mat1.Set(0, 0, 1)
	mat1.Set(0, 1, 2)
	mat1.Set(1, 0, 0)
	mat1.Set(1, 1, 1)
	mat1.Set(2, 0, 4)
	mat1.Set(2, 1, 0)

	mat2 := mat.Mul(mat1)
	if mat2.row != 2 || mat2.collum != 2 ||
		mat2.Get(0, 0) != 7 ||
		mat2.Get(0, 1) != 8 ||
		mat2.Get(1, 0) != 9 ||
		mat2.Get(1, 1) != 2 {
		t.Fail()
	}

	mat3 := mat1.Mul(mat)
	if mat3.row != 3 || mat3.collum != 3 {
		t.Fail()
	}

	mat4 := NewMatrix(2, 1)
	mat4.Set(0, 0, 3)
	mat4.Set(1, 0, 1)

	mat5 := mat4.Mul(mat1)
	mat5.Print()

	assertPanic(t, func() { mat.Mul(mat4) })
}
