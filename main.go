package main

import (
	"fmt"
	"math"
)

const (
	lowerBound float64 = 136818
	upperBound float64 = 685979
	pwLength   int     = 6
)

func getIntUnits(whole float64) (units []float64) {
	curVal := whole

	for i := pwLength - 1; i >= 0; i-- {
		pow10 := math.Pow10(i)

		unit := curVal / pow10
		unit = math.Floor(unit)

		curVal -= (pow10 * unit)

		units = append(units, unit)
	}

	return
}

func checkForDoubles(digits []float64, passed *int) {
	for i := range digits {
		if i == len(digits)-1 {
			break
		}

		if digits[i] == digits[i+1] {
			(*passed)++
			break
		}
	}
}

func checkIncreasing(digits []float64, passed *int) {
	var allGood bool

	for i := range digits {
		if i == len(digits)-1 {
			break
		}

		if digits[i] <= digits[i+1] {
			allGood = true
		} else {
			allGood = false
			break
		}
	}

	if allGood {
		(*passed)++
	}
}

func isPossible(digits []float64) bool {
	var checksPassed int

	checkForDoubles(digits, &checksPassed)

	checkIncreasing(digits, &checksPassed)

	if checksPassed == 2 {
		return true
	}

	return false
}

func main() {
	start := lowerBound
	end := upperBound

	var numOptions int

	for passwdTrial := start; passwdTrial < end; passwdTrial++ {
		units := getIntUnits(passwdTrial)

		if isPossible(units) {
			numOptions++
		}
	}

	fmt.Println(numOptions)
}
