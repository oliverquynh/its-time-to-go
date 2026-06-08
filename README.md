# It's time to GO!

This is where I learn Go programming language.

## General

- Initialize a GO project.

```bash
cd /path/to/somewhere

mkdir its-time-to-go

cd its-time-to-go

go mod init its-time-to-go
```

- A GO program starts from a `main` function, inside a `main` package.

```go
// main.go or whatever.
package main

import "fmt"

func main() {
    fmt.Println("This is where a Go program starts from.")
}
```

> A directory allows only 1 file which declares "package main"

## Zero values

When declaring variables without assigning values, Go will assign default values based on their types. Those default values are zero values.

- integer: 0
- float: 0
- bool: false
- string: ""


## Data Types

- Go has 4 signed integer types: int8, int16, int32 and int64.
- Go has 4 unsigned integer types: uint8, uint16, uint32 and uint64.
- int and uint are platform-dependent types, meaning their size changes depending on the CPU architecture you compile for
- Go has 2 floating point types: float32 and float64.
- bool is similar to other languages which contains 2 values `true` and `false`.
