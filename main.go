package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/bradhe/stopwatch"
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


	sw := stopwatch.Start()
	for i := 0; i < 1000000; i++ {
		SayHello("Index ke: " + strconv.Itoa(i))
	}
	sw.Stop()

	fmt.Println(".....")

	ChannelExample()

	time.Sleep(1 * time.Microsecond)

	// print execute time
	fmt.Println("Execute Time: ", sw.Milliseconds(), "")
}