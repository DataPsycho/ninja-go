# Highlights

- Error Handling
- Checking Error
- Printing and Logging

Remember, **Always Check Errors.**

## Writting To a File

We can apply error handling and defer when writing any operation. Here we are trying to create a file in the desktop while checking if there is any error and then using defer to close the file.

```go
package main

import (
    "fmt"
    "io"
    "os"
    "strings"
)

func main() {
    f, err := os.Create("names.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer f.Close()
    r := strings.NewReader("Data Psycho")
    io.Copy(f, r)
}
```

## Writing Log into Error

```go
package main

import (
    "log"
    "os"
)

func main() {
    f, err := os.Create("log.txt")
    if err != nil {

    }
    defer f.Close()
    log.SetOutput(f)

    f2, err := os.Open("no-file.txt")
    if err != nil {
        // fmt.Println("Error Happened", err)
        log.Println("Error Happened", err)
        // log.Fatalln(err)
        // panic(err)
    }
    defer f2.Close()
}
```

Using log fatal the defere function will not run. Panic will stop execution of current go routine running. Log package might be better choose in case of logging the error.

## Recover

Recover is only useful inside of defer function. Here we can see how to manage panic and over come panic with recover.

```go
package main

import "fmt"

func main() {
    f()
    fmt.Println("Returned normally from f.")

}

func f() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered in f", r)
        }
    }()
    fmt.Println("Calling g.")
    g(0)
    fmt.Println("Returned normally from g.")
}

func g(i int) {
    if i > 3 {
        fmt.Println("Panicking")
        panic(fmt.Sprintf("%v", i))
    }
    defer fmt.Println("Defer in g", i)
    fmt.Println("Panicking in g", i)
    g(i + 1)
}
```

## Adding Info into The Errors

We can create our own erro message using `New` method which take the advantage or Error interface of go.

```go
package main

import (
    "errors"
    "fmt"
    "log"
)

func main() {
    f()
    fmt.Println("Returned normally from f.")

}

func f() {
    _, err := sqrt(-10)
    if err != nil {
        log.Fatalln(err)
    }
}

func sqrt(f float64) (float64, error) {
    if f < 0 {
        return 0, errors.New("Stat C++ lab: square root of negative number")
    }
    return 42, nil
}
```
