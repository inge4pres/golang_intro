package main

import "fmt"

var first = 1
var second int

func init() {
	second = 2
}
func main() {
	third := 3
	numbers := make([]int, 0)
	numbers = []int{first, second, third}
	fmt.Printf("sum = %d", sum(numbers...))
}

func sum(numbers ...int) int {
	var result int
	for _, n := range numbers {
		result += n
	}
	return result
}
