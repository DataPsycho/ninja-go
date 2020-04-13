# Highlights

- Channels

## Channel 

It can be consider a place where we can send data for further use.

in case of just create channel and send a value to the channel is not possible. Because give away and take away of the data to some channel should be happened at the same time.

```go
package main

import "fmt"

func main() {
    c := make(chan int)
    c <- 42

    fmt.Println(<-c)
}
```

Here, the main function create a channel and try to put 42 in that channel but the channel blocks because the send and receive is not happeing at the same time.

To fix that we can put that part of code into go routine with will create different async. call and send and reveive will happen at the same time. As the `c <-` and `<-c` is happening one after another but they should be happeing at the same time. Using a go routine we can arrange the give and take design pattern at the same time.

```go
package main

import "fmt"

func main() {
    c := make(chan int)

    go func() {
        c <- 42
    }()
    fmt.Println(<-c)
}
```

Another way is to use buffer channel, which will allow a value to sit in the channel and wait for take it out.

```go
package main

import "fmt"

func main() {
    c := make(chan int, 1)

    c <- 42

    fmt.Println(<-c)
}
```

But if we send more value than the waiting list of channel the buffer will fail.

```go
package main

import "fmt"

func main() {
    c := make(chan int, 1)

    c <- 42
    // This will block the channel
    c <- 43

    fmt.Println(<-c)
}
```

## Directional Channel

We can create channel which can only send values or only receive values to/from the channels. The above example were bidirectional channel and can do both.

## Range

Range is pulling values from a channel until the channel is closed.

```go
package main

import "fmt"

func main() {
    c := make(chan int)
    // send
    go func() {
        for i := 0; i < 100; i++ {
            c <- i
        }
        close(c)
    }()
    // receive

    for v := range c {
        fmt.Println(v)
    }
    fmt.Println("About to Exit!")
    }
```

## Select

Select statement can be used to pull values off from multiple channel in a infinite loop.

```go
package main

import "fmt"

func main() {
    eve := make(chan int)
    adam := make(chan int)
    moses := make(chan int)

    go send(eve, adam, moses)

    // receive
    receive(eve, adam, moses)

    fmt.Println("About To Exit!")
}

func send(e, a, m chan<- int) {
    for i := 0; i < 100; i++ {
        if i%2 == 0 {
            e <- i
        } else {
            a <- i
        }
    }
    close(e)
    close(a)
    m <- 0
}

func receive(e, a, m <-chan int) {
    for {
        select {
        case v := <-e:
            fmt.Println("From the eve Channal", v)
        case v := <-a:
            fmt.Println("From the adam Channal", v)
        case v := <-m:
            fmt.Println("From the moses Channal", v)
            return
        }
    }
}
```
