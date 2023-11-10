package main

import "fmt"

func factorialLoop(value int) (int, int) {
	result := 1

	for i := value; i > 0; i-- {
		result *= i
	}

	return result, 10
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	}
	return value * factorialRecursive(value-1)
}

func main() {
	loop, i := factorialLoop(10)
	fmt.Println(loop, i)
}
