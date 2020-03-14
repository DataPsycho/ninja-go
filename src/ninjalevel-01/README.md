# Ninja Level 01

- Package Mangement

## Packages

Package information can be find in golang.org/doc. After finding the
package we can go to godoc.org with package name will show all api
available for that package.

## Short Declearation Operatior (:=)

This allows us to declare and assign a value to a variable. We need to be careful about keyword, operator, special signs. which can not be used as variable name. For example:

```go
x := 42
```  

After assignment it is possible to re-assign a new value:

```go
x = 99
```

## Declaring Variable with `var`

Using var keyword will also create a new variable:

```go
var y = 90
```

It is possible to initialize a variable with no value. In case it will take default assign value from the compiler:

```go
// will initialize with i := 0
var i int
```

Tips: Always search for spec in go doc org for language specification.

## Difference Between var and :=

`var` can be used as golobal level (out side of funcion) but `:=` can be only used inside of functional definitions.

## String Literals in GO

Raw string literal, which is tripple qoute in python can be written by using back coute.

```go
// The inner double quote will be preserved
var a string = `James saidn "Hi !"`
```

## Types in Go

- Primitive Data Type: Primitive datatypes are basic built in data types.
- Composite Data Type: Composit data types are build on top of primitive data types like array, struct etc.
- User Defined Data type: Custom types are are user declared types using primitive data types (yes you can create your own type)

## Exercises

### Exercise 1

Short Decleration Variable and Printing:

```go
package main

import "fmt"

func main() {
    x := 42
    y := "James Bond"
    z := true
    // multiple print
    // assigning single value will print single value
    fmt.Println(x, y, z)
}

output:
42 James Bond true
```

### Exercise 2

Print zero or default data for a variable with var variable assignment keyword.

```go
package main

import "fmt"

func main() {
    var x int
    var y string
    var z bool
    // for string it will print a empty space between 0 and false
    fmt.Println(x, y, z)
}

output:
0  false
```

### Exercise 3

Format and print with Sprintf.

```go
package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
    // For sprint you need location
    // identifier
    // Sprint + Format: Sprintf
    s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
    fmt.Println(s)
}

output:
42      James Bond      tru
```

### Exercise 4

User defined types and pringing data types.

```go
package main

import "fmt"

type mytype int

var myvar mytype

func main() {
    fmt.Println(myvar)
    // %T for printing the type
    fmt.Printf("%T\n", myvar)
    myvar = 5
    fmt.Println(myvar)
}

output:
0
main.mytype
5
```

### Exercise 5

Converting from one data type to another.

```go
package main

import "fmt"

type mytype int

var x mytype
var y int

func main() {
    fmt.Println(x)
    // %T for printing the type
    fmt.Printf("%T\n", x)
    x = 42
    fmt.Println(x)
    // Here is the conversion
    y = int(x)
    fmt.Printf("%T", y)
    fmt.Println("\n")
}

output:
0
main.mytype
42
int
```
