package main

import "strconv"

func getNumericValue(n string) (int, error) {
	value, err := strconv.Atoi(n)
	if err != nil {
		return 0, err
	}
	return value, nil
}

func isLeastThanOne(n int) bool {
	return n < 1
}
