package main

import (
	"testing"
)

// INFO: I know, numbers are global, it is not the best idea, in production it will not have any place! ;]

func TestGetNumbersIndexReturnVaidIndex(t *testing.T) {
	numbers = []int{10, 20, 30, 100, 110, 120, 140}

	var index = getNumberIndex(100, 50)

	if index != 3 {
		t.Errorf("getNumberIndex(100), 50 = %d; want 3", index)
	}
}

func TestGetNumbersIndexReturnClosestIndex(t *testing.T) {
	numbers = []int{10, 20, 30, 100, 110, 120, 140}

	var index = getNumberIndex(102, 50)

	if index != 3 {
		t.Errorf("getNumberIndex(102, 50) = %d; want 3", index)
	}
}


func TestGetNumbersIndexReturnMinusOneIfNoTolerate(t *testing.T) {
	numbers = []int{10, 20, 30, 100, 110, 120, 140}

	var index = getNumberIndex(105, 1)

	if index != -1 {
		t.Errorf("getNumberIndex(105, 1) = %d; want -1", index)
	}
}


func TestGetNumbersIndexReturnMinusOneIfNoNumberExists(t *testing.T) {
	numbers = []int{10, 20, 30, 100, 110, 120, 140}

	var index = getNumberIndex(400, 30)

	if index != -1 {
		t.Errorf("getNumberIndex(400, 30) = %d; want -1", index)
	}
}

