package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	i, j := 0, 1
	return func() int{
		i, j = j, i + j
		//if i == 0 {
		  // i = 1
		//} else if j == 0 {
		  // j = 1
		//} else {
			//x := i
			//i = i + j
			//j = x
		//}
		return i 
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
