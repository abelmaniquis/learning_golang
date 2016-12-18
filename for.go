package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

}

/*

for is Go's only looping construct. There are three different types
of for loops

The most basic type, with a single condition.

a classic initial/condition/after for loop

for without a condition will loop repeatedly until
you break out of a loop or return from the enclosing
function

We’ll see some other for forms later when we look at range statements,
channels, and other data structures.

*/
