package main

import (
	"fmt"
	"time"
)

func main() {
	// var wg sync.WaitGroup
	// wg.Add(2)

	// go func() {
	// 	count("fish")
	// 	count("sheep")
	// 	wg.Done()
	// }()

	// go func() {
	// 	count("bears")
	// 	wg.Done()
	// }()
	// wg.Wait()

	// =======================================================
	c := make(chan string)
	go count("Sheep", c)

	for msg := range c { // continues to iterate over channel until it's closed
		fmt.Println(msg)
	}

}

func count(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing // send value to channel c
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}

	close(c) // close only from sender, otherwise error
}
