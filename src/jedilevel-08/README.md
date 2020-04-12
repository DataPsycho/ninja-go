# Highlingts

- Reading Standard Library
- Working with Standard Library
- Application
- Json marshal and un-marshal

## Package Documentaiton

The package code and documentation can be found as follows:
`golang.org>documentation>Package Documentation>package Name` or easyly can be found in godoc.org. For example we can find the `json` package and docuemntaion in godoc.org/encoding/json or we can search by the name. The `index`, `Example` and `File` link will provide more insights on those libraries.

## Working With Json Marshal and Unmarshal

Jason Marshal method take a unlimited value as argument and return a `byte` like object with successful execution and a error status. Later we can use `write` method from `os` to convert the `byte` in to json view. **For unmarshal the argument value have to be a pointer**.

So, __Main take aways is, Considering json data as protocol of data transfer we will be able to convert string in to slice of byte or slice of byte to string using Marshal and Unmarshal method.__

```go

type person struct {
    first string
    last string
    age int
}

p1 := person{
    First: "Data",
    Last: "Psycho",
    Age: 29,
}

p2 := person{
    First: "Data",
    Last: "Nerd",
    Age: 30,
}

// slice or persons
people := []person{p1, p2}

// marshalling person
bs, err := json.Marshal(people)
if err != nil {
    fmt.Prinln(err)
}
fmt.Println(string(bs))
```

When we get a json string we need to convert the string in to byte which can be redable ty go Unmarshal function. **But also we need to define a data structure which will be used by unmarshal method when serialize the byte to go native data structure, in that case we also need to provide tags which will map each value or a struct to its identifier by `json:Identifier`** :

```go
type person struct {
    Name string `json:Name`
}

s := `[{"Name": "Data Psycho"}, {"Name": "Data Nerd"}]`
bs := []bytes(s)

var people []person

err := json.Unmarshal(bs, &people)

if err != nil {
    fmt.Prinln(err)
}
fmt.Println("All the People", people)
```

## Json Encode and Decode

In case of serialize the data directly in to program and not using any marshal or unmarshal we can use `encode` and `decode` methods. So for decode will equivalant to unmarshal.

## Writer Interface

`Writer` is an interface to write serialized values in to some file or in the program for furtuer manipulation. 

```go
type Writer interface {
    Writer(p []byte) (n int, err error)
}
```

So any UDT (User Defined Type) will have that method will be able to peform `io` methods of writer.

## Sorting Slice

Go built in `sort` package provide sorting functionality on any primitive or UDT (User Defined Type).

```go
xi := []int{4, 1, 12, 11}
// sort package need to be imported
sort.Ints(xi)
```

## Sorting Values of Composit Data

lets say we have a composite thpe `typeA` and we have a slice  `x` or values of `typeA` and we want to sort the slice by any attribute of the `typeA`. We should create a new type which will have the same type of `typeA` the if the new type will have `len swap & less` attached we can use `sort.Sort()` to sort the values of slice `x` after converting the type of `x` to new type as follows:

```go
package main

import (
    "fmt"
    "sort"
)

func main() {
    people := []person{
        {"Data Psycho", 29},
        {"Data Nerd", 32},
        {"String Rider", 30},
    }
    fmt.Println(people)
    // Sorting after conversion
    sort.Sort(ByAge(people))
    fmt.Println(people)
}

// A type
type person struct {
    name string
    age  int
}

// Funciton attacked to person
// for spacial string print [name: age name: age ...]
func (p person) String() string {
    return fmt.Sprintf("%s: %d", p.name, p.age)
}

// A new type of slice of person type
type ByAge []person

// implementing the 3 methods to Byage type
func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].age < a[j].age }
```

Playground [link](https://play.golang.org/p/T5z6y7N0fRn).

## Bcrypt

```go
s := "password123"
bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
if err != nil {
    fmt.Println(err)
}
fmt.Println(bs)
```