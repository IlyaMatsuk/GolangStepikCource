// https://stepik.org/lesson/360357/step/8?unit=344766

func task(c chan int, N int) {
	c <- N + 1
}