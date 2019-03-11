package main

import (
	"fmt"
)

type MyError interface {
	message() string
}

func add(x, y int) (int, string) {
	if x == y {
		return x + y, "The same"
	}
	return x + y, "Different"
}

func main() {
	fmt.Printf("hello, world\n")
	r, s := add(100, 100)
	fmt.Printf("Res: %d - %s", r, s)
}
