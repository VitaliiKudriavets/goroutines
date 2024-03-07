package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go sayHello(ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	Say("The End")
}

func Say(word string) {
	fmt.Println(word)
}

func sayHello(exit chan int) {
	for i := 0; i <= 5; i++ {
		time.Sleep(100 * time.Millisecond)
		exit <- i
	}
}
