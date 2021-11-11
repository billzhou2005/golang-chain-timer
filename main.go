package main

import (
	"fmt"
	"time"
)

func chanProcess(delay time.Duration, ch1 chan int, ch2 chan int) {

	t := time.NewTimer(delay)

	for {
		select {
		case <-ch1:
			fmt.Println("Send operation on ch1 works!")
			t.Reset(delay)
			continue
		case v := <-ch2:
			fmt.Println("Receive operation on ch2 works!", len(ch2), v)
			t.Reset(delay)
			continue
		case _ = <-t.C:
			fmt.Println("ch1 & ch2 has no value, repeat time interval:", delay)
			t.Reset(delay)
		}
	}
}

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func(ch chan int) { ch <- 1 }(ch1)
	go func(ch chan int) {
		for i := 0; i < 5; i++ {
			ch <- i
			time.Sleep(100 * time.Millisecond)
		}
	}(ch2)

	time.Sleep(time.Millisecond)
	delay := 12 * time.Second
	chanProcess(delay, ch1, ch2)
}
