package main

import (
	"fmt"
)

func main() {
	nedStark := func() {
		name, location, age, isDead := "Ned Stark", "Winterfell", 50, true

		fmt.Printf("%s (%d) of %s is dead?: %b \n", name, age, location, isDead)
	}
	robbStark := func() {
		name, location, age, isDead := "Robb Stark", "Winterfell", 18, true
		fmt.Printf("%s (%d) of %s is dead?: %b \n", name, age, location, isDead)
	}
	nedStark()
	robbStark()

}
