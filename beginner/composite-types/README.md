# Go Composite Types Tutorial

Welcome All! In this tutorial, we are going to be looking at the various different composite data types that are available in the Go programming language.

If you haven’t already, I’d suggest you check out my other tutorial in this course on [Basic Data Types](https://github.com/ntthang-dev/golang-tutorial/tree/main/beginner/basic-types). You’ll need to know about these basic data types in order to understand some of the composite data types.

## Arrays

Let’s dive into our first Composite data type, the array and see how we can declare `arrays` and work with them.

Let’s start by declaring an array of all of the days of the week. This will be an empty array for now:

```go
// declaring an empty array of strings
var days []string

// declaring an array with elements
days := [...]string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}
```

If we want to query the first element in the array, or a specific element, we can do so in a very similar fashion to other languages:

```go
fmt.Println(days[0]) // prints 'monday'
fmt.Println(days[5]) // prints 'saturday'
```

### Slices

The difference between slices and arrays in very subtle and it’s something that has most definitely caught me out a few times in the past. Slices in Go allow you to access a subset of an underlying array’s elements.

Slices are made up of three things, a `pointer`, a `length`, and a `capacity`. Let’s try and visualize this with an example. Say, for instance, we have an array of the days of the week, we could use slices to extract only those days that are week days.

```go
days := [...]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
weekdays := days[0:5]
fmt.Println(weekdays)
// This returns: [Monday Tuesday Wednesday Thursday Friday]
```

### Maps

Maps are Go’s representation of hash tables, a data structure that allows you to map one arbitrary data type to another. For instance, let’s create a map of YouTube channel names to the numbers of subscribers for that channel:

```go
youtubeSubscribers := map[string]int{
  "TutorialEdge":     2240,
  "MKBHD":            6580350,
  "Fun Fun Function": 171220,
}

fmt.Println(youtubeSubscribers["MKBHD"]) // prints out 6580350
```

This represents a mapping between a `string` data type and an `int` data type.

### Structs

In Go, we have this concept of a `struct`. These `structs` allow us to create data types that are aggregates of other data types.

Say for instance, we had this concept of a `Person` within our application. We could create a person `struct`, that has a number of fields within it, we could for instance, have a `name` field which is of type `string` and an `age` field which is of type `int` like so:

```go
// our Person struct
type Person struct {
  name string
  age int
}

// declaring a new `Person`
var myPerson Person
```

The advantage of using these `struct` is that we can effectively treat all of these values or `fields` as they are called as a single entity and modify that easily.

```go
// declaring a new `elliot`
elliot := Person{name: "Elliot", age: 24}

// trying to roll back time to before I was injury prone
elliot.age = 18
```

### Nested Structs

Structs are incredibly extensible due to the fact we can create nested structs within our structs. For example, imagine we had a `Team` struct who had a number of people within that team:

```go
package main

import (
    "fmt"
)

func main() {
    type Person struct {
        name string
        age  int
    }

    // our Team struct
    type Team struct {
        name    string
        players [2]Person
    }

    // declaring an empty 'Team'
    var myTeam Team
    fmt.Println(myTeam)

    players := [...]Person{Person{name: "Forrest"}, Person{name: "Gordon"}}
    // declaring a team with players
    celtic := Team{name: "Celtic FC", players: players}
  fmt.Println(celtic)

}
```

## Conclusion

So, hopefully you found this tutorial useful and it has given you insight as to how you can use more advanced data types within your own Go programs.