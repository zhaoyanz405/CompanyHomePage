package main

import "fmt"

func main() {
	b := [...]string{1: "a", 2: "b", 3: "C"}
	c := b[0:1]
	fmt.Printf("%d", cap(c))
	fmt.Printf("%d", len(c))
}
