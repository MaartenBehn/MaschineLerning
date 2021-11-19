package main

import (
	"fmt"
	"testing"
)

func TestNet(t *testing.T) {
	NewNet(2, 1, 2, 2, 1, 1)
}

const randomWeigths = true
const printDebug = false

func TestLearnNet(t *testing.T) {
	net := NewNet(2, 1, 2, 2, 1, 1)

	net.input.Set(1, 0, 0.7)
	net.input.Set(2, 0, 0.6)

	if randomWeigths {
		for i := 0; i < len(net.layers); i++ {
			net.layers[i].setRandomWeigts()
		}
	} else {
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
	}

	net.outputLayer.expected.Set(0, 0, 0.9)
	net.outputLayer.expected.Set(1, 0, 0.2)

	if printDebug {
		fmt.Println("--- Parameter ---")
		fmt.Println("Input Layer:")
		net.input.Print()
		fmt.Print("\n")
		printLayers(net)
	}

	for i := 0; i < 200; i++ {
		fmt.Printf("\r--- Round %d ---", i)

		net.forwardPath()

		if printDebug {
			fmt.Println("\n--- Forward Path ---")
			printLayers(net)
		}

		net.backwardPath()

		if printDebug {
			fmt.Println("--- Backard Path ---")
			printLayers(net)
		}
	}
	fmt.Println("")

	printLayers(net)
	if net.outputLayer.output.Get(0, 0)-
		net.outputLayer.expected.Get(0, 0) > 0.1 {
		t.Fail()
	}
}

func printLayers(net *NeuralNet) {
	fmt.Println("--- Layers ---")
	for i, layer := range net.layers {
		fmt.Printf("Layer %d:\n", i)

		fmt.Printf("Input:\n")
		layer.input.Print()

		fmt.Printf("Weigths:\n")
		layer.weights.Print()

		fmt.Printf("NetInput:\n")
		layer.netInput.Print()

		fmt.Printf("Output:\n")
		layer.output.Print()

		fmt.Printf("ErrSig:\n")
		layer.errSig.Print()

		fmt.Printf("Expected:\n")
		if layer.expected == nil {
			fmt.Printf("nill\n")
		} else {
			layer.expected.Print()
		}
	}
}
