package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	WaitGroup()
	ErrGroupWait()
}

func WaitGroup() {
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

func ErrGroupWait() {
	q2 := make([]int, 0)

	for i := 0; i < 100; i++ {
		q2 = append(q2, i)
	}

	var eg errgroup.Group
	for i := 0; i < 100; i++ {
		i := i
		eg.Go(func() error {
			doProcess(i)
			return nil
		})
		q2 = q2[1:]
	}

	if err := eg.Wait(); err != nil {
		panic(err)
	}
	fmt.Println("処理終了")
}

func doProcess(i int) {
	fmt.Println(i)
	time.Sleep(10 * time.Second)
}
