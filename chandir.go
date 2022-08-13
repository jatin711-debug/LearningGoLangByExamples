package main

import "fmt"


func ping(ch1 chan <-int,mssg int){
	ch1 <- mssg
}

func pong(ch1 <-chan int,ch2 chan<- int) {
	mssg := <-ch1
	ch2 <- mssg
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int,2)
	sum(s[:len(s)/2], c)
	sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
