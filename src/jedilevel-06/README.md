# Highlights

- Functions, Returns Arguments, Types
- Interfaces, Polymorfhism
- Annonymous Functions
- Callback
- Clouser

## Signature of Function

In Go functions will have the following signature:

```go
func (r recever) identifier(parameter(s)) (return(s)) {
    Inner Code ...
}
```

Here `func identifier(parameters) {}` are must have signature of function, other elements are optional. Here is an example:

```go
func foo() {
    ... do some thing ...
}
```

## Function with Varaidic Parameter

Providing function parameter with `...<datatype>` will invoce the variatic feature of a function which means It will be able to take any number of value of a specific type.

```go
func foo(x ...int){
    fmt.Println(x)
}
```

That function will accept any number of `integer` as parameter. Now lets add returns into a sum function.

```go
// return type will be an integer
func sum(x ...int) int {
    sum := 0
    // loop over the input array
    for _, v := range x {
        sum += v
    }
    return sum
}
```

But we must remember, **The variadic parameter must have to be the final parameter in the function**

## Passing Slice to Function

In the above function, the indentifier `x` is expecting integers but it is also possible to pass a slice of integers using `...` variadic operator. Which we can call unfurling an slice to use it as parameter.

```go
func main() {
    slice = []int{1, 3, 4, 5}
    // using slice... instead of raw
    // integer value the, sum function is above
    x = sum(slice...)
}
```

In case of having variadic parameter if we pass nothing it will return zero as sum as the initial zero for for integers is zero.

```go
func main() {
    slice = []int{}
    // using slice... instead of raw
    // integer value the, sum function is above
    x = sum(slice...)
}
```

## `defer` statement

If we need to run some memory intensive operation and close it just before next function call `defer` should be used top of that function.

```go
func main() {
    defer foo()
    bar()
}

func foo() {
    ...
}

func bar() {
    ...
}
```

In that situation `bar` will run first will close any execution inside of `bar` and run `foo`.

## Writing Methods(Utilizing the Recever of Function)

Having a recever in the function will narrow down the type exceptence. Usually some type is going to be attached to a function. Which make the function accessible by any value of the same type. We can think is as methods of a `class` in OOP the way the `@classmethod` has access to its initializers. But here in go the methods is assciated with that particular instance of value of a type.

```go
type TypeA struct {
    ...
    ...
}

type TypeB struct {
    TypeA
    ...
    ...
}

func main() {
    b1 := TypeB{
        TypeA : TypeA{
            ...
            ...
        }
        ...
        ...
    }
    // execution of underlying function
    // for falue of TypeB
    b1.method_for_b()
}

// Defining a Method for TypeB

func (t TypeB) method_for_b() {
    ...
    ...
}
```

## Interfaces and Polymorphism

A value of a particular type can be assigned to more that one type, so a value can be more than one type. So interface is allow us to define behaviour.

Remember, **A value can be more that one type.**

So interfcaces can preserve some method signatures for other types with the same name and same output (Ex. area of a rectengle, area of circle has seperate way to calculate but at the end its a type of area). Based on the type passed to interface, interface will be able to select the right version of the method to calculate the output.

```go
// keyword-Identifier-type: type-TypeA-struct
type TypeA struct {
    ...
    ...
}

type TypeB struct {
    TypeA
    ...
    ...
}

// simple method for type A
func (t TypeA) simple_method() {
    ...
    ...
}

// Same simple method for type B
func (t TypeB) simple_method() {

}

// So any method has simple_method
// also will be considered as the type of InterfaceC
type InterfaceC interface {
    simple_method()
}

// defining a new method for InterfaceC
// now TypeA and TypeB will also have access to 
// method_for_c
func metnod_for_c(c InterfaceC) {
    ...
    ...
}

func main() {
    a1 := TypeA {
        ...
        ...
    }
    b1 := TypeB {
        TypeA : TypeA {
            ...
            ...
        }
        ...
        ...
    }
    // as b1 belongs to both types
    // which is a example of polymophism
    a1.simple_method()
    b1.simple_method()
    method_for_c(a1)
    method_for_c(b1)
}

// Now as the method simple_method also attached to
// interfaceC the value is also belongs to type of
// InterfaceC
```

## Anonymous Function

Annonymous function can be used when we do not need to indetify a function and also if the function is used for very small life cycle.

```go
func main() {
    // function without parameter
    func() {
        ...
        ...
    }() // here is the peren. to call the function

    // funciton with argument can except a integer
    func(x int) {
        ...
        ...
    }(2)
}
```

## Inline Function

We can create lambda like function using func expression in go.

```go
func main() {
    f := func(x int) {...}
    f(2)
}
```

## Pass a Function to Another Function

As function is considered as type in Go, we can asign function into a variable.

```go
func main() {
    // store the return from abc in to
    // a variable called value
    value := abc()
    fmt.Println(value)
}

func abc() string {
    ...
    ...
    return sometning
}
```

Function also can return a function.

```go
func main() {
    f := abc()
    fmt.Println(f())
}

func abc() func() int {
    // an arbitary function
    return func() int {
        return 4566
    }
}
```

## Callback (Taking Function as Argument)

Functions are considered as first class citizen in Go. So we can pass function as argument to another function. The callback general syntax will be `func name(f func() return_type, another_parameter type) return_type`. Let's see how to sum only even numbers using call back.

```go
func main() {
    ii := []int{1, 2, 3, 4, 5,}
    s := sum(ii...)
    even_sum = even(ii)
}

// sum will be used as call back function to sum
// even numbers in the next step
func sum(xi ...int) int {
    total := 0
    for _, v := range xi {
        total += v
    return total
    }
}
// even is taking a function as parameter with identifier
// f and a variadic parameter of type int with identifier
// vi
func even(f func(xi ...int) int, vi ...int) int {
    var yi []int
    for _, v := range vi {
        if v % 2 == 0 {
            yi = append(yi, v)
        }
    }
    return f(yi...)
}
```

## Closure

Use to create a small code block inside of functions of other code blocks to narrow down the scope of any particular operation.

```go
// The return value of a_func is a function with type int
// limiting the scope of variable some_variable with clouser
func a_func () func() int {
    var some_variable int
    return func() int {
        ... operation with some_variable ...
    }

}
```

Referances:

1. [Interfaces - You Tube](https://www.youtube.com/watch?v=EGRXKV6j-v0)
