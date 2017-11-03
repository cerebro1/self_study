package main

import "fmt"


func add( num1, num2 int) int {
	num3 := num1 + num2
	return num3
}

func main(){
	fmt.Println(add(10,20))
}
