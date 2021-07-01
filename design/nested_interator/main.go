package main
import "fmt"

type NestedIterator struct {
	vals []int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	var vals []int
	var dfs func([]*NestedInteger)

	dfs = func(nestedList []*NestedInteger){

		// append the element recursively into vals
		for _,val := range nestedList{
			if val.IsInteger{
				vals = append(vals,val)
			}else{
				dfs(val)
			}
		}
	}

	dfs(nestedList)
	return &NestedIterator{
		vals:vals,
	}
    
}

func (this *NestedIterator) Next() int {
	ele := this.vals[0]
	this.vals = this.vals[1:]
	return ele
}

func (this *NestedIterator) HasNext() bool {
	return len(this.vals) == 0
    
}

func main(){

}