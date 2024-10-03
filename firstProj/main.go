package main

import (
	"errors"
	"fmt"
)

func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("factorial is not defined for negative numbers")
	}
	if n == 0 {
		return 1, nil
	}
	factorial := 1
	for i := 1; i <= n; i++ {
		factorial *= i
	}
	return factorial, nil
}

func IntToBinary(num int) (string, error) {
	if num < 0 {
		return nil, errors.New("negative numbers are not allowed")
	}
}

func main() {
	fmt.Println(GetCharacterAtPosition("Зенит чемпион!", 20))
}
