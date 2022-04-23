package main

import (
	"errors"
	"fmt"
)

func main() {
	scores := []int{10, 5, 8, 9, 7}
	total := sum(scores)
	fmt.Println("total dari sum:", total)

	result, err := calculate(10, 2, "a")
	if err != nil {
		fmt.Println("terjadi kesalahan")
		fmt.Println((err.Error()))
	} else {
		fmt.Println("Operation:", result)
	}

}

func sum(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

func calculate(number, numberTwo int, operation string) (int, error) {
	var result int
	var errorResult error
	switch operation {
	case "+":
		result = number + numberTwo
	case "-":
		result = number - numberTwo
	case "*":
		result = number * numberTwo
	case "/":
		result = number / numberTwo
	default:
		errorResult = errors.New("unknown operation")
	}

	return result, errorResult
}
