package main

import "fmt"

func exchange(nums []int) []int {
	i, j := 0, 0
	for j = 0; j < len(nums); j++ {
		if nums[j]%2 != 0 {

			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	return nums
}

func exchange_gopher(nums []int) []int {
	i := 0
	for j, v := range nums {
		if v%2 != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	return nums

}

func main() {
	nums := []int{1, 2, 3, 4}
	// fmt.Println(exchange(nums))
	fmt.Println(exchange_gopher(nums))
}
