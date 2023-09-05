package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	for {
		go pa()
		go pb()
		time.Sleep(1e9)

	}

}

func pa() {

	fmt.Print("A ")
	runtime.Gosched()

}
func pb() {
	fmt.Print("B ")
	runtime.Gosched()

}
