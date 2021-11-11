package main

import (
	"runtime"
	"time"
)

func main() {
	quit := make(chan bool)
	for i := 0; i != runtime.NumCPU(); i++ {
		go func() {
			for {
				select {
				case <-quit:
					break
				default:
				}
			}
		}()
	}

	time.Sleep(time.Second * 15)
	for i := 0; i != runtime.NumCPU(); i++ {
		quit <- true
	}
}