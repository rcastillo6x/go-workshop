package main

import "fmt"

func main() {
	var numbersSimple = make([]int, 3, 5)
	printSlice(numbersSimple, "numberSimple")

	numbersSimple = nil

	if numbersSimple == nil {
		fmt.Printf("slice is nil")
	}

	// Part 2
	/* create a slice */
	numbersInt := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbersInt, "numbersInt")

	/* print the original slice */
	fmt.Println("numbers ==", numbersInt)

	/* print the sub slice starting from index 1(included) to index 4(excluded)*/
	fmt.Println("numbers[1:4] ==", numbersInt[1:4])

	/* missing lower bound implies 0*/
	fmt.Println("numbers[:3] ==", numbersInt[:3])

	/* missing upper bound implies len(s)*/
	fmt.Println("numbers[4:] ==", numbersInt[4:])

	numbers1 := make([]int, 0, 5)
	printSlice(numbers1, "number1")

	/* print the sub slice starting from index 0(included) to index 2(excluded) */
	number2 := numbersInt[:2]
	printSlice(number2, "number2")

	/* print the sub slice starting from index 2(included) to index 5(excluded) */
	number3 := numbersInt[2:5]
	printSlice(number3, "number3")
}
func printSlice(x []int, name string) {
	fmt.Printf("slice = %s len=%d cap=%d slice=%v\n", name, len(x), cap(x), x)
}
