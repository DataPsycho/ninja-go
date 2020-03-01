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
