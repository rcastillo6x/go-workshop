# Loops | [Home](../../README.md)

## Go loops

Every programming language has a way of looping and Go is no exception. Go offers the for loop, which allows you to iterate over many kinds of data types.

 
### The for loop

The for loop allows you to iterate a predefined number of times, for as long as a condition is valid, or according to a value that is calculated at the beginning of the for loop. Such values include the size of a slice or an array, or the number of keys on a map. This means that the most common way of accessing all the elements of an array, a slice, or a map is the for loop.

The simplest form of a for loop follows. A given variable takes a range of predefined values:

```
for i := 0; i < 100; i++ { 
} 
```

Generally speaking, a for loop has three sections: the first one is called the initialization, the second one is called the condition, and the last one is called the afterthought. All sections are optional.

In the previous loop, the values that i will take are from 0 to 99. As soon as i reaches 100, the execution of the for loop will stop. In this case, i is a local and temporary variable, which means that after the termination of the for loop, i will be garbage collected at some point and disappear. However, if i was defined outside the for loop, it will keep its value after the termination of the for loop. In this case, the value of i after the termination of the for loop will be 100 as this was the last value of i in this particular program at this particular point.

You can completely exit a for loop using the break keyword. The break keyword also allows you to create a for loop without an exit condition, such as i < 100 used in the preceding example, because the exit condition can be included in the code block of the for loop. You are also allowed to have multiple exit conditions in a for loop. Additionally, you can skip a single iteration of a for loop using the continue keyword.

### The while loop

Go does not offer the while keyword for writing while loops but allows you to use a for loop instead of a while loop. This section will present two examples where a for loop does the job of a while loop.

Firstly, let's look at a typical case where you might want to write something like while(true):

```
for { 
} 
```

It is the job of the developer to use the break keyword to exit this for loop.

However, the for loop can also emulate a do...while loop, which can be found in other programming languages.

As an example, the following Go code is equivalent to a do...while(anExpression) loop:

```
for ok := true; ok; ok = anExpression { 
} 
```

As soon as the ok variable has the false value, the for loop will terminate.

There is also a for condition {} loop in Go where you specify the condition and the for loop is executed for as long as the condition is true.

### The range keyword

Go also offers the range keyword, which is used in for loops and allows you to write easy-to-understand code for iterating over supported Go data types including Go channels. The main advantage of the range keyword is that you do not need to know the cardinality of a slice, a map, or a channel in order to process its elements one by one. 