package main

import "sync"

func main() {
	var a string
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		a = "hello world"
		wg.Done()
	}()
	wg.Wait()
	println(a)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(x int) {
			sendRPC(x)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
func sendRPC(i int) {
	println(i)
}
