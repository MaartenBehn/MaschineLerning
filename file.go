package main

import (
	"golang.org/x/image/bmp"
	"image"
	"os"
)

// Func
func loadFile(path string) (file *os.File) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return file
}

func loadBmp(path string) (img image.Image) {
	file := loadFile(path)
	img, err := bmp.Decode(file)
	if err != nil {
		panic(err)
	}
	return img
}
