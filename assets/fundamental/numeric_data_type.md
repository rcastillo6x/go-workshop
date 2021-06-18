# Numeric Data Type | [Home](../../README.md)

## Integers

Go offers support for four different sizes of signed and unsigned integers, named int8, int16, int32, int64; and uint8, uint16, uint32, and uint64, respectively. The number at the end of each type shows the number of bits used for representing each type.

Additionally, int and uint exist and are the most efficient signed and unsigned integers for your current platform. Therefore, when in doubt, use int and uint, but have in mind that the size of these types changes depending on the architecture.

The difference between signed and unsigned integers is the following: if an integer has eight bits and no sign, then its values can be from binary 00000000 (0) to binary 11111111 (255). If it has a sign, then its values can be from -127 to 127. This means that you get to have seven binary digits to store your number because the eighth bit is used for keeping the sign of the integer. The same rule applies to the other sizes of unsigned integers

## Floating-point numbers

Go supports only two types of floating-point numbers: float32 and float64. The first one provides about six decimal digits of precision, whereas the second one gives you 15 digits of precision.

## Complex numbers

Similar to floating-point numbers, Go offers two complex number types named complex64 and complex128. The first one uses two float32: one for the real part and the other for the imaginary part of the complex number, whereas complex128 uses two float64. Complex numbers are expressed in the form of a + bi, where a and b are real numbers, and i is a solution of the equation x2 = −1.

All these numeric data types are illustrated in the code of numbers.go, which will be presented in three parts