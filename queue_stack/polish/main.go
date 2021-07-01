package main

import "fmt"
import "strconv"
func main(){
	a,b:=13,4
	fmt.Println(a/b)

}
func evalRPN(tokens []string) int {

	stack:=[]int{}
	for _,v := range tokens{
		token,err:=strconv.Atoi(v)
		
		if err == nil{ //number
			stack = append(stack,token)
		}else{ // operator
			num2,num1:= stack[len(stack)-1],stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			switch v{
			case "+":
				stack = append(stack,num1+num2)
			case "-":
				stack = append(stack, num1-num2)

			case "*":
				stack = append(stack, num1*num2)
			default:
				stack = append(stack, num1 / num2)
			}
		}
	}
	return stack[0]

 
}