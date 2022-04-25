package main

import (
	"fmt"
	"math/cmplx"
)

func main() {

	x := complex(5, 8)
	y := complex(3, 4)
	mod_x := cmplx.Abs(x)
	mod_y := cmplx.Abs(y)
	conj_x := cmplx.Conj(x)
	phase_x := cmplx.Phase(x)
	mod, phase := cmplx.Polar(x)

	fmt.Println("x = ", x)
	fmt.Println("Modulus of x = ", mod_x)
	fmt.Println("Modulus of y = ", mod_y)
	fmt.Println("Conjugate of x = ", conj_x)
	fmt.Println("Phase of x = ", phase_x)
	fmt.Printf("Polar Form : %v, %v\n", mod, phase)

}
