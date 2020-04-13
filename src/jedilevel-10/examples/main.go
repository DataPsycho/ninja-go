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
