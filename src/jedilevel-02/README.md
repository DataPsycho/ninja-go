# Highlights

- CS Fundamentals
- Encoding of Character ASCii and UTF-8
- Neumeral, String and Other Types

## Type Importance

Go compiler does decide auto assignment of different specs of a type (int8, int64 for int) by itself. But for software development point of view its our duty to save each and every bytes of allocation carefully where the use of different version of types come to the action.

## uint{8, ..., 64} VS int{8, ..., 64}

As 1 single bit can have two switch on or off. 8 bits makes 1 bytes and the rest calculation goes according to 2^n. So in theory 8 bits can have 2^8 can store 256 seperate combination. As in Go counting start from 0 for example:

- A single uint8 can store 0 to 255 using its all possible combination
- A single int8 can store -128 to 127 the negative and positive integers

Using only int/uint the compilar will decide based on the machine architecture to choose between type variant.

The `byte` keyword can be used instead of `uint8` and `rune` can be used instead of `int32`. As each character is up to 4 bytes in UTF-8 and 1 byte is 8 bit so we need total 8*4, 32 bit to store any kind of character exist in the world.

## Storing Strings

String mades up with slices of bytes, Where slices is a composit data type. Also as the first part of UTF8 is ASCii so it uses the same numeber code for character we can decode a string as followes:

```go
s := "Hello, Playground"
bs := []byte(s)
```

Now the pringing will shows up a slice with numbers `[72 101 ... 100]`, where `72` represent the word `H`.

## iota a Predeclieared Identifier

iota is a special identifier can be used along with `const` reserved keyword can generate incremental value in a sequence. Probably has the simirar functinality like python enumerate function.

## Referances

- [Type specs](https://golang.org/ref/spec#Types); go doc
- [Types in Go](https://www.golang-book.com/books/intro/3); Golang Book
- [Strings, bytes, runes and characters in Go](https://blog.golang.org/strings); Blog Post
- [iota in Go](https://go101.org/article/keywords-and-identifiers.html)
