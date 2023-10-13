package main

import "fmt"

func main() {
	x := 10
	if x >= 5 {
		fmt.Println(x) // 10
		x := 5
		fmt.Println(x) // 5

	}
	fmt.Println(x) // 10

	// shadowing with multiple assign
	{
		x := 10
		if x > 5 {
			x, y := 5, 20
			fmt.Println(x, y) // 5 20
		}
		fmt.Println(x) // 10
	}

	//!! shadowing the universe variables
	fmt.Println(true) // true
	true := "lol"
	fmt.Println(true) // lol

}
