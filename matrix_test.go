package main

import "testing"

func TestMatrix(t *testing.T) {
	mat := NewMatrix(2, 2)
	mat.Set(0, 0, 1)
	if mat.Get(0, 0) != 1 {
		t.Error("Get Set wrong")
	}
}
