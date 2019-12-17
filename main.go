package main

import (
	"fmt"
	"bufio"
	"os"
	"math"
	"strconv"
)

func getFuelReq(mass float64) (fuelReq float64) {
	fuelReq = math.Floor(mass/3) - 2

	return
}

func main() {
	fh, err := os.Open("data.txt")
	defer fh.Close()
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(fh)

	var totalFuelReq int

	for scanner.Scan() {
		mass, err := strconv.ParseFloat( scanner.Text(), 64 )
		if err != nil {
			panic(err)
		}

		moduleFuelReq := getFuelReq(mass)
		fuelReq := moduleFuelReq

		for {
			fuelReq = getFuelReq(fuelReq)
			if fuelReq < 1 { 
				break
			}

			moduleFuelReq += fuelReq
		}

		totalFuelReq += int(moduleFuelReq)
	}

	fmt.Println(totalFuelReq)
}