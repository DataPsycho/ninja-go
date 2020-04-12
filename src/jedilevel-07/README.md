# Highlights

- Pointers(location) of a value of a particular type
- `&` vs `*` operator
- Use of Pointers
- Method Sets for a type

Pointer is type like any other premitive or composite types in Go.

## `&` -  Address in Memory

The memory address of any value `x` of type let say `typeA` can be seen using `&x` in go, which should output some arbitary value such as `0x10...`.

## `*` - See the Value of a Pointer Type

If `x` is a value of type `pointer` let say `typeA` the we can access the value stored in that pointer location with using `&x`.

## Change the value in the address

It is possible to rewrite the value of the pointer is pointing to. But the main and reference variable both will be affected by the change.

```go
a := 12
// createing a pointer
b:= &a
// changing the value of the pointer
// a will aslo be assigned to 43
*b = 43
```

## Use of Pointers

Usually passing a large amount of data (values of some particular data type) in to the function is memory consuming as the data is copied in to the function and do the operation and return a new result and it does not touch the main data.

But we can pass the location of the data in the memory in the function and do all the expensive operation, that will not copy the data again but it will peform the operation on the existing locations of the data in memory, but in that case the change will rewrite the main data.

```go
func main() {
    a_int := 32
    fmt.Println("Before passing pointer I am:", a_int)
    opt_on_pointer(&a_int)
    fmt.Println("After passing pointer I am:", a_int)
}

func opt_on_pointer(y *int) {
    constant := 123
    *y = constant
}
```

Link in [Playground](https://play.golang.org/p/UB-Uo89Zbv9).

## Method sets

When we have a method with recever of type pointer, the method can operate only on the pointer of the value.

But if the recever is general type it can operate on the value or either on the pointer of the value.
