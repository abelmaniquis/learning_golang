//https://gist.github.com/LordZamy/2adcb6d879fcef557d3d

package main

import (
	"fmt"
)

func main() {
	A := []int{3, 4, 2, 1, 4, 5, 7, 12, 6, 19, 4}
	fmt.Println(sort(A))
}

func sort() {
	if len(A) <= 1 {
		return A
	}

	left, right := split(A)
	left = sort(left)
	right = sort(right)
	return merge(left, right)

}

func split(A []int) ([]int, []int) {
	return A[0 : len(A)/2], A[len(A)/2:]
}

func merge() {

}
