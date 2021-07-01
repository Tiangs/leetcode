package main
import "fmt"
import "sort"
func main(){
	arr:=[]int{1,2,2}
	fmt.Println(subsetsWithDup(arr))

}
func subsetsWithDup(nums []int) (ans [][]int) {
    sort.Ints(nums)
    t := []int{}
    var dfs func(bool, int)
    dfs = func(choosePre bool, cur int) {

		//primary return condition
        if cur == len(nums) {
            ans = append(ans, append([]int(nil), t...))
            return
        }
        dfs(false, cur+1)

		//cutting branch condition
        if !choosePre && cur > 0 && nums[cur-1] == nums[cur] {
            return
        }
        t = append(t, nums[cur])
        dfs(true, cur+1)
        t = t[:len(t)-1]
    }
	fmt.Println(t)
    dfs(false, 0)
    return
}
