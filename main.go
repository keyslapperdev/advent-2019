package main

import (
	"fmt"
	"math"
	"os"
)

type opFunc func([]int, *IntCode)

type IntCode struct{
	fullCode *[]int
	startIdx int
	endIdx   int
}

func (ic IntCode) getInstructions(start, end int) []int {
	return (*ic.fullCode)[start:end]	
}

func getOpFunc(opc int) (opFunc, int) {
	_, opcode := parseOpc(opc)

	switch opcode {
	case 1:
		return opc1, 4
	case 2:
		return opc2, 4
	case 3:
		return opc3, 2
	case 4:
		return opc4, 2
	case 5:
		return opc5, 3
	case 6:
		return opc6, 3
	case 7:
		return opc7, 4
	case 8:
		return opc8, 4
	case 99:
		return opc99, 1
	default:
		return opcERR, 1
	}
}

func opcGeneric(instructions []int, fullCode *[]int) (int, int, int) {
	pModes, _ := parseOpc(instructions[0])

	var (
		operandA    int
		operandB    int
		operandC    int
	)

	if pModes[2] == 1 {
		operandA = instructions[1] //Immediate Mode
	} else {
		operandA = (*fullCode)[instructions[1]] //Position Mode
	}

	if len(instructions) > 2 {
		if pModes[1] == 1 {
			operandB = instructions[2]
		} else {
			operandB = (*fullCode)[instructions[2]]
		}

		if len(instructions) > 3 {
			operandC = instructions[3]
		}
	}

	return operandA, operandB, operandC
}

func opc1(instructions []int, ic *IntCode) {
	p1, p2, dest := opcGeneric(instructions, ic.fullCode)

	result := p1 + p2
	(*ic.fullCode)[dest] = result
}

func opc2(instructions []int, ic *IntCode) {
	p1, p2, dest := opcGeneric(instructions, ic.fullCode)

	result := p1 * p2
	(*ic.fullCode)[dest] = result
}

func opc3(instructions []int, ic *IntCode) {
	var input int
	fmt.Print("Please enter [Device ID]: ")
	fmt.Scan(&input)

	(*ic.fullCode)[instructions[1]] = input
}

func opc4(instructions []int, ic *IntCode) {
	p1, _, _ := opcGeneric(instructions, ic.fullCode)

	fmt.Printf("Output: %d\n", p1 )
}

func opc5(instructions []int, ic *IntCode) {
	p1, newIdx, _ := opcGeneric(instructions, ic.fullCode)

	if p1 != 0 {
		ic.startIdx = newIdx
	}
}

func opc6(instructions []int, ic *IntCode) {
	p1, newIdx, _ := opcGeneric(instructions, ic.fullCode)

	if p1 == 0 {
		ic.startIdx = newIdx
	}
}

func opc7(instructions []int, ic *IntCode) {
	p1, p2, dest := opcGeneric(instructions, ic.fullCode)

	var result int
	if p1 < p2 {
		result = 1
	} else {
		result = 0
	}

	(*ic.fullCode)[dest] = result
}

func opc8(instructions []int, ic *IntCode) {
	p1, p2, dest := opcGeneric(instructions, ic.fullCode)

	var result int
	if p1 == p2 {
		result = 1
	} else {
		result = 0
	}

	(*ic.fullCode)[dest] = result
}

func opc99(instructions []int, ic *IntCode) {
	os.Exit(0)
}

func opcERR(instructions []int, ic *IntCode) {
	fmt.Printf("Function Errored: %v\n", instructions )
	os.Exit(1)
}

func parseOpc(op int) ([]int, int) {
	var units []int
	opc := float64(op)

	for i := 4; i >= 0; i-- {
		pow10 := math.Pow10(i)

		unit := opc / pow10
		unit = math.Floor(unit)

		opc -= (pow10 * unit)

		units = append(units, int(unit))
	}

	opcode := (units[3] * 10) + units[4]
	params := units[:3]

	return params, opcode
}

func main() {
	data := []int{3, 225, 1, 225, 6, 6, 1100, 1, 238, 225, 104, 0, 1002, 36, 25, 224, 1001, 224, -2100, 224, 4, 224, 1002, 223, 8, 223, 101, 1, 224, 224, 1, 223, 224, 223, 1102, 31, 84, 225, 1102, 29, 77, 225, 1, 176, 188, 224, 101, -42, 224, 224, 4, 224, 102, 8, 223, 223, 101, 3, 224, 224, 1, 223, 224, 223, 2, 196, 183, 224, 1001, 224, -990, 224, 4, 224, 1002, 223, 8, 223, 101, 7, 224, 224, 1, 224, 223, 223, 102, 14, 40, 224, 101, -1078, 224, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 2, 224, 1, 224, 223, 223, 1001, 180, 64, 224, 101, -128, 224, 224, 4, 224, 102, 8, 223, 223, 101, 3, 224, 224, 1, 223, 224, 223, 1102, 24, 17, 224, 1001, 224, -408, 224, 4, 224, 1002, 223, 8, 223, 101, 2, 224, 224, 1, 223, 224, 223, 1101, 9, 66, 224, 1001, 224, -75, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 6, 224, 1, 223, 224, 223, 1102, 18, 33, 225, 1101, 57, 64, 225, 1102, 45, 11, 225, 1101, 45, 9, 225, 1101, 11, 34, 225, 1102, 59, 22, 225, 101, 89, 191, 224, 1001, 224, -100, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 1, 224, 1, 223, 224, 223, 4, 223, 99, 0, 0, 0, 677, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1105, 0, 99999, 1105, 227, 247, 1105, 1, 99999, 1005, 227, 99999, 1005, 0, 256, 1105, 1, 99999, 1106, 227, 99999, 1106, 0, 265, 1105, 1, 99999, 1006, 0, 99999, 1006, 227, 274, 1105, 1, 99999, 1105, 1, 280, 1105, 1, 99999, 1, 225, 225, 225, 1101, 294, 0, 0, 105, 1, 0, 1105, 1, 99999, 1106, 0, 300, 1105, 1, 99999, 1, 225, 225, 225, 1101, 314, 0, 0, 106, 0, 0, 1105, 1, 99999, 8, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 329, 1001, 223, 1, 223, 108, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 344, 1001, 223, 1, 223, 7, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 359, 101, 1, 223, 223, 7, 226, 677, 224, 102, 2, 223, 223, 1006, 224, 374, 101, 1, 223, 223, 1008, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 389, 101, 1, 223, 223, 8, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 404, 101, 1, 223, 223, 8, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 419, 1001, 223, 1, 223, 1107, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 434, 1001, 223, 1, 223, 1107, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 449, 1001, 223, 1, 223, 107, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 464, 1001, 223, 1, 223, 1008, 677, 677, 224, 1002, 223, 2, 223, 1006, 224, 479, 1001, 223, 1, 223, 1108, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 494, 1001, 223, 1, 223, 1108, 677, 677, 224, 1002, 223, 2, 223, 1006, 224, 509, 1001, 223, 1, 223, 107, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 524, 101, 1, 223, 223, 1007, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 539, 1001, 223, 1, 223, 1107, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 554, 1001, 223, 1, 223, 1008, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 569, 101, 1, 223, 223, 1108, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 584, 101, 1, 223, 223, 108, 677, 677, 224, 1002, 223, 2, 223, 1006, 224, 599, 1001, 223, 1, 223, 1007, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 614, 101, 1, 223, 223, 107, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 629, 101, 1, 223, 223, 1007, 226, 226, 224, 102, 2, 223, 223, 1005, 224, 644, 1001, 223, 1, 223, 108, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 659, 1001, 223, 1, 223, 7, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 674, 1001, 223, 1, 223, 4, 223, 99, 226}

	intCode := IntCode{
		fullCode: &data,
	}

	var curCode []int

	for {
		//Get process for handling opcode and amount of required
		//parameters
		process, paramCount := getOpFunc((*intCode.fullCode)[intCode.startIdx])
		intCode.endIdx = intCode.startIdx + paramCount

		//Load variable with current instruction snippit
		curCode = intCode.getInstructions(intCode.startIdx, intCode.endIdx)

		//Move forward the index by the number of consumed
		//parameters
		intCode.startIdx += paramCount

		//Use returned function to process code
		process(curCode, &intCode)
	}
}
