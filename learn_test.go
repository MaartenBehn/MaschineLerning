package main

import (
	"fmt"
	"testing"
)

func TestMultibleTestSets(t *testing.T) {

	learnSettings := NewLearnSettings(4, 4, 2, 8)
	learnSettings.datasets[0].inputs[0] = 1
	learnSettings.datasets[0].inputs[1] = 0
	learnSettings.datasets[0].inputs[2] = 0
	learnSettings.datasets[0].inputs[3] = 0
	learnSettings.datasets[0].name = "Zero"

	learnSettings.datasets[1].inputs[0] = 1
	learnSettings.datasets[1].inputs[1] = 1
	learnSettings.datasets[1].inputs[2] = 0
	learnSettings.datasets[1].inputs[3] = 1
	learnSettings.datasets[1].name = "One"

	learnSettings.datasets[2].inputs[0] = 1
	learnSettings.datasets[2].inputs[1] = 1
	learnSettings.datasets[2].inputs[2] = 1
	learnSettings.datasets[2].inputs[3] = 0
	learnSettings.datasets[2].name = "Two"

	learnSettings.datasets[3].inputs[0] = 1
	learnSettings.datasets[3].inputs[1] = 1
	learnSettings.datasets[3].inputs[2] = 1
	learnSettings.datasets[3].inputs[3] = 1
	learnSettings.datasets[3].name = "Zero"

	net := CreateNetFromLearnSettings(learnSettings, 1)
	net.setRandomWeigts()

	for i := 0; i < 10000; i++ {
		for _, dataset := range learnSettings.datasets {
			net.loadDataSet(dataset)
			net.forwardPath()
			net.backwardPath()
		}
	}

	for _, dataset := range learnSettings.datasets {
		net.loadDataSet(dataset)
		net.forwardPath()
		fmt.Println("Dataset: " + dataset.name)
		net.printOutput()
	}
}
