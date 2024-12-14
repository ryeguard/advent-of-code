# goac

Go utilities for Advent of Code:

- Utility/helper functions in the package `goac`.
- A CLI in [`cmd`](./cmd/)

## Utils

Use the utility functions in your Go code like so:

```go
//main.go
package main

import (
    "fmt"

    "github.com/ryeguard/advent-of-code/goac"
)

func main() {
    input, err := goac.ReadInput("./path/to/file")
    if err != nil {
        panic(err)
    }

    fmt.Println(input)
}
```

## CLI

### Usage

To build and run:

```bash
# from the cmd dir (/goac/cmd)
go run . help

# from the project root
go run ./goac/cmd/main.go help
```
