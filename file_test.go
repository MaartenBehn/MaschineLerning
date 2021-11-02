package main

import (
	"testing"
)

func TestLoadFile(t *testing.T) {
	loadFile("file.go")
	assertPanic(t, func() { loadFile("") })
}

func TestLoadImage(t *testing.T) {
	loadBmp("images/Dreieck000.bmp")
	assertPanic(t, func() { loadBmp("file.go") })
	assertPanic(t, func() { loadBmp("") })
}

func assertPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	f()
}
