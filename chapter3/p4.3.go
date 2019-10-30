package main

import (
	"fmt"
)

func reverse1(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverse2(s *[3]string) {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	a := new([3]string)
	*a = [3]string{"a", "b", "c"}
	reverse2(a)
	fmt.Println(*a)

	b := [3]string{"a", "b", "c"}
	reverse2(&b)
	fmt.Println(b)

}
