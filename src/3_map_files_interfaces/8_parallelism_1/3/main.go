// https://stepik.org/lesson/360357/step/10?unit=344766

func removeDuplicates(inputStream chan string, outputStream chan string) {
	defer close(outputStream)
	var prev_str string
	for str := range inputStream {
		if prev_str != str {
			outputStream <- str
		}
		prev_str = str
	}
}