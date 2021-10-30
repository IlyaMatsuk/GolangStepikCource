package main

import (
	"fmt"
	"time"
)

// https://stepik.org/lesson/359395/step/8?unit=343626

const now = 1589570165

func main() {
	var m, s time.Duration
	fmt.Scanf("%d мин. %d сек.", &m, &s)
	t := time.Unix(now, 0).UTC().Add(m * time.Minute).Add(s * time.Second)
	fmt.Println(t.Format(time.UnixDate))
}
