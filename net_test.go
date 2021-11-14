package main

import (
	"fmt"
	"testing"
)

func TestNet(t *testing.T) {
	NewNet()
}

func TestForward(t *testing.T) {
	net := NewNet()

	net.input.Set(1, 0, 0.7)
	net.input.Set(2, 0, 0.6)

	/*
		net.layers[0].weights.Set(0, 0, 0.3)
		net.layers[0].weights.Set(0, 1, 0.8)
		net.layers[0].weights.Set(0, 2, 0.5)
		net.layers[0].weights.Set(1, 0, -0.2)
		net.layers[0].weights.Set(1, 1, -0.6)
		net.layers[0].weights.Set(1, 2, 0.7)

		net.outputLayer.weights.Set(0, 0, 0.2)
		net.outputLayer.weights.Set(0, 1, 0.4)
		net.outputLayer.weights.Set(0, 2, 0.3)
		net.outputLayer.weights.Set(1, 0, 0.1)
		net.outputLayer.weights.Set(1, 1, -0.4)
		net.outputLayer.weights.Set(1, 2, 0.9)
	*/
	for i := 0; i < len(net.layers); i++ {
		net.layers[i].setRandomWeigts()
	}

	net.outputLayer.expected.Set(0, 0, 0.9)
	net.outputLayer.expected.Set(1, 0, 0.2)

	net.forwardPath()

	for i := 0; i < 20; i++ {
		net.backwardPath()
		net.forwardPath()

		fmt.Println("Output:")
		fmt.Println(net.outputLayer.output)

		fmt.Println("Err:")
		fmt.Println(net.outputLayer.errSig)

		fmt.Println("----------------------------------")
	}

}
