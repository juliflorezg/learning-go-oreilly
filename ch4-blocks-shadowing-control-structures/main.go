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
	// true := "lol"
	// fmt.Println(true) // lol

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

	//* 4. for-range statement

	// used for strings, slices, arrays and maps (also with channels)

	evenValues := []int{2, 4, 6, 8, 10, 12}

	for i, v := range evenValues {
		fmt.Println(i, v)
	}
	// for range without the index
	for _, v := range evenValues {
		fmt.Println(v)
	}
	//* if we want just the keys or indexes
	uniqueNames := map[string]bool{"John": true, "Miles": true, "Cerberus": true}
	for name := range uniqueNames {
		fmt.Println(name)
	}

	// for-range with maps

	map1 := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	//when we run this code, we don't always get the same result, the order differs, this is for security reasons (check page 74 for explanation)
	for i := 0; i < 3; i++ {
		fmt.Println("loop:::", i)
		for k, v := range map1 {
			fmt.Println(k, v)
		}
	}

	// for-range for strings
	samples := []string{"hello", "apple_π!"}
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
		}
		fmt.Println()
	}

	// if we modify the value in a for-range loop, it doesn't modify the source
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for _, v := range evenVals {
		v *= 2
	}
	fmt.Println(evenVals) // [2 4 6 8 10 12]

	// we can put a label on a for loop::
outer:
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
			if r == 'l' {
				// this continues the iteration on the 'outer' for loop, so it goes to the second word
				continue outer
			}
		}
	}

	//> switch statement
	words := []string{"a", "cow", "smile", "gopher",
		"octopus", "anthropologist"}
	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short word!")
		case 5:
			wordLen := len(word)
			fmt.Println(word, "is exactly the right length:", wordLen)
		case 6, 7, 8, 9:
		default:
			fmt.Println(word, "is a long word!")
		}
	}
	// outputs:
	// a is a short word!
	// cow is a short word!
	// smile is exactly the right length: 5
	// anthropologist is a long word!

	// to make it output smt for cases 6, 7, 8, 9:::

	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short word!")
		case 5:
			wordLen := len(word)
			fmt.Println(word, "is exactly the right length:", wordLen)
		case 6, 7, 8, 9:
			wordLen := len(word)
			fmt.Println(word, "has length:", wordLen)
		default:
			fmt.Println(word, "is a long word!")
		}
	}

	//* break in switch
	// used to get out a case, not necessary
	// another use is to break the current loop (we need to label the loop in order for this to work)

	//?this doesn't exit the loop properly:::
	for i := 0; i < 10; i++ {
		switch {
		case i%2 == 0:
			fmt.Println(i, "is even")
		case i%3 == 0:
			fmt.Println(i, "is divisible by three but not 2")
		case i%7 == 0:
			fmt.Println(i, "exit the loop!!")
			break
		default:
			fmt.Println(i, "boring!!!")
		}
	}

	//0 is even
	// 1 boring!!!
	// 2 is even
	// 3 is divisible by three but not 2
	// 4 is even
	// 5 boring!!!
	// 6 is even
	// 7 exit the loop!!
	// 8 is even
	// 9 is divisible by three but not 2

	//? this, on the other hand, exits the loop as intended
loop:
	for i := 0; i < 10; i++ {
		switch {
		case i%2 == 0:
			fmt.Println(i, "is even")
		case i%3 == 0:
			fmt.Println(i, "is divisible by three but not 2")
		case i%7 == 0:
			fmt.Println(i, "exit the loop!!")
			break loop
		default:
			fmt.Println(i, "boring!!!")
		}
	}

	// 	0 is even
	// 1 boring!!!
	// 2 is even
	// 3 is divisible by three but not 2
	// 4 is even
	// 5 boring!!!
	// 6 is even
	// 7 exit the loop!!

}
