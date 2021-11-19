package main

type LearnSettings struct {
	numOfHiddenLayers        int
	numOfNodesInHiddenLayers int
	datasets                 []DataSet
}

type DataSet struct {
	inputs     []float64
	name       string
	expectedId int
}

func NewLearnSettings(numOfDatasets int, numOfInputs int, numOfHiddenLayers int, numOfNodesInHiddenLayers int) *LearnSettings {
	l := &LearnSettings{
		numOfHiddenLayers:        numOfHiddenLayers,
		numOfNodesInHiddenLayers: numOfNodesInHiddenLayers,
		datasets:                 make([]DataSet, numOfDatasets),
	}
	for i := 0; i < numOfDatasets; i++ {
		l.datasets[i].inputs = make([]float64, numOfInputs)
	}
	return l
}
