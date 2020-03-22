# Highlights

- Functions, Returns Arguments, Types

## Signature of Function

In Go functions will have the following signature:

```go
func (r recever) identifier(parameter(s)) (return()) {Inner Code ...}
```

Here `func identifier() {}` are must have signature of function, other elements are optional. Here is an example:

```go
func foo() {
    ... do some thing ...
}
```

## Function with Varaidic Parameter

Providing function paranter with `...<datatype>` will invoce the variatic feature of a function which means It will be able to take any number of value of a specific type.

```go
func foo(x ...int){
    fmt.Println(x)
}
```

That function will accept any number of `integer` as parameter.
