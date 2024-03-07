package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	data := make(chan int)
	exit := make(chan int)
	go sayHello(ch)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-data)
		}
		exit <- 0
	}()
	selectOne(data, exit)

	for i := range ch {
		fmt.Println(i)
	}

	Say("The End")
}

func Say(word string) {
	fmt.Println(word)
}

func sayHello(exit chan int) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		exit <- i
	}
	close(exit)
}

func selectOne(data, exit chan int) {
	x := 0
	for {
		select {
		case data <- x:
			x += 1
		case <-exit:
			fmt.Println("exit")
			return
		}
	}
}
