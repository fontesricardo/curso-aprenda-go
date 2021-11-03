package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	fmt.Println("CPUs:", runtime.NumCPU())

	var contador int64
	totaldegoroutines := 10

	var wg sync.WaitGroup
	wg.Add(totaldegoroutines)

	for i := 0; i < totaldegoroutines; i++ {
		go func() {
			runtime.Gosched()
			atomic.AddInt64(&contador, 1)
			fmt.Println("Contador:\t", atomic.LoadInt64(&contador))
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Valor contador:", contador)
}
