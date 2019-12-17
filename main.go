package main

import (
	"fmt"
)

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

func process(start, end, n, v int ) int {
	var result int

	code := []int{1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,1,6,19,1,9,19,23,1,6,23,27,1,10,27,31,1,5,31,35,2,6,35,39,1,5,39,43,1,5,43,47,2,47,6,51,1,51,5,55,1,13,55,59,2,9,59,63,1,5,63,67,2,67,9,71,1,5,71,75,2,10,75,79,1,6,79,83,1,13,83,87,1,10,87,91,1,91,5,95,2,95,10,99,2,9,99,103,1,103,6,107,1,107,10,111,2,111,10,115,1,115,6,119,2,119,9,123,1,123,6,127,2,127,10,131,1,131,6,135,2,6,135,139,1,139,5,143,1,9,143,147,1,13,147,151,1,2,151,155,1,10,155,0,99,2,14,0,0}
	code[1] = n
	code[2] = v

	for {
		if code[ start ] == 99 {
			result = code[0]
			break
		}

		if code[ start ] != 1 && code[ start ] != 2 {
			result = code[0]
			break
		}

		curCode := code[ start : end +1 ]
		var num int

		switch curCode[0] {
		case 1:
			num = add( code[ curCode[1] ], code[ curCode[2] ])
		case 2:
			num = multiply( code[ curCode[1] ], code[ curCode[2] ])
		}
		
		code[ curCode[3] ] = num

		start += 4
		end   += 4
	}

	if result == 19690720 {
		return 100 * n + v
	}

	return 0
}

func main() {
	var out  int
	startIdx := 0
	endIdx 	 := 3

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100 ; verb++ {
			out = process(startIdx, endIdx, noun, verb)
			if out != 0 {
				break
			}
		}

		if out != 0  {
			break
		}
	}

	fmt.Printf("Result=%d\n", out)
}