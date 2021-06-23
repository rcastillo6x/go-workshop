# Arrays | [Home](../../README.md)

Go programming language provides a data structure called the array, which can store a fixed-size sequential collection of elements of the same type. 
An array is used to store a collection of data, but it is often more useful to think of an array as a collection of variables of the same type.

Instead of declaring individual variables, such as number0, number1, ..., and number99, you declare one array variable such as numbers and use numbers[0], numbers[1], and ..., numbers[99] to represent individual variables. A specific element in an array is accessed by an index.

All arrays consist of contiguous memory locations. The lowest address corresponds to the first element and the highest address to the last element

## Declaring Arrays

To declare an array in Go, a programmer specifies the type of the elements and the number of elements required by an array as follows −

```
    var variable_name [SIZE] variable_type
```

This is called a single-dimensional array. The arraySize must be an integer constant greater than zero and type can be any valid Go data type. For example, to declare a 10-element array called balance of type float32, use this statement −

```
    var balance [10] float32
```

Here, balance is a variable array that can hold up to 10 float numbers.

## Initializing Arrays

You can initialize array in Go either one by one or using a single statement as follows −

```

    var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
 
```

The number of values between braces { } can not be larger than the number of elements that we declare for the array between square brackets [ ].

If you omit the size of the array, an array just big enough to hold the initialization is created. Therefore, if you write −

```

 var balance = []float32{1000.0, 2.0, 3.4, 7.0, 50.0}

 ```

## Accessing Array Elements

An element is accessed by indexing the array name. This is done by placing the index of the element within square brackets after the name of the array. For example −

```
float32 salary = balance[9] 
```

## Multi-dimensional Arrays

Arrays can have more than one dimension. However, using more than three dimensions without a serious reason can make your program difficult to read and might create bugs.

Go programming language allows multidimensional arrays. Here is the general form of a multidimensional array declaration −

```

var variable_name [SIZE1][SIZE2]...[SIZEN] variable_type
 
```

The following Go code shows how you can create an array with two dimensions (twoD) and another one with three dimensions (threeD)

```

twoD := [4][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12},  {13, 14, 15, 16}} 
threeD := [2][2][2]int{{{1, 0}, {-2, 4}}, {{5, -1}, {7, 0}}} 
 
```

## Initializing Two-Dimensional Arrays

Multidimensional arrays may be initialized by specifying bracketed values for each row. Following is an array with 3 rows and each row has 4 columns.

```
a = [3][4]int{  
   {0, 1, 2, 3} ,   /*  initializers for row indexed by 0 */
   {4, 5, 6, 7} ,   /*  initializers for row indexed by 1 */
   {8, 9, 10, 11}   /*  initializers for row indexed by 2 */
}
```

## Accessing Two-Dimensional Array elements

An element in two dimensional array is accessed by using the subscripts, i.e., row index and column index of the array. For example −

int val = a[2][3]

Accessing, assigning, or printing a single element from one of the previous two arrays can be done easily. As an example, the first element of the twoD array is twoD[0][0] and its value is 1.

Therefore, accessing all the elements of the threeD array with the help of multiple for loops can be done as follows:

```
    for i := 0; i < len(threeD); i++ { 
        for j := 0; j < len(v); j++ { 
            for k := 0; k < len(m); k++ { 
            } 
        } 
    }  
```

As you can see, you need as many for loops as the dimensions of the array in order to access all of its elements.