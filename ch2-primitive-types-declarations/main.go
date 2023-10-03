package main

import "fmt"

func main() {
	// we can use underscore in the middle of a literal to, for example, group by thousands:
	num1 := 1_799

	str1 := "Greetings and \n\"salutations\""
	str2 := `Greetings and
"salutations"`

	fmt.Println(num1)
	fmt.Println(str1)
	fmt.Println(str2)

	// booleans
	var flag bool // since this doesn't have any value assigned it'll be false (its zero value)
	var isAwesome = true

	fmt.Println("flag:", flag)
	fmt.Println("isAwesome:", isAwesome)

	var var1 uint8 = 255 // integer from 0 to 255
	var var2 byte = 255  // same as uint8 but is better and more usual to find uint8 as byte, so we'll use byte

	fmt.Printf("type of var1: %T\n", var1)
	fmt.Printf("type of var2: %T\n", var2)

	// four usual arithmetic operators +, -, *, / and modulus with %

	// in order to divide this two, we have to make a type conversion
	var firstNumber = float64(23)
	var secondNumber = float64(44)
	var result = firstNumber + secondNumber
	fmt.Println(result)
	result = firstNumber - secondNumber
	fmt.Println(result)
	result = firstNumber * secondNumber
	fmt.Println(result)
	result = firstNumber / secondNumber
	fmt.Println(result)

	if firstNumber > 0 && secondNumber > 0 {
		fmt.Println("both are greater than zero")
	}
	if firstNumber > 0 || secondNumber > 0 {
		fmt.Println("at least one of them is greater than zero")
	}
	// just one of them, like logical XOR
	if (int(firstNumber)%2 == 0 || int(secondNumber)%2 == 0) && !(int(firstNumber)%2 == 0 && int(secondNumber)%2 == 0) {
		fmt.Println("just one of them is even")
	}
	fmt.Println(int(firstNumber) % 2)
	fmt.Println(int(secondNumber) % 2)

	// type conversion
	var x int = 15
	var y float64 = 2.5
	var z float64 = float64(x) + y
	var w int = x + int(y)

	fmt.Println("z:", z)
	fmt.Println("w:", w)

	// this is used to give the variable the zero value
	var x1 int
	fmt.Println(x1)
	// declaring multiple variables of the same type
	var first, second int = 10, 20
	fmt.Println("first", first)
	fmt.Println("second", second)
	//zero values of the same type
	var zero1, zero2 int
	fmt.Println("zero1", zero1)
	fmt.Println("zero2", zero2)
	//multiple variables with dif types
	var dif1, dif2 = 69, "hellooo"
	fmt.Println("dif1", dif1)
	fmt.Println("dif2", dif2)

	// declaring multiple variables at once
	var (
		x3     int
		y3         = 20
		z3     int = 30
		d3, e3     = 40, "hello"
		f, g   string
	)

	fmt.Println(x3, y3, z3, d3, e3, f, g)

	// using short declaration format
	f2 := 10
	f2, g2 := 22, 44
	f2, g2, h3 := 36, 54, 69
	fmt.Println(f2, g2, h3)

}
