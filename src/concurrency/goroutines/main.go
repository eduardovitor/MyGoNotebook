package main

import (
	"fmt"
	"time"
)

func message(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go message("goroutine")
	message("normal")
}
