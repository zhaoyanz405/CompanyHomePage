package main

import "fmt"

func main() {
	p := new(int)
	fmt.Println(*p)
	fmt.Println(p)
	fmt.Println(&p)

	x := 1
	fmt.Println(x)
	fmt.Println(&x)
	fmt.Println(*&p)
}
