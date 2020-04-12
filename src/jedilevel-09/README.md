# Highlights

- Concurrency in Brief
- Waitgroup
- Mutex, Atomic

## Concurrency

Concurrency is a design pattern in Go which runs part of a operation of a program  in parallel. Using `go` in front of the function execution line will create a seperate go routine implice a asyncronous call and next part of the code will not wait for that go routine to complete. Concurrency is not a gurantee but a week promise that, if possible the goroutine will run in parallel. One of the condition to satisfy for go routine is system have to have multiple CPUs.

## WaitGroup - Return from Goroutine

To get the result our from the concurrent programs we can apply WaitGroup interface from base package `sync`. We have to use the following methods to archicts the concurrency proberly: Add(), Wait(), Done()`.

```go
package main

import (
    "fmt"
    "runtime"
)

var wg sync.WaitGroup

func main() {
    fmt.Println("Goroutines:", runtime.NumGoroutine())
    fmt.Println("===========The Funcs ===========")
    wg.Add(1)
    go a_func()
    b_func()

    fmt.Println("===========The Routines ===========")
    fmt.Println("Goroutines:", runtime.NumGoroutine())
    wg.Wait()

}

func a_func() {
    for i := 1; i <= 10; i++ {
        fmt.Println("Talking from A:", i)
    wg.Done()
    }
}

func b_func() {
    for i := 1; i <= 10; i++ {
        fmt.Println("Talking from B:", i)
    }
}

```

`Add` will signal a that many number of async. call we are expecting at the end. `Done` will signal after the successful execution.`Wait` will wait untill all the routines added by `Add` are `Done` with execution.

## Race Conditions

Sharing same variable in by multiple goroutine is a not a safe design architecture. As any point in time any go routine can chanage the state of the variable without preserving the previous state of the variable and can lead to a bad design.

```go
package main

import (
    "fmt"
    "runtime"
    "sync"
)

var wg sync.WaitGroup

func main() {
    fmt.Println("CPUs:", runtime.NumCPU())
    fmt.Println("GOroutines:", runtime.NumGoroutine())

    counter := 0
    // Setting up waitGroup
    const gs = 100
    var wg sync.WaitGroup
    wg.Add(gs)

    for i := 0; i <= gs; i++ {
        go func() {
            v := counter
            runtime.Gosched()
            v++
            counter = v
            wg.Done()
        }()
        fmt.Println("GOroutines:", runtime.NumGoroutine())
    }

    wg.Wait()
    fmt.Println("Coun:", counter)
}
```

The above program has problem with writing the goroutines on the same location. The final counter result will not be hundread what we expect but will provide very incosistent result. `go run -race <filename>.go` will raise the condition.

## Handling Raise Condition - Mutex

Mutex put a lock condition on a certain variable untill some process is using the variable.

```go
package main

import (
    "fmt"
    "runtime"
    "sync"
)

var wg sync.WaitGroup

func main() {
    fmt.Println("CPUs:", runtime.NumCPU())
    fmt.Println("GOroutines:", runtime.NumGoroutine())

    counter := 0
    // Setting up waitGroup
    const gs = 100
    var wg sync.WaitGroup
    wg.Add(gs)

    // Mutex to Lock a Task
    var mu sync.Mutex  

    for i := 0; i <= gs; i++ {
        go func() {
            // Locking Counter
            mu.Lock()
            v := counter
            runtime.Gosched()
            v++
            counter = v
            // Unlocking Counter
            mu.Unlock()
            wg.Done()
        }()
        fmt.Println("GOroutines:", runtime.NumGoroutine())
    }

    wg.Wait()
    fmt.Println("Coun:", counter)
}
```

But a better way to handle such operation is Channels.

## Async Call with Atomic

The package atomic has a special counting style using integers or units type in a async. fashion. The passing argument must have to be a pointer to a decleared variable which it will increment.

```go
package main

import (
    "fmt"
    "runtime"
    "sync"
    "sync/atomic"
)

var wg sync.WaitGroup

func main() {
    fmt.Println("CPUs:", runtime.NumCPU())
    fmt.Println("GOroutines:", runtime.NumGoroutine())

    var counter int64
    // Setting up waitGroup
    const gs = 100
    var wg sync.WaitGroup
    wg.Add(gs)

    for i := 0; i < gs; i++ {
        go func() {
            atomic.AddInt64(&counter, 1)
            runtime.Gosched()
            fmt.Println("Counter\t", atomic.LoadInt64(&counter))
            wg.Done()
        }()
        fmt.Println("GOroutines:", runtime.NumGoroutine())
    }

    wg.Wait()
    fmt.Println("GOroutines:", runtime.NumGoroutine())
    fmt.Println("Coun:", counter)
}
```
