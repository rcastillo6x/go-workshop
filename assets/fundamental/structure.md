# Structure | [Home](../../README.md)

A Go program basically consists of the following parts −

* Package Declaration
* Import Packages
* Functions
* Variables
* Statements and Expressions
* Comments

```
package main

import "fmt"

func main() {
	var msg string = "This is a message"
	/*
		This is a comment
	*/
	fmt.Println("Var msg is equal to :" + msg)
}
```