package main

import (
	"testing"
)

func TestMath(t *testing.T) {
	t.Log(sigmoidDerivationFunc(1))
	t.Log(sigmoidDerivationFunc(0.6394))
}
