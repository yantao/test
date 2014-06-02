package main

import (
	"fmt"
)

const PI = 3.14

var name = "yt"

type newType int

type mystruct struct {
}

type golang interface{}

func main() {
	//fmt.Println("Hello world 书剑")
	//a := [2]int{}
	//fmt.Println(a)

	// a := [2]int{1, 2}
	// b := [2]int{2, 4}
	// fmt.Println(a == b)
	// fmt.Println(len(a))

	// fmt.Fprintln(a.size())

	a := [...]int{1, 234, 89, 2314, 92, 784, 2}
	fmt.Println(a)

	num := len(a)

	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if a[i] > a[j] {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
	fmt.Println(a)

	fmt.Println(a)
}
