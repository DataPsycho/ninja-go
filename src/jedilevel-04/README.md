# Highlights

- Composite data types
- Data Structures and use of them

## Array

Sequence of elements of single types can be store together in an array.

```go
var myArray [5]int
```

After assigning an array it is possilbe to use several builtin function to manipulate the arrays such as `len(myArray)` for length of an array. Here 5 is the length and int is datatype of the `myArray`.

## Slices

Compose values of different data types together.

```go
x := []int{4, 5, 7, 8 12}
```

Here, `[]int` indicates a slice with integer datatype and `{}` indicates the values inside os a slice.

Looping over the slice:

```go
x := []int{4, 5, 7, 8 12}
// Using Range we have direct access to the values
for index, value := range x {
    fmt.Println(index, value)
}
```

Equivalant python code is:

```python
for index, value in enumerate(x):
    print(index, value)
```

Slicing a slice we can use `:` operator and same like python.

```go
x := []int{4, 5, 7, 8 12}
x[1:]
```

Code will return everything from 1st position.

It is also possible to make a larger array with capacity but with some initial values.

```go
x := make([]int, 10, 100)
```

Here `make` is making an arary with 10 initial values but there but it is possible to replace those zero values and we are able to append more 90 values.

Slice of slice or multidimentional slice is also possible.

```go
child1 := []string{"X", "Y", "Z"}
child2 := []string{"a", "b", "c"}
// [][]string refers it has an outer slice
// each inner slice is slice of stirng
parent := [][]string{child1, child2}
```

### Appending Slice

Slice can be expandable but arrays are not.

```go
x := []int{4, 5, 7, 8 12}
x = append(x, 1, 2, 3)
y = []int{98, 45}
x = append(x, y...)
```

As `append` is a function so we have to catch the outcomes into a new or existing save typed variable. Here `...` unpack all the values from the variable `y` and append to x.

### Delete from Slice

One way is to break a slice in index position and re append them together.

```go
x := []int{4, 5, 7, 8 12}
x = append(x[:1], x[4:]...)
```

### Maps

Maps are key value pairs like pyton dictionary.

```go
m := map[string]int{
    "Person1": 23,
    "Person2": 33,
}
```

Calling a map key wihich does not exist will return zero, but also the underlying function return 2 resulst 0 and OK as true or false. So we can catch the absense.

```go
m := map[string]int{
    "Person1": 23,
    "Person2": 33,
}

// ok variable will return false
v, ok := m["Person3"]
```

But if we want to do other operation using that statement we are able too in one statement:

```go
m := map[string]int{
    "Person1": 23,
    "Person2": 33,
}

// ok variable will return true
if v, ok := m["Person1"]; ok {
    [some operation with v]
}

```

Adding element to a map:

```go
m := map[string]int{
    "Person1": 23,
    "Person2": 33,
}

m["Person3"] = 34
```

Unpacking keys and values:

```go
m := map[string]int{
    "Person1": 23,
    "Person2": 33,
}

for k, v := range m {
    [some operation with k and value]
}
```

We can also delete the value from a map using key.

```go
m := map[string]int{
    "Person1": 23,
    "Person2": 33,
    "Person3": 34,
}

if v, ok := m["Person3"]; ok {
    delete(m, "Person3")
}
```

## References

- Go lang Docmentation
