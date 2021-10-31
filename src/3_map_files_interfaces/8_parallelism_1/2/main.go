// https://stepik.org/lesson/360357/step/9?unit=344766

func task2(c chan string, str string) {
	for i := 0; i < 5; i++ {
		c <- str + " "
	}
}