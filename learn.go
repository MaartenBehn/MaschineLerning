package main

type LearnSettings struct {
	numOfHiddenLayers        int
	numOfNodesInHiddenLayers int
	numOfInputs              int
	datasets                 []*DataSet
	names                    []string
}

type DataSet struct {
	inputs []float64
	name   string
	id     int
}

func NewLearnSettings(numOfDatasets int, numOfInputs int, numOfHiddenLayers int, numOfNodesInHiddenLayers int) *LearnSettings {
	l := &LearnSettings{
		numOfHiddenLayers:        numOfHiddenLayers,
		numOfNodesInHiddenLayers: numOfNodesInHiddenLayers,
		numOfInputs:              numOfInputs,
		datasets:                 make([]*DataSet, numOfDatasets),
	}
	for i := 0; i < numOfDatasets; i++ {
		l.datasets[i] = &DataSet{
			inputs: make([]float64, numOfInputs),
			name:   "undefined",
		}
	}
	return l
}

func CreateNetFromLearnSettings(settings *LearnSettings, learparam float64) *NeuralNet {

	for _, dataset := range settings.datasets {

		isContained := false
		for i, name := range settings.names {
			if dataset.name == name {
				isContained = true
				dataset.id = i
				break
			}
		}

		if !isContained {
			dataset.id = len(settings.names)
			settings.names = append(settings.names, dataset.name)
		}
	}

	return NewNet(settings.numOfNodesInHiddenLayers, settings.numOfHiddenLayers, settings.numOfInputs, len(settings.names), learparam)
}

func (net *NeuralNet) loadDataSet(dataset *DataSet) {
	for i, input := range dataset.inputs {
		net.input.Set(i+1, 0, input)
	}

	for i := 0; i < net.outputLayer.expected.row; i++ {
		net.outputLayer.expected.Set(i, 0, 0)
	}
	net.outputLayer.expected.Set(dataset.id, 0, 1)
}
