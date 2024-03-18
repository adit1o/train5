package main

import (
	"fmt"
	"strconv"
	"time"
)



func SayHello(s string) {
	fmt.Println("Hello " + s)
}

func ChannelExample() {
	ch := make(chan string)
	defer close(ch)

	// anon func
	go func() {
		time.Sleep(2 * time.Second)
		ch <- "Hello World"
		fmt.Println("send done.")
	}()

	output := <- ch
	fmt.Println("Channel Value: ", output)
	time.Sleep(1 * time.Second)
}

func main() {
	fmt.Println(".....")

	for i := 0; i < 1000000; i++ {
		go SayHello("Index ke: " + strconv.Itoa(i))
	}

	fmt.Println(".....")

	ChannelExample()

	time.Sleep(1 * time.Microsecond)
}