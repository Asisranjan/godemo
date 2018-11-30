package main 

import "fmt"

func main() {
	num1,num2 := 5.6, 6.4

	fmt.Println(add(num1, num2))
}

func add(x float64,y float64) float64{
	return x + y
}