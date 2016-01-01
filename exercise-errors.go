// exercise-errors.go
package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	// 为了避免可能出现的无穷递归，应在递归之前转换值的类型
	return fmt.Sprint("cannot Sqrt negative number: ", float64(e))
}
func Sqrt(x float64) (float64, error) {
	z := 1.0
	if x >= 0 {
		for {
			if math.Abs(z*z-x) > 0.01 {
				z = z - (z*z-x)/2*z
			} else {
				break
			}
		}
		return z, nil
	} else {
		return 0, ErrNegativeSqrt(x)
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
