package main

const nodesPerLayer = 2
const hiddenLayersPerNet = 1
const inputNodes = 2
const outputNodes = 2

type NeuralNet struct {
	input        *matrix
	hiddenLayers [hiddenLayersPerNet]Layer
	outputLayer  Layer
}

type Layer struct {
	input   *matrix
	weights *matrix
	result  *matrix
	output  *matrix
}

func NewNet() *NeuralNet {
	net := &NeuralNet{}

	net.input = NewMatrix(inputNodes, 1)

	net.hiddenLayers[0].input = net.input
	net.hiddenLayers[0].weights = NewMatrix(nodesPerLayer, net.input.row)

	for i := 1; i < len(net.hiddenLayers); i++ {
		net.hiddenLayers[i].weights = NewMatrix(nodesPerLayer, len(net.inputLayer))
	}

	return net
}

func doForwardPathForLayer(layer Layer) {
	layer.result = layer.weights.Mul(layer.input)

	layer.output = NewMatrix(layer.result.row, layer.result.collum)
	for i := 0; i < layer.result.row; i++ {
		layer.output.Set(i, 0, sigmoidFunc(layer.result.Get(i, 0)))
	}
}

func (net *NeuralNet) forwardPath() {
	for i := 0; i < hiddenLayersPerNet; i++ {
		doForwardPathForLayer(net.hiddenLayers[i])
	}
	doForwardPathForLayer(net.outputLayer)
}
