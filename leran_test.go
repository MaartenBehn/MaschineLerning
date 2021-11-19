package main

import "testing"

func TestMultibleTestSets(t *testing.T) {

	learnSettings := NewLearnSettings(4, 4, 2, 8)
	learnSettings.datasets[0].inputs[0] = 1
	learnSettings.datasets[0].inputs[1] = 0
	learnSettings.datasets[0].inputs[2] = 0
	learnSettings.datasets[0].inputs[3] = 0
	learnSettings.datasets[0].expectedId = 0
	learnSettings.datasets[0].name = "0"

	learnSettings.datasets[1].inputs[0] = 0
	learnSettings.datasets[1].inputs[1] = 1
	learnSettings.datasets[1].inputs[2] = 0
	learnSettings.datasets[1].inputs[3] = 0
	learnSettings.datasets[1].expectedId = 1
	learnSettings.datasets[1].name = "1"

	learnSettings.datasets[2].inputs[0] = 0
	learnSettings.datasets[2].inputs[1] = 0
	learnSettings.datasets[2].inputs[2] = 1
	learnSettings.datasets[2].inputs[3] = 0
	learnSettings.datasets[2].expectedId = 2
	learnSettings.datasets[2].name = "2"

	learnSettings.datasets[3].inputs[0] = 0
	learnSettings.datasets[3].inputs[1] = 0
	learnSettings.datasets[3].inputs[2] = 0
	learnSettings.datasets[3].inputs[3] = 1
	learnSettings.datasets[3].expectedId = 3
	learnSettings.datasets[3].name = "3"

}
