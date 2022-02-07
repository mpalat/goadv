package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	fmt.Println("main started")
	n := 3
	// create 20 tasks
	tasks := make([]func(), n)
	for i := 0; i < n; i++ {
		wg.Add(1)
		tasks[i] = func(idx int) func() {
			return func() {
				fmt.Println("Task ", idx, " invoked")
				wg.Done()
			}
		}(i)
	}
	fmt.Printf("%d tasks initiated", n)
	// initiate the execution of all tasks
	for _, task := range tasks {
		go task()
	}
	// wait for the tasks to complete
	wg.Wait()
	fmt.Println("main finished")

}

func f(i int, initWg, doneWg sync.WaitGroup) {
	initWg.Done()
	fmt.Printf("Task No:%d\n", i)
	doneWg.Done()
}
