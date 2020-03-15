# Highlits

- Control Flow
- Iterator, loops

## For Loop

For loop in go consist of init, condition and post argument, here init is initialization of if the loop,
conditon is the condition for the loop it stop and argument is the interval for iteration.

```go
for [init]; [condition]; [post] {
    if [condition] {
        ...
    }
}
```

```go
// here post is the intervel for each step
for i := 0; i <= 100; i++ {
    ...
}
```

Alternative with for as while, also called for with single condition:

```go
x := 1
for x < 10 {
    ...
    // incrementing x
    x++
}
```

## Continue and Break

Break will break the control flow in go and continue will skip the any next operation if the condition failed. Otherwise it will run the whole control flow.

```go
// brake statement
if [condition] {
    break
}

// Continue statement
if [condition] {
    continue
}

```

## If Statement

Simple true, false codition can be done with true, false or anti true, false combination.

```go
if [false/true/!true/!false] {
    [some statement]
}
```

Multiple condtion is also possible to impose with `;` operator.

```go
if x:=42, x == 42 {
    [some statement]
}
```

Else can also be combined with else if as follows:

```go
if [condition] {
    [some statement]
} else if [condition] {
    [some statement]
} else [condition] {
    [some statement]
}
```

## Switch-Case and Fallthrough

Based on the control flow of one of the statement is true the rest of the condition will not be evaluated in `switch`.

```go
switch {
    case [condition]:
        [some statement]
    case [condition]:
        [some condition]
}
```

To evaluate all statement in `switch` we can use `fallthrough` statement.

```go
switch {
    case [condition]:
        [some statement]
    case [true condition]:
        [some condition]
        fallthrough
    case [true condition]:
        [some condition]
}
```

If noting fallthrough the `default case` statement can be use to have a default return.

```go
switch {
    case [condition]:
        [some statement]
    case [true condition]:
        [some condition]
        fallthrough
    case [true condition]:
        [some condition]
        fallthrough
    default:
        [some statement]
}
```

It is also possible to add default value in `switch` and case will try to evaluate it. Switch also except multiple default value.

```go
switch [value] {
    case [a value]:
        [another value]
    default:
        [some statement]
}
```

## Referances

- [For Statement](https://golang.org/ref/spec#For_statements); Go Lang Spec For Statement
- [If statement](https://golang.org/ref/spec#If_statements); Go Lang Spec If statement
