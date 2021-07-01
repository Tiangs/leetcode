package main
import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	if matrix == nil{
		return false
	}
	m,n := len(matrix),len(matrix[0])
	left,right := 0,m*n-1

	for left<= right{
		mid := (right-left)/2+left
		val := matrix[mid/n][mid%n]
		if target<val{
			right = mid-1
		}else if target>val{
			left = mid +1
		}else{
			return true
		}

	}
	return false
}

func main(){
	matrix:= [][]int{{1,3,5,7},{10,11,16,20},{23,30,34,60}}
	fmt.Println(searchMatrix(matrix,3))

}


