package main
import "fmt"

func spiralOrder(matrix [][]int) []int {
	if len(matrix) ==0{
        return []int{}
    }
	res := []int{}
	direction := 0
	M:= len(matrix)
	N:= len(matrix[0])

	//left
	L := 0
	//right
	R:= N -1
	//bottom
	B:= M-1
	//top
	T:= 0

	for L<=R && T<=B{
		if direction == 0{
			//Top level (first row)
			for i:=L;i<=R;i++{
				res = append(res,matrix[T][i])
			}
			T++
		}else if direction == 1{

			//rightmost column
			for i:=T;i<=B;i++{
				res = append(res,matrix[i][R])
				
			}
			R--
		}else if direction ==2{
			//bottom line
			for i:= R;i>=L;i--{
				res = append(res,matrix[B][i])
			}
			B--

		}else if direction ==3{
			//leftmost cols 
			for i:= B;i>=T;i--{
				res = append(res,matrix[i][L])
			}
			L++
		}

		direction = (direction +1)%4
	}

	return res

}
 
func main(){
	matrix := [][]int {{1,2,3},{4,5,6},{7,8,9}}
	fmt.Println(spiralOrder(matrix))

}