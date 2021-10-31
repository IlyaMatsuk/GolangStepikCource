package main

import (
	"fmt"
	"time"
)

// https://stepik.org/lesson/345547/step/3?unit=329291

func work() {
	time.Sleep(1 * time.Second)
	fmt.Println("Done")
}

func main() {
	done := make(chan struct{})
	go func() {
		work()
		close(done)
	}()
	<-done
}
