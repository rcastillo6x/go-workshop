# Slide | [Home](../../README.md)

Go slices are very powerful and it would not be an exaggeration to say that slices could totally replace the use of arrays in Go. There are only a few occasions when you will need to use an array instead of a slice. The most obvious one is when you are absolutely sure that you will need to store a fixed number of elements.

***
**Tip**

Slices are implemented using arrays internally, which means that Go uses an underlying array for each slice.
***


As slices are passed by reference to functions, which means that what is actually passed is the memory address of the slice variable, any modifications you make to a slice inside a function will not get lost after the function exits. Additionally, passing a big slice to a function is significantly faster than passing an array with the same number of elements because Go will not have to make a copy of the slice; it will just pass the memory address of the slice variable

## Performing basic operations on slices

You can create a new slice literal as follows:
```

aSliceLiteral := []int{1, 2, 3, 4, 5} 
 
```

However, there is also the make() function, which allows you to create empty slices with the desired length and capacity based on the parameters passed to make(). The capacity parameter can be omitted. In that case, the capacity of the slice will be the same as its length. So, you can define a new empty slice with 20 places that can be automatically expanded when needed as follows:

```
    integer := make([]int, 20) 
```

Please note that Go automatically initializes the elements of an empty slice to the zero value of its type, which means that the value of the initialization depends on the type of the object stored in the slice. It is good to know that Go initializes the elements of every slice created with make.
You can access all the elements of a slice in the following way:

```
    for i := 0; i < len(integer); i++ { 
        fmt.Println(integer[i]) 
    } 
```

If you want to empty an existing slice, the zero value for a slice variable is nil:

```
aSliceLiteral = nil 
```

You can add an element to the slice, which will automatically increase its size, using the append() function:

```
integer = append(integer, 12345) 
```

You can access the first element of the integer slice as integer[0], whereas you can access the last element of the integer slice as integer[len(integer)-1].

Lastly, you can access multiple continuous slice elements using the [:] notation. The next statement selects the second and the third elements of a slice:

```
integer[1:3] 
```

Additionally, you can use the [:] notation for creating a new slice from an existing slice or array:

```
s2 := integer[1:3] 
```

## len() and cap() functions
A slice is an abstraction over array. It actually uses arrays as an underlying structure. The len() function returns the elements presents in the slice where cap() function returns the capacity of the slice (i.e., how many elements it can be accommodate)

## Subslicing

Slice allows lower-bound and upper bound to be specified to get the subslice of it using[lower-bound:upper-bound]. For example −

```
package main

import "fmt"

func main() {
   /* create a slice */
   numbers := []int{0,1,2,3,4,5,6,7,8}   
   printSlice(numbers)
   
   /* print the original slice */
   fmt.Println("numbers ==", numbers)
   
   /* print the sub slice starting from index 1(included) to index 4(excluded)*/
   fmt.Println("numbers[1:4] ==", numbers[1:4])
   
   /* missing lower bound implies 0*/
   fmt.Println("numbers[:3] ==", numbers[:3])
   
   /* missing upper bound implies len(s)*/
   fmt.Println("numbers[4:] ==", numbers[4:])
   
   numbers1 := make([]int,0,5)
   printSlice(numbers1)
   
   /* print the sub slice starting from index 0(included) to index 2(excluded) */
   number2 := numbers[:2]
   printSlice(number2)
   
   /* print the sub slice starting from index 2(included) to index 5(excluded) */
   number3 := numbers[2:5]
   printSlice(number3)
   
}
func printSlice(x []int){
   fmt.Printf("len = %d cap = %d slice = %v\n", len(x), cap(x),x)
}
```

## Byte slices

A byte slice is a slice where its type is byte. You can create a new byte slice named s as follows:

```
s := make([]byte, 5) 
```

Go knows that most slices of bytes are used to store strings and so makes it easy to switch between this type and the string type. There is nothing special in the way you can access a byte slice compared to the other types of slices. It is just that byte slices are used in file input and output operations

## The copy() function

You can create a slice from the elements of an existing array and you can copy an existing slice to another one using the copy() function. However, as the use of copy()

```
package main 
 
import ( 
    "fmt" 
) 
 
func main() { 
    a6 := []int{-10, 1, 2, 3, 4, 5} 
    a4 := []int{-1, -2, -3, -4} 
    fmt.Println("a6:", a6) 
    fmt.Println("a4:", a4) 
 
    copy(a6, a4) 
    fmt.Println("a6:", a6) 
    fmt.Println("a4:", a4) 
    fmt.Println() 
    
```

So, in the preceding code, we define two slices named a6 and a4, we print them, and then we try to copy a4 to a6. As a6 has more elements than a4, all the elements of a4 will be copied to a6. However, as a4 has only four elements and a6 has six elements, the last two elements of a6 will remain the same.

## append()  
One can increase the capacity of a slice using the append() function.

```
package main

import "fmt"

func main() {
   var numbers []int
   printSlice(numbers)
   
   /* append allows nil slice */
   numbers = append(numbers, 0)
   printSlice(numbers)
   
   /* add one element to slice*/
   numbers = append(numbers, 1)
   printSlice(numbers)
   
   /* add more than one element at a time*/
   numbers = append(numbers, 2,3,4)
   printSlice(numbers)
   
   /* create a slice numbers1 with double the capacity of earlier slice*/
   numbers1 := make([]int, len(numbers), (cap(numbers))*2)
   
   /* copy content of numbers to numbers1 */
   copy(numbers1,numbers)
   printSlice(numbers1)   
}
func printSlice(x []int){
   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
```