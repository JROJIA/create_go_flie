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
		go pb()
		go pd()
		time.Sleep(1e9)

	}

}

func pa() {

	fmt.Print("A ")
	runtime.Gosched()

}
func pb() {
	fmt.Print("c ")
	runtime.Gosched()

}
func pd() {
	fmt.Print("d ")
	runtime.Gosched()

}
