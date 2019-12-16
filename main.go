package main

import (
	"fmt"
	"bufio"
	"os"
	"math"
	"strconv"
)

func getFuelReq(mod float64) (fuelReq int) {
	fuelReq = int(math.Floor(mod/3) - 2)

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
		input, err := strconv.ParseFloat( scanner.Text(), 64 )
		if err != nil {
			panic(err)
		}

		totalFuelReq += getFuelReq(input)
	}

	fmt.Println(totalFuelReq)
}