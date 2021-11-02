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

	net.forwardPath()

	fmt.Println(net.outputLayer.output)

	net.outputLayer.expected.Set(0, 0, 0.9)
	net.outputLayer.expected.Set(1, 0, 0.2)

	net.backwardPath()

	fmt.Println(net.layers[0].errSig)
	fmt.Println(net.layers[1].errSig)

	fmt.Println(net.layers[0].weights)
	fmt.Println(net.layers[1].weights)
}
