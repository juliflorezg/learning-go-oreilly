package main

import (
	"fmt"
)

func main() {

	//> arrays
	var x [3]int // zero values -> [0 0 0]
	fmt.Println(x)

	var y = [3]int{10, 20, 30}
	fmt.Println(y)

	//sparse array with some of the values being zero values, the syntax 5: and 10: means that it will populate those positions up to that one specified (5 or 10) but it's not inclusive.
	// another way to see it is that we specify the values in order and we can specify a position and its value, position 5 will be 4, and position 10 will be 100, the ones we don't specify will be zero values
	var z = [12]int{1, 5: 4, 6, 10: 100, 15} // [1, 0, 0, 0, 0, 4, 6, 0, 0, 0, 100, 15]
	fmt.Println(z)

	var w = [...]int{2, 3, 4, 5}
	fmt.Println(w)
	var a = [4]int{2, 3, 4, 5}
	fmt.Println(w == a) // prints true

	// we can simulate multidimensional arrays with:
	var multiArray [2][3]int
	fmt.Println(multiArray)

	// for reading in an array, we use bracket syntax
	fmt.Println(a[0])

	//> slices
	var slice1 = []int{1, 2, 3}
	fmt.Println(slice1)
	fmt.Println("slice1 length::", len(slice1))
	var slice2 = []int{1, 5: 4, 6, 10: 100, 15} // index 5 will be 4, index 10 will be 100
	fmt.Println(slice2)                         //[1 0 0 0 0 4 6 0 0 0 100 15]

	// multidimensional slices:
	var multiSlice [][]int
	fmt.Println(multiSlice)

	//reading from slices
	slice2[0] = 10
	fmt.Println(slice2[5])

	// declare a slice without using a literal \
	var nilSlice []int
	fmt.Println(nilSlice) // []
	fmt.Println("nilSlice length::", len(nilSlice))
	fmt.Println(nilSlice == nil) // true

	nilSlice = append(nilSlice, 66)
	fmt.Println(nilSlice) // [66]

	//append more than one value at a time
	slice1 = append(slice1, 4, 5, 6)

	slice3 := []int{7, 8, 9}
	slice1 = append(slice1, slice3...)
	fmt.Println(slice1) // [1 2 3 4 5 6 7 8 9]

	// Let's see how capacity grows automatically in Go
	var x1 []int
	fmt.Println(x1, len(x1), cap(x1))
	x1 = append(x1, 10)
	fmt.Println(x1, len(x1), cap(x1))
	x1 = append(x1, 20)
	fmt.Println(x1, len(x1), cap(x1))
	x1 = append(x1, 30)
	fmt.Println(x1, len(x1), cap(x1))
	x1 = append(x1, 40)
	fmt.Println(x1, len(x1), cap(x1))
	x1 = append(x1, 50)
	fmt.Println(x1, len(x1), cap(x1))

	// the output will be like this:
	/*
		[] 0 0
		[10] 1 1
		[10 20] 2 2
		[10 20 30] 3 4
		[10 20 30 40] 4 4
		[10 20 30 40 50] 5
	*/
	// every time the slice runs out of capacity, the go runtime will double its capacity (it its less than 1024), beyond that will grow at least 25%
	//? make for slices

	y1 := make([]int, 5) // make a new slice of int with length 5, all the values will be zero values
	fmt.Println(y1)      // [0 0 0 0 0]

	// we can specify a capacity when making a slice
	y2 := make([]int, 5, 10)
	fmt.Println(y1, len(y2), cap(y2)) // [0 0 0 0 0] 5 10

	// non nil slice with zero length and defined capacity
	y3 := make([]int, 0, 10)
	fmt.Println(y3)        // []
	fmt.Println(y3 == nil) // false

}
