package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a bool = true
	var b int32
	b = 111
	c := false
	fmt.Println(b)
	fmt.Println(a)
	fmt.Println(c)

	z, x, t := 1, 2, 3
	fmt.Println(z)
	fmt.Println(x)
	fmt.Println(t)

	var zz float64 = 1.1
	var bb = int(zz)
	fmt.Println(bb)

	var aaa int = 99
	var sss = strconv.Itoa(aaa)
	fmt.Println(sss)
}
