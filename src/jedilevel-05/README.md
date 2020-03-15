# Highlits

- Struct composit data structure
- Operation/Manipulation of struct type

## Struct Building Block

```go
type person struct{
    first string
    last string
}
// instance of that struct
p1 := person{first: "Data", last: "Psycho"}
p2 := person{first: "Data", last: "Monger"}
```

It is possible to access each element of that oblect using `dot` notation.

```go
type person struct{
    first string
    last string
}
// instance of that struct
p1 := person{first: "Data", last: "Psycho"}
p2 := person{first: "Data", last: "Monger"}
fmt.Println(p1.first)
fmt.println(p2.first)
```

### Embeding one type to anoter type

We can take one defined type and extend another new struct using existint struct.

```go
type person struct{
    first string
    last string
}

type musician struct{
    person
    instrument string
}

// Creating an instance
ms1 = musician{
    // have to be defined seperately
    // Here person is considered as imported type
    person: person{first: "String", last: "Rider"},
    instrument: "Guitar"
}
```

With out having assigned the struct to a variable and create it in inside of function with decleraton will indtroduce a annonymous type.

```go
p1 = struct{
    first string
    last string
}{
    first: "Data",
    last: "Psycho",
}
```
