package main

import (
	"fmt"
	"math/rand"
	"time"
)

const bias = 1

type NeuralNet struct {
	hiddenLayersAmmount int
	inputNodes          int
	outputNodes         int
	learnparam          float64

	input       *matrix
	layers      []Layer
	outputLayer *Layer
}

type Layer struct {
	nodesAmount int
	input       *matrix
	weights     *matrix
	netInput    *matrix
	output      *matrix
	errSig      *matrix
	expected    *matrix
}

func NewNet(nodesPerHiddenLayer int, hiddenLayersPerNet int, inputNodes int, outputNodes int, lernparam float64) *NeuralNet {
	net := &NeuralNet{
		hiddenLayersAmmount: hiddenLayersPerNet,
		inputNodes:          inputNodes,
		outputNodes:         outputNodes,
		learnparam:          lernparam,
	}

	net.input = NewMatrix(inputNodes+1, 1)
	net.input.Set(0, 0, bias)

	net.layers = make([]Layer, hiddenLayersPerNet+1)
	for i := 0; i < hiddenLayersPerNet+1; i++ {

		if i == 0 {
			net.layers[i].input = net.input
		} else {
			net.layers[i].input = net.layers[i-1].output
		}

		if i < hiddenLayersPerNet {
			net.layers[i].nodesAmount = nodesPerHiddenLayer
		} else {
			net.layers[i].nodesAmount = outputNodes
		}

		net.layers[i].weights = NewMatrix(net.layers[i].nodesAmount, net.layers[i].input.row)
		net.layers[i].output = NewMatrix(net.layers[i].nodesAmount+1, 1)
		net.layers[i].output.Set(0, 0, bias)
		net.layers[i].netInput = NewMatrix(net.layers[i].nodesAmount, 1)
		net.layers[i].errSig = NewMatrix(net.layers[i].nodesAmount, 1)
	}

	net.outputLayer = &net.layers[hiddenLayersPerNet]
	net.outputLayer.expected = NewMatrix(outputNodes, 1)

	return net
}

func (net *NeuralNet) setRandomWeigts() {
	for i := 0; i < len(net.layers); i++ {
		net.layers[i].setRandomWeigts()
	}
}

func (layer Layer) setRandomWeigts() {
	rand.Seed(time.Now().Unix())
	for i := 0; i < layer.weights.row; i++ {
		for j := 0; j < layer.weights.collum; j++ {
			layer.weights.Set(i, j, rand.Float64())
		}
	}
}

func (net *NeuralNet) forwardPath() {
	for i := 0; i < net.hiddenLayersAmmount+1; i++ {
		layer := net.layers[i]
		result := layer.weights.Mul(layer.input)

		for i := 0; i < result.row; i++ {
			val := result.Get(i, 0)
			layer.netInput.Set(i, 0, val)
			layer.output.Set(i+1, 0, sigmoidFunc(val))
		}
	}
}

func (net *NeuralNet) backwardPath() {

	layer := *net.outputLayer
	for i := 0; i < layer.errSig.row; i++ {
		sig := sigmoidDerivationFunc(layer.netInput.Get(i, 0)) *
			(layer.expected.Get(i, 0) - layer.output.Get(i+1, 0))

		layer.errSig.Set(i, 0, sig)
	}

	for k := net.hiddenLayersAmmount - 1; k >= 0; k-- {
		layer := net.layers[k]
		for i := 0; i < layer.errSig.row; i++ {
			sig := sigmoidDerivationFunc(layer.netInput.Get(i, 0))

			a := 0.0
			backLayer := net.layers[k+1]
			for l := 0; l < backLayer.errSig.row; l++ {
				a += backLayer.errSig.Get(l, 0) * backLayer.weights.Get(l, i+1)
			}
			sig *= a

			layer.errSig.Set(i, 0, sig)
		}
	}

	for _, layer := range net.layers {
		for i := 0; i < layer.errSig.row; i++ {
			for j := 0; j < layer.input.row; j++ {
				weigthDelta := net.learnparam * layer.errSig.Get(i, 0) * layer.input.Get(j, 0)
				weigth := layer.weights.Get(i, j) + weigthDelta
				layer.weights.Set(i, j, weigth)
			}
		}
	}
}

func (net *NeuralNet) print() {
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

func (net *NeuralNet) printOutput() {
	net.outputLayer.output.Print()
}
