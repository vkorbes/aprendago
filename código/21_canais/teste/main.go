package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	q := make(chan bool)

	go boop(c, q)

	for {
		select {
		case v, ok := <-c:
			if ok {
				fmt.Println(v)
			}
		case <-q:
			return
		}
	}

}

func boop(c chan<- int, q chan<- bool) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
	time.Sleep(5 * time.Microsecond)
	q <- true
	close(q)
}
