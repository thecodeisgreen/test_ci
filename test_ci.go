package main

import (
	"fmt"
	"test_ci/models/counter"

	"test_ci/tcig.io/version"
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

func main() {
	c := counter.New(15)
	c.Inc()
	c.Inc()
	fmt.Println(version.Get().String())
	fmt.Println("welcome", c.Get())
}
