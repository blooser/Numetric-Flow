package main

import (
	"bufio"
	"os"
	"strconv"
	logger "github.com/sirupsen/logrus"
)

var numbers[] int

func loadNumbers(){
	file, err := os.Open("numbers.txt")	

	if err != nil {
		logger.Error(err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
	
		number, err := strconv.Atoi(line)

		if err != nil {
			continue
		}
		
		numbers = append(numbers, number)
	}

	if err := scanner.Err(); err != nil {
		logger.Error(err)
	}

	logger.WithFields(logger.Fields{
		"size": len(numbers),
	}).Info("Numbers loaded")
}


func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}

func getNumberIndex(number int, tolerance float64) int {
	//NOTE: This is binary search with tracking the closest number with tolerance (%) parameter

	low := 0
	high := len(numbers) - 1

	closest := -1
	closestPercentage := 100.0

	for low <= high {
		mid := low + (high-low)/2
		midval := numbers[mid]

		diff := abs(midval - number)
		percentageDiff := (float64(diff) / float64(number)) * 100.0

		// Let's calculate percentage diff!
		if percentageDiff < closestPercentage {
			closestPercentage = percentageDiff	
			closest = mid
		}


		// Basic binary search
		if midval == number {
			return mid
		} else if midval < number {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	// Do we fit in tolerance?
	if closestPercentage <= tolerance {
		return closest
	}

	return -1
}
