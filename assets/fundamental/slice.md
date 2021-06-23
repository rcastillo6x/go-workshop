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
