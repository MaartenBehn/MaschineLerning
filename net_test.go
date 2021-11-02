package main

import (
	"fmt"
	"testing"
)

func TestNet(t *testing.T) {
	net := NewNet()
	t.Log(net)
}

func TestBackward(t *testing.T) {
	net := NewNet()
	net.inputLayer[0] = 1
	net.inputLayer[1] = 0.7
	net.inputLayer[2] = 0.6

	net.hiddenLayers[0].weights = NewMatrix(nodesPerLayer, len(net.inputLayer))
	net.hiddenLayers[0].weights.Set(0, 0, 0.3)
	net.hiddenLayers[0].weights.Set(0, 1, 0.8)
	net.hiddenLayers[0].weights.Set(0, 2, 0.5)
	net.hiddenLayers[0].weights.Set(1, 0, -0.2)
	net.hiddenLayers[0].weights.Set(1, 1, -0.6)
	net.hiddenLayers[0].weights.Set(1, 2, 0.7)

	net.hiddenLayers[0].input = NewMatrix(len(net.inputLayer), 1)
	net.hiddenLayers[0].input.Set(0, 0, 1)
	net.hiddenLayers[0].input.Set(1, 0, 0.7)
	net.hiddenLayers[0].input.Set(2, 0, 0.6)

	net.hiddenLayers[0].result = net.hiddenLayers[0].weights.Mul(net.hiddenLayers[0].input)

	net.hiddenLayers[0].output = NewMatrix(net.hiddenLayers[0].result.row, net.hiddenLayers[0].result.collum)
	for i := 0; i < net.hiddenLayers[0].result.row; i++ {
		net.hiddenLayers[0].output.Set(i, 0, sigmoidFunc(net.hiddenLayers[0].result.Get(i, 0)))
	}

	fmt.Print(net.hiddenLayers[0].output)
}
