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

	//! we should never specify a capacity that's less that the length

	//? slicing
	// we can get a slice from another slice
	x2 := []int{1, 2, 3, 4}
	y4 := x2[:2] // [1 2] end index is not inclusive
	z1 := x2[1:] // [2 3 4] from index one (inclusive) to the end of the slice
	d := x2[1:3] // [2 3] initial index inclusive, end index not inclusive
	e := x2[:]   // [1 2 3 4] the entire slice, since no index where passed
	fmt.Println(x2)
	fmt.Println(y4)
	fmt.Println(z1)
	fmt.Println(d)
	fmt.Println(e)

	// when we slice from an existing slice, Go doesn't create a copy of those elements for the new slice, it shares memory between the original slice and the new sub-slice created, meaning that any change in one of the elements will be reflected in the parent slice or sub-slice that contains those elements:::

	originalSlice := []int{1, 2, 3, 4}
	firstSubSlice := originalSlice[:2]  // [1 2]
	secondSubSlice := originalSlice[3:] // [4]

	fmt.Println(cap(originalSlice), cap(firstSubSlice), cap(secondSubSlice)) // 4 4 1
	fmt.Println("originalSlice", originalSlice)
	fmt.Println("firstSubSlice", firstSubSlice)
	fmt.Println("secondSubSlice", secondSubSlice)

	firstSubSlice = append(firstSubSlice, 30)
	fmt.Println("originalSlice after append", originalSlice) // [1 2 30 4]
	fmt.Println("firstSubSlice after append", firstSubSlice) // [1 2 30 ]

	// wtf, why the 3rd element in the original slice was modified ?
	// bc when we slice from another slice, the sub-slice capacity is set to the original slices capacity, minus the offset of the sub-slice on the parent slice (how far from the beginning is the sub-slice)
	// When we make the firstSubSlice slice from originalSlice, the length is set to 2, but the capacity is set to 4, the same as originalSlice. Since the capacity is 4, appending onto the end of firstSlice puts the value in the third position of originalSlice

	//* a more confusing example :::

	originalSlice2 := make([]int, 0, 5)
	originalSlice2 = append(originalSlice2, 1, 2, 3, 4)
	firstSubSlice2 := originalSlice2[:2]
	secondSubSlice2 := originalSlice2[2:]
	fmt.Println(cap(originalSlice2), cap(firstSubSlice2), cap(secondSubSlice2)) // 5 5 3  -> 3 bc of the originalSlice2 capacity and the offset (2 positions)
	fmt.Println((originalSlice2), (firstSubSlice2), (secondSubSlice2))

	firstSubSlice2 = append(firstSubSlice2, 30, 40, 50)
	originalSlice2 = append(originalSlice2, 60)
	secondSubSlice2 = append(secondSubSlice2, 70)
	fmt.Println("originalSlice2 after append", originalSlice2)   // [1 2 30 40 70]
	fmt.Println("firstSubSlice2 after append", firstSubSlice2)   // [1 2 30 40 70]
	fmt.Println("secondSubSlice2 after append", secondSubSlice2) // [1 2 30 40 70]

	//* in order to avoid this confusing situations, we must not, (NEVER EVER swear by the most holy thing) use append with a sub-slice or use a full slice expression to make sure that append doesn't cause any overwrite

	// the full slice expression user a third parameter, it tells us the last position in the parent slice's capacity that's available for the sub-slice

	originalSlice3 := make([]int, 0, 5)
	originalSlice3 = append(originalSlice3, 1, 2, 3, 4)
	firstSubSlice3 := originalSlice3[:2:2]                                      //> here we tell that parents capacity will be available up to index 2 (non inclusive) for the sub-slice
	secondSubSlice3 := originalSlice3[2:4:4]                                    //> the same here, but it will be up to index 4
	fmt.Println(cap(originalSlice3), cap(firstSubSlice3), cap(secondSubSlice3)) // 5 2 2  -> we've changed the sub-slices capacity to 2 with the full slice expression
	fmt.Println((originalSlice3), (firstSubSlice3), (secondSubSlice3))

	firstSubSlice3 = append(firstSubSlice3, 30, 40, 50)
	originalSlice3 = append(originalSlice3, 60)
	secondSubSlice3 = append(secondSubSlice3, 70)
	fmt.Println("originalSlice3 after append", originalSlice3)   // [1 2 3 4 60]
	fmt.Println("firstSubSlice3 after append", firstSubSlice3)   // [1 2 30 40 50]
	fmt.Println("secondSubSlice3 after append", secondSubSlice3) // [3 4 70]

	// we can convert an array to an slice and we can also slice an array, and the same sharing memory issues can be found if we're not careful

	originalArray := [4]int{5, 6, 7, 8}
	sliceFromArray1 := originalArray[:2] // [5 6]
	sliceFromArray2 := originalArray[2:] // [7 8]
	originalArray[0] = 100
	fmt.Println("originalArray:", originalArray)     // [100 6 7 8]
	fmt.Println("sliceFromArray1:", sliceFromArray1) // [100 6]
	fmt.Println("sliceFromArray2:", sliceFromArray2) // [7 8]

	// to  make an independent copy of an slice, we must use the built-in function copy

	sliceToCopy := []int{1, 2, 3, 4}
	resultSlice := make([]int, 4)

	//> the following would be read as -copy in resultSlice as many items as possible from sliceToCopy-
	numOfItemsCopied := copy(resultSlice, sliceToCopy) // the -copy- function takes two parameters, first one is the destination slice and second one is the source slice, and it returns the number of items copied.

	fmt.Println("numOfItemsCopied", numOfItemsCopied) //4
	fmt.Println("resultSlice", resultSlice)           // [1 2 3 4]

	//copying just a portion of the source slice
	resultSlice2 := make([]int, 2)
	numOfItemsCopied2 := copy(resultSlice2, sliceToCopy)

	fmt.Println("numOfItemsCopied2", numOfItemsCopied2) // 2
	fmt.Println("resultSlice2", resultSlice2)           // [1 2]

	//copying from the middle of the source slice
	resultSlice3 := make([]int, 2)
	copy(resultSlice3, sliceToCopy[2:]) //* we can not assign the return value of copy to a variable if we don't need it

	fmt.Println("resultSlice3", resultSlice3) // [3 4]

	// this is also possible (page 47 in the book)

	{
		x := []int{1, 2, 3, 4}
		num := copy(x[:3], x[1:])
		fmt.Println(x, num) // [2 3 4 4] 3
	}
	// that block right there shows us how to copy a section from a slice to another section in the same slice 🤯, here we copied the last three numbers from that slice into the first three numbers in the same slice

	// we can also use copy in arrays, by slicing them:
	{
		fmt.Println("using copy in arrays")
		x := []int{1, 2, 3, 4}
		d := [4]int{5, 6, 7, 8}
		y := make([]int, 2)
		copy(y, d[:])
		fmt.Println(y) // [5 6]
		copy(d[:], x)
		fmt.Println(d) // [1 2 3 4]
	}

	//> strings and runes
	{
		var s string = "Hello there"
		var b byte = s[6]   // byte for 't'
		fmt.Println("b", b) // 116
		var s2 string = s[4:7]
		fmt.Println("s2::", s2) // o t
		var s3 string = s[:5]
		fmt.Println("s3::", s3) // Hello
		var s4 string = s[6:]
		fmt.Println("s4::", s4) // there
	}
	// code points that are multiple bytes long
	{
		var s string = "Hello 🌞"
		fmt.Println(len(s)) // 10 (10 bytes long)
		var s2 string = s[4:7]
		fmt.Println("s2::", s2) // o �
		var s3 string = s[:5]
		fmt.Println("s3::", s3) // Hello
		var s4 string = s[6:]
		fmt.Println("s4::", s4) // 🌞
		var b byte = s[6]
		var b1 byte = s[7]
		var b2 byte = s[8]
		var b3 byte = s[9]
		fmt.Println(b)
		fmt.Println(b1)
		fmt.Println(b2)
		fmt.Println(b3)
	}
	// go allows type conversion between runes, strings and bytes
	{
		var a rune = 'x'          // 120
		var s string = string(a)  // string from a rune
		var b byte = 'y'          // 121
		var s2 string = string(b) // string from a byte
		fmt.Println("a::", a)
		fmt.Println(s)
		fmt.Println(b)
		fmt.Println(s2)

	}
	// we can't convert a number to a string ie  5 -> "5"
	{
		var x int = 65
		var y = string(x) //!! not "65" but "A"
		fmt.Println(y)
		s := fmt.Sprint(x)
		fmt.Println(s)
		fmt.Printf("Type of s (%v) is:: %T", s, s)
	}

	//we can convert a string in slices of runes and bytes
	{
		var s string = "Hello 🌞 yay"
		var bs []byte = []byte(s) // byte slice -> [72 101 108 108 111 32 240 159 140 158 32 121 97 121]
		var rs []rune = []rune(s) // rune slice -> [72 101 108 108 111 32 127774 32 121 97 121]
		fmt.Println(bs)
		fmt.Println(rs)
	}

	//> maps
	// declare a empty map and set it to its zero value(nil)

	var nilMap map[string]int
	fmt.Println("nilMap", nilMap) // map[]
	fmt.Println(nilMap == nil)    // true
	// A nil map has length zero, we can read an nil map and will return its zero value for the map's value type, but if we try to write to a nil mal it will cause a panic.
	fmt.Println(nilMap["key1"]) // 0
	fmt.Println(nilMap["key2"]) // 0
	// nilMap["key1"] = 3 //! panic: assignment to entry in nil map

	// we can use the := syntax to declare an empty map, it will have length 0 but we can read and write from it

	emptyMap := map[string]int{}
	fmt.Println(emptyMap["one"]) // 0
	emptyMap["one"] = 1
	fmt.Println(emptyMap["one"]) // 1

	// non empty map with := syntax
	teams := map[string][]string{
		"Orcas":   []string{"Fred", "Ralph", "Bijou"},
		"Lions":   []string{"Sarah", "Peter", "Billie"},
		"Kittens": []string{"Waldo", "Raul", "Ze"},
	}

	fmt.Println(teams)

	// we can make an empty map with predefined size if we do know how many key-value pairs will it hold but not which ones
	ages := make(map[int][]string, 10)
	fmt.Println(ages)

	// we can use comparable types as keys for maps, this means we can't use slices or maps as keys, since we can't compare them using == or !=

	// this block show us how to read and write from an to a map

	{
		totalWins := map[string]int{}
		totalWins["Orcas"] = 1
		totalWins["Lions"] = 2
		fmt.Println(totalWins["Orcas"])   // 1
		fmt.Println(totalWins["Kittens"]) // 0 (zero value bc it hasn't been set and its type is int, so its zero value will be 0)
		totalWins["Kittens"]++            // we can use the increment operator to operate this zero value and increment its value to 1
		fmt.Println(totalWins["Kittens"]) // 1
		totalWins["Lions"] = 3
		fmt.Println(totalWins["Lions"]) // 3
	}

	//? the comma, ok idiom
	// it help us know if a key exists in a map

	m := map[string]int{
		"hello": 5,
		"world": 0,
	}

	v, ok := m["hello"]
	fmt.Println(v, ok) // 5 true
	v1, ok1 := m["world"]
	fmt.Println(v1, ok1) // 0 true
	v2, ok2 := m["bye"]
	fmt.Println(v2, ok2) // 0 false

	//* we can delete from a map using the delete built-in function
	{
		m := map[string]int{
			"hello": 5,
			"world": 10,
		}
		delete(m, "world")
		fmt.Println(m)
	}

	//* we can use a map as a set, by setting the key type to the type of values we want in the set and the value type to boolean, like this::
	{
		// page 55
		intSet := map[int]bool{}
		intList := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
		for _, v := range intList {
			intSet[v] = true
		}
		fmt.Println(intSet)
		fmt.Println(len(intList), len(intSet))
		fmt.Println(intSet[5])   // true
		fmt.Println(intSet[500]) // false (zero value for 500 key, is not in the list)
		if intSet[100] {
			fmt.Println("100 is in the list")
		}
	}

	//* we can also use struct instead of boolean for the set. The advantage is that an empty struct uses zero bytes whereas a boolean uses one byte. The disadvantage is that it makes the code harder to read and we have to use the comma, ok idiom to check if the value is in the set::

	{
		intSet := map[int]struct{}{}
		vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
		for _, v := range vals {
			intSet[v] = struct{}{}
		}
		fmt.Println(intSet[5])
		if _, ok := intSet[5]; ok {
			fmt.Println("5 is in the set")
		}
	}

}
