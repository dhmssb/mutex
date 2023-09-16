package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(4)

	bathroom := sync.Mutex{}

	takeAShower := func(name string) {
		defer wg.Done()

		fmt.Printf("%s: I want to take a shower. I'm trying to acquire the bathroom\n", name)
		bathroom.Lock()
		fmt.Printf("%s: I have the bathroom now, taking a shower\n", name)
		time.Sleep(500 * time.Microsecond)
		fmt.Printf("%s: I'm done, I'm unlocking the bathroom\n", name)
		bathroom.Unlock()
	}

	go takeAShower("Maman")
	go takeAShower("Budi")
	go takeAShower("Dudung")
	go takeAShower("Riko")

	wg.Wait()
	fmt.Println("main: Everyone is Done. Shutting down...")
}
