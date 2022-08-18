package main

import (
	"fmt"
	"time"
)

func main() {

	channel_1 := make(chan string)
	channel_2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		channel_1 <- "channel-1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		channel_2 <- "channel-2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-channel_1:
			fmt.Println("received", msg1)
		case msg2 := <-channel_2:
			fmt.Println("received", msg2)
		}
	}
}
