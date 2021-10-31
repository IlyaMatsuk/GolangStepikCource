package main

import (
	"fmt"
	"sync"
	"time"
)

// https://stepik.org/lesson/345547/step/5?unit=329291

func work() {
	time.Sleep(1 * time.Second)
	fmt.Println("Done")
}

func main() {
	ws := new(sync.WaitGroup)
	count := 10
	ws.Add(count)
	for i := 0; i < count; i++ {
		go func() {
			work()
			ws.Done()
		}()
	}
	ws.Wait()
}
