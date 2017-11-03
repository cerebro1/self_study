package main

import "fmt"


func add( num1, num2 int) (num4, num3 int) {
	num3 = num1 + num2
	num4 = num1 + num3
	return
}

func main(){
	fmt.Println(add(10,20))
}
