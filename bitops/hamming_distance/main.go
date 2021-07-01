package main

import "fmt"
func main(){
	fmt.Println(hammingDistance(1,4))

}

func hammingDistance(x int, y int) int {
	ans:=0

	//XOR to find the different bit
	t:= x^y

	//count when last bit is 1
	for t>0{
		ans+=t&1
		t>>=1
	}

	return ans
}