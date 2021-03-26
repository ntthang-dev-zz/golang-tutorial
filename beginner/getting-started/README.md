# Getting Started With Go

Go is an absolutely incredible language to build a wide variety of different applications in. From command-line interfaces to distributed microsystems and even cloud platforms, its' simplicity and concurrency baked in makes it a powerful choice of language for many development teams.

In this tutorial, I will be embarking on a mission to help get you up and running with the language so that you can go off and build ever-more brilliant applications and help push forward technology.

We’ll be focusing on getting a really simple `Hello World` style application up and running. Once you’ve got everything working correctly, we can start the rest of our journey learning the more complex aspects of the language such as functions, methods, and eventually things like concurrency and reflection.

## Prerequisites

Before you can follow this article, you will need the following:

- You will need Go installed on your development machine. If you need to install this you should check out the official download page:
[Official Go Download](https://golang.org/dl/)

## Getting Started

Let’s dive in with getting everything installed and writing an incredibly simple program to get our toes wet.

You’ll first want to go to the official [Getting Started page](https://golang.org/doc/install) which will contain a link to install the various different versions of Go currently available.

At this point, we’ll want to install the latest version, which, at the time of writing this is `go 1.11.4`.

Installing this using the `Installed` kind should add the official Go binary to your machine’s PATH. With this done, we should be able to run `go version` within a terminal:

```sh
$ go version
go version go1.11 darwin/amd64
```
If this works correctly, we are all set to start writing our own Go programs.

Let’s open up our code editor of choice and then create a new directory in which our `Hello World` project will reside.

Within this directory, we’ll create a new file called `main.go` which will contain our relatively simple Go program. We’ll also want to open up a terminal at this directory location and run the following commands:

```sh
$ GOMODULES11=ON
$ go mod init github.com/hello/world
```
This will initialize our project and allow us the ability to separate our Go code into sub-packages in the future. It’ll also allow us to retrieve any external dependencies we may want with minimal fuss.

Now, within our `main.go` file, we’ll want to add the following code:

```go
// the first statement of every go source file
// must be a package declaration. If we aren't doing anything
// fancy, this tends to be package main.
package main

// We then want to use the fmt package
// which features a `print` function - Println
import "fmt"

// We then need to define our main function.
// Think of this as the entry point to our Go
// program
func main() {
    // within this main function, we then
    // want to call a function within the fmt
    // package called Println() in order to print
    // out `Hello World`
    fmt.Println("Hello World")
}
```

And that’s all you need! Once you’ve added these 5 lines of code, we can set about running and compiling this using our `go` binary which is now on our `PATH`.

If we wanted to compile this into a binary executable, we can do so by again using the `go` binary like so:

```sh
$ go build main.go
$ ./main
```

## Conclusion

So, in this simple tutorial, we managed to successfully start our journey into the world of Go development.

If you are interested in further improving your Go skills, I recommend checking out the next article in this series on the basic types available in Go - Go Basic Types Tutorial
