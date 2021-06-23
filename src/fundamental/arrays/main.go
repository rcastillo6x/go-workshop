package main

import (
	"fmt"
)

func main() {
	var n [10]int /* n is array of 10 integers */
	var i, j int

	for i = 0; i < 10; i++ {
		n[i] = i + 100 /* set element */
	}

	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}
	// Two dimension
	/* an array with 5 rows and 2 columns*/
	var a = [5][2]int{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}}
	var i, j int

	/* output each array element's value */
	for i = 0; i < 5; i++ {
		for j = 0; j < 2; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}

	// Three Dimensions

	oneD := [4]int{1, 2, 4, -4}
	twoD := [4][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	threeD := [2][2][2]int{{{1, 0}, {-2, 4}}, {{5, -1}, {7, 0}}}

	fmt.Println("The length of", oneD, "is", len(oneD))
	fmt.Println("The first element of", twoD, "is", twoD[0][0])
	fmt.Println("The length of", threeD, "is", len(threeD))

	fmt.Println("\nArrays of 3 dimensions- Form 1:")

	for i := 0; i < len(threeD); i++ {
		v := threeD[i]
		for j := 0; j < len(v); j++ {
			m := v[j]
			for k := 0; k < len(m); k++ {
				fmt.Print("x: ", i, " y: ", j, " z:", k, " value: ", m[k], " , ")
			}
		}
		fmt.Println()
	}

	fmt.Println("\nArrays of 3 dimensions. Form 2:")
	for _, v := range threeD {
		for _, m := range v {
			for _, s := range m {
				fmt.Print(s, " ")
			}
		}
		fmt.Println()
	}

}
