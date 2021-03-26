# Go Basic Types Tutorial

In this tutorial, we are going to be looking at all of the basic data types available to us within the Go language. By the end of this tutorial, you should be comfortable with the various different types available within the language and hopefully some understanding as to how you can use these within your own Go programs.

This kind of material can be fairly dry and boring to learn so I’ll try and spice things up and make it somewhat interesting whilst also covering the necessary basics.

## Data Types

So, to get us started, it’s important to know that there are 4 distinct categories of types within the Go programming language:

- **Basic Types** - what we’ll be covering in this tutorial
- **Aggregate Types** - These are arrays and structs
- **Reference Types** - These are your pointers and slices
- **Interface Types** - These are your standard interfaces

## Integers

The first basic type we will cover is the `Integer` type.

We can either use signed or unsigned integers within our program and we can specify the size of integer we need. Why do we want to specify the size you may ask, well, imagine you were trying to optimize your program’s memory utilization. If you know a certain number isn’t going to exceed a specific value, you can then select a size that is suitable for that value:

We can create new integer variables by typing either `uint` or `int` with the size of the int appended to the end. If we wanted an 8 bit unsigned integer then it would look like `var myint uint8` and so on:

```go
// all numeric types default to 0

// unsigned int with 8 bits
// Can store: 0 to 255
var myint uint8
// signed int with 8 bits
// Can store: -127 to 127
var myint int8

// unsigned int with 16 bits
var myint uint16
// signed int with 16 bits
var myint int16

// unsigned int with 32 bits
var myint uint32
// signed int with 32 bits
var myint int32

// unsigned int with 64 bits
var myint uint64
// signed int with 64 bits
var myint int64
```

It should be noted that if you try and assign a larger value to an int than it can handle, i.e:

```go
var myint int8
myint = 2500
```

The Go compiler will fail to run or build the program and will output the fact that 2500 overflows `int8`. However, if you were to overflow the integer at runtime, then you may start seeing odd results. For example, try running this program and examine the output:

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println("Hello World")

    var myint int8
    for i := 0; i < 129; i++ {
        myint += 1
    }
    fmt.Println(myint) // prints out -127
}
```

This is due to the fact that the result of this particular operation has caused the signed integer to overflow. This is something to watch out for within your programs!

## Standard int Type

If all of this verbosity in defining your integer values is too much, then you should know that for most of the time, you can default to just int. This int data type is typically either 32 bits in size or 64 bits depending on whether your underlying system is a 32-bit system or a 64-bit system.

For simplicity, it’s best to default to this data type and you’ll see this being the most widely used in the wild.

### Conversion of Types

When it comes to working with multiple variables with different data types, you will more often than not have to cast your various integer variables to `int`. This will handle conversion from things like `uint8` and `int16` to a standard 32, or 64 bit signed int and from there you will be able to do things like addition, multiplication and subtraction.

```go
var men uint8
men = 5
var women int16
women = 6

var people int
// this throws a compile error
people = men + women
// this handles converting to a standard format
// and is legal within your go programs
people = int(men) + int(women)
```

## Floating Point Numbers

Next, we come to Floating Point numbers. These come in 2 distinct sizes, either `float32` or `float64` and allow you to work with exceptionally large numbers that don’t fit inside a standard `int64` data type.

```go
var f1 float32
var f2 float64
```

Let’s have a look at how you would declare and work with floats now:

```go
var maxFloat32 float32
maxFloat32 = 16777216
fmt.Println(maxFloat32 == maxFloat32+10) // you would typically expect this to return false
// it returns true
fmt.Println(maxFloat32+10) // 16777216
fmt.Println(maxFloat32+2000000) // 16777216
```

### Converting float to int and back again

If you want to convert integers to floats or floats to int then you can achieve that by casting the variable as your desired data type.

```go
// converting from int to float
var myint int
myfloat := float64(myint)

// converting from float to int
var myfloat2 float64
myint2 := int(myfloat2)
```

## Complex Numbers

Ok, so we’ve covered both integers and floating points, but there is another commonly overlooked numeric data type and that is the complex number data type. These, much like the floating point data type, come in 2 distinct sizes, you can either go for `complex64` or `complex128`.

## Booleans

Now that we’ve covered all the basic numeric data types, we can move on to the other basic data types available in Go. The first of which, is the bool, or boolean data type.

A bool, represents either `true` or `false`. Let’s see how this can be used within one of our Go programs:

```go
var amazing bool
amazing = true
if amazing {
  subscribeToChannel()
}
```

Nice and simple, but what happens if we want to do a bit of boolean logic within our programs? Well, using the `||` and `&&` operators we can!

```go
var isTrue bool = true
var isFalse bool = false
// AND
if isTrue && isFalse {
  fmt.Println("Both Conditions need to be True")
}
// OR - not exclusive
if isTrue || isFalse {
  fmt.Println("Only one condition needs to be True")
}
```

## Strings

Strings within the Go language are what we would call character slices. We can declare a new string variable using `string`:

```go
var myName string
myName = "Elliot Forbes"
```

## Constants

Constants are our final basic data type within the Go language. They allow us to specify immutable values that will not change throughout the course of our programs execution.

```go
const meaningOfLife = 42
```

## Conclusion

So, we’ve managed to cover a lot in this tutorial and I’m hoping you enjoyed it! If you need any further help or would like to know more then please let me know in the comments section below!