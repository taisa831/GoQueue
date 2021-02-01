package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("test")

	q := make([]int, 0)

	for i := 0; i < 100; i++ {
		q = append(q, i)
	}

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			doProcess(i)
		}(i)
		q = q[1:]
	}
	wg.Wait()
	fmt.Println("処理終了")
}

func doProcess(i int) {
	fmt.Println(i)
	time.Sleep(10 * time.Second)
}
