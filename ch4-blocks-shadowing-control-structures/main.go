package main

import (
	"fmt"
	"math/rand"
)

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

	//> if statement

	n := rand.Intn(10)
	if n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That's a good number:", n)
	}

	// scoping a variable to an if statement

	if n := rand.Intn(10); n == 0 {
		// this variable n is shadowing the previous created at line 35 on function block
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That's a good number:", n)
	}

	// variable created on if statement block only lives within that block

	if num := rand.Intn(10); num == 0 {
		fmt.Println("That's too low")
	} else if num > 5 {
		fmt.Println("That's too big:", num)
	} else {
		fmt.Println("That's a good number:", num)
	}
	// fmt.Println(num) // undefined

	//* for statement (four ways)
	// there are 4 ways we can write a for loop in go

	//* 1. complete, conventional for loop

	// we can't use var to initialize the variable in for loop, it's ilegal
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//* 2. condition-only for loop

	// this one is similar to a while statement found in C, Java, JS, Python
	i := 1
	for i < 100 {
		fmt.Println("i:::", i)
		i = i * 2
	}

	//* 3. the infinite for loop

	// for {
	// 	fmt.Println("HELLO!!")
	// }

	//* break & do while
	// in go, we don't have a do, while statement like in some other languages. If we want something like that we must use a infinite loop with an if statement at the end:::

	// for {
	// 	// do stuff
	// 	if !condition{
	// 		break //? break exits the loop immediately
	// 	}
	// }

	//* continue
	// continue skips the rest of the body in the for loop and jumps to the next iteration
	// it helps us write idiomatic code

	//? from this
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			if i%5 == 0 {
				fmt.Println("FizzBuzz")
			} else {
				fmt.Println("Fizz")
			}
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

	//> to this
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		}
		if i%3 == 0 {
			fmt.Println("Fizz")
			continue
		}
		if i%5 == 0 {
			fmt.Println("Buzz")
			continue
		}
		fmt.Println(i)
	}
}
