package main

// https://stepik.org/lesson/345547/step/13?unit=329291

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		select {
		case f_data := <-firstChan:
			out <- f_data << 1
		case s_data := <-secondChan:
			out <- s_data * 3
		case <-stopChan:
			break
		}
	}()
	return out
}

func main() {

}
