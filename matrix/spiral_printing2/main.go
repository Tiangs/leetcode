package main
import "fmt"

func generateMatrix(n int) [][]int {

	if n==1{
        return [][]int{{1}}
    }

	// res := make([][]int , n)
	res:= make([][]int, n)       // initialize a slice of dy slices
	for i:=0;i<n;i++ {
		res[i] = make([]int, n)  // initialize a slice of dx unit8 in each of dy slices
	}

	direction := 0

	//left
	L := 0
	//right
	R:= n -1
	//bottom
	B:= n-1
	//top
	T:= 0
	count :=1 

	for L<=R && T<=B{
		if direction == 0{
			//Top level (first row)
			for i:=L;i<=R;i++{
				// res = append(res,matrix[T][i])
				res[T][i] = count
				count++
			}
			T++
		}else if direction == 1{

			//rightmost column
			for i:=T;i<=B;i++{
				// res = append(res,matrix[i][R])
				res[i][R] = count
				count++
			}
			R--
		}else if direction ==2{
			//bottom line
			for i:= R;i>=L;i--{
				// res = append(res,matrix[B][i])
				res[B][i] = count
				count++
			}
			B--

		}else if direction ==3{
			//leftmost cols 
			for i:= B;i>=T;i--{
				// res = append(res,matrix[i][L])
				res[i][L] = count
				count++
			}
			L++
		}

		direction = (direction +1)%4
	}

	return res

}
 
func main(){
	fmt.Println(generateMatrix(4))
}