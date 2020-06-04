package main

import (
	"fmt"
)

func main() {
	s := createSpiral("20")
	for _, v := range s {
		fmt.Println(v)
	}

}

func createSpiral(n string) [][]int {
	value, err := getNumericValue(n)
	if err != nil {
		return make([][]int, 0)
	}
	if isLeastThanOne(value) {
		return make([][]int, 0)
	}
	spiral := make([][]int, value)
	for i := 0; i < value; i++ {
		for j := 0; j < value; j++ {
			spiral[i] = append(spiral[i], 0)
		}
	}
	counter := 1
	row := 0
	col := 0
	up := movementUp{}
	left := movementLeft{}
	left.setNext(&up)
	down := movementDown{}
	down.setNext(&left)
	right := movementRight{}
	right.setNext(&down)
	up.setNext(&right)
	counter, col, row = right.execute(spiral, counter, row, col)
	return spiral
}
