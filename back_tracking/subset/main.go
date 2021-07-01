package main
import "fmt"
func main(){
	arr := []int{1,2,3}
	fmt.Println(subsets(arr))
}
func subsets(nums []int) (ans [][]int) {
    set := []int{}
    var dfs func(int)
    dfs = func(level int) {
        if level == len(nums) {
			//put the current set into ans
            ans = append(ans, append([]int(nil), set...))
            return
        }
        set = append(set, nums[level]) //include current element
        dfs(level + 1)
        set = set[:len(set)-1] // not include current element
        dfs(level + 1)
    }
    dfs(0)
    return
}
