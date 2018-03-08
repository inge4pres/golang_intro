package main

import "fmt"

func main() {
	friends := "julia,mark,kent"

	// all my friends are going to be at my party, reference them all
	invited := &friends
	fmt.Println("invited ", *invited)

	// kent is not my friend anymore
	friends = "julia,mark"
	fmt.Println("invited ", *invited)

	//side-note
	fmt.Println(invited)
}
