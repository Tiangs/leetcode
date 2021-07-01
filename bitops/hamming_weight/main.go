
package main
import "fmt"
func main(){
    fmt.Println(hammingWeight(76))
}

func hammingWeight(num uint32) (ones int) {
    for i := 0; i < 32; i++ {

        if 1<<i&num > 0 {
            fmt.Println(1<<i)
            ones++
        }
    }
    return
}

 