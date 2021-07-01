package main

import "fmt"

func dailyTemperatures(T []int) []int {

	stack := []int{}
	res := make([]int, len(T))

	for i, v := range T {

		for len(stack) != 0 && T[stack[len(stack)-1]] < v {
			index := stack[len(stack)-1]

			//attribute result
			res[index] = i - index

			//pop the last element until the incoming element is <= stack[-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	return res
}

func main() {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(temperatures))
	// fmt.Println(temperatures[:len(temperatures)-1])

}
