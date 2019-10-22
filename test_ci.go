package main

import (
	"fmt"
	"test_ci/models/counter"
)

func F1() string {
	return "F1"
}

func F2() string {
	return "F2"
}

func F3() string {
	return "F3"
}

func F4() int {
	counter.Inc()
	counter.Inc()
}

func main() {
	fmt.Println("welcome")
}
