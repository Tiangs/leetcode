package main

func main(){
/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
}

type ParkingSystem struct {
	arr  []int
}


func Constructor(big int, medium int, small int) ParkingSystem {

	
	return ParkingSystem{
		arr:[]int{0,big,medium,small},
	}

}


func (this *ParkingSystem) AddCar(carType int) bool {
	if this.arr[carType]>0{
		this.arr[carType]--
		return true
	}
	return false

}


