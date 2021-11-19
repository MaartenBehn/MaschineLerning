package main

import (
	"math/rand"
	"time"
)

type NeuralNet struct {
	nodesPerLayer      int
	hiddenLayersPerNet int
	inputNodes         int
	outputNodes        int
	bias               float64
	lernparam          float64

	input       *matrix
	layers      []Layer
	outputLayer *Layer
}

type Layer struct {
	input    *matrix
	weights  *matrix
	netInput *matrix
	output   *matrix
	errSig   *matrix
	expected *matrix
}

func NewNet(nodesPerLayer int, hiddenLayersPerNet int, inputNodes int, outputNodes int, bias float64, lernparam float64) *NeuralNet {
	net := &NeuralNet{
		nodesPerLayer:      nodesPerLayer,
		hiddenLayersPerNet: hiddenLayersPerNet,
		inputNodes:         inputNodes,
		outputNodes:        outputNodes,
		bias:               bias,
		lernparam:          lernparam,
	}

	net.input = NewMatrix(inputNodes+1, 1)
	net.input.Set(0, 0, bias)

	for i := 0; i < len(net.layers); i++ {

		if i == 0 {
			net.layers[i].input = net.input
		} else {
			net.layers[i].input = net.layers[i-1].output
		}
		net.layers[i].weights = NewMatrix(nodesPerLayer, net.layers[i].input.row)
		net.layers[i].output = NewMatrix(nodesPerLayer+1, 1)
		net.layers[i].output.Set(0, 0, bias)
		net.layers[i].netInput = NewMatrix(nodesPerLayer, 1)
		net.layers[i].errSig = NewMatrix(nodesPerLayer, 1)
	}

	net.outputLayer = &net.layers[hiddenLayersPerNet]
	net.outputLayer.expected = NewMatrix(nodesPerLayer, 1)

	return net
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
	for i := 0; i < net.hiddenLayersPerNet+1; i++ {
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

	for k := net.hiddenLayersPerNet - 1; k >= 0; k-- {
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
			for j := 0; j < net.nodesPerLayer+1; j++ {
				weigthDelta := net.lernparam * layer.errSig.Get(i, 0) * net.input.Get(j, 0)
				weigth := layer.weights.Get(i, j) + weigthDelta
				layer.weights.Set(i, j, weigth)
			}
		}
	}
}
