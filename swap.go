package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("use the force", "you will")
	fmt.Println(a, b)
}
