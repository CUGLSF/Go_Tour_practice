// exercise-fibonicci-closure.go
package main

import "fmt"

// fibonacci 函数会返回一个返回 int 的函数。
func fibonacci() func() int {
	f_zero := 0
	f_one := 1
	return func() int {
		f_zero, f_one = f_one, f_zero+f_one
		return f_zero
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		if i == 0 {
			fmt.Println(0)
		}
		fmt.Println(f())
	}
}
