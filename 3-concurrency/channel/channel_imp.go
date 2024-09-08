package channel

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/exp/rand"
)

func ChannelImpl() {
	ch := make(chan string)

	sayHelloTo := func(person string) {
		ch <- fmt.Sprintf("hello %s", person)
	}

	go sayHelloTo("rama")
	go sayHelloTo("rami")
	go sayHelloTo("ramu")

	for i := 0; i < 3; i++ {
		fmt.Printf("%d-%s\n", i+1, <-ch)
	}
}

func Iteration() {
	var messages = make(chan string)

	for _, each := range []string{"wick", "hunt", "bourne"} {
		go func(who string) {
			var data = fmt.Sprintf("hello %s", who)
			messages <- data
		}(each)
	}

	for i := 0; i < 3; i++ {
		fmt.Println(<-messages)
	}
}

func ChannelWithTimeout() {
	ch := make(chan string)

	sendData := func(msg string) {
		time.Sleep(5000 * time.Millisecond)
		ch <- msg
	}

	receiver := func(ch <-chan string) {
		select {
		case data := <-ch:
			fmt.Print("send data ", data, "\n")
		case <-time.After(time.Second * 2):
			fmt.Print("got timeout\n")
		}
	}

	for i := 0; i < 5; i++ {
		go sendData(fmt.Sprintf("hello-%d", i+1))
	}

	for i := 0; i < 5; i++ {
		receiver(ch)
	}
}

// this is the best architecture
// called fan in and fan out
func ChannelGroup() {
	var wg sync.WaitGroup
	wg.Add(5)
	ch := make(chan string)

	start := time.Now()
	fmt.Println("start")

	sendData := func(msg string) {
		duration := time.Duration(rand.Intn(5)+1) * time.Second
		fmt.Println(duration)
		time.Sleep(duration)
		ch <- msg
	}

	receiver := func(ch <-chan string, wg *sync.WaitGroup) {
		select {
		case data := <-ch:
			fmt.Print("send data ", data, "\n")
		case <-time.After(time.Second * 2):
			fmt.Print("got timeout\n")
		}

		wg.Done()
	}

	for i := 0; i < 5; i++ {
		go sendData(fmt.Sprintf("hello-%d", i+1))
	}

	// receiver now using go routine
	for i := 0; i < 5; i++ {
		go receiver(ch, &wg)
	}

	// current time taken is the longest processing, not accumulation all of processing
	end := time.Now()
	fmt.Printf("finish %s\n", end.Sub(start))

	wg.Wait()
}
