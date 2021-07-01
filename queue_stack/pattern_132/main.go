package main

import "fmt"
import "math"

func main(){
	arr := []int{3, 1, 4, 2}//true
	arr2 := []int{1,0,1,-4,-3}//false
	arr3 := []int{3,5,0,3,4}//true
	fmt.Println(find132pattern(arr))
	fmt.Println(find132pattern(arr2))
	fmt.Println(find132pattern(arr3))

}

func find132pattern(nums []int) bool {

	stack := []int{}
	second := math.MinInt64

	for i:= len(nums)-1;i>=0;i--{

		for len(stack) != 0 && nums[i]> stack[len(stack)-1]{
			second = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

		}
		stack = append(stack,nums[i])
		if stack[len(stack)-1]<second{
			return true
		}
	}

	return false
}

func min(a,b int)int{
	if a<b{
		return a 
	}
	return b
}