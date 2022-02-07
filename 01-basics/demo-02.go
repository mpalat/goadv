package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	fmt.Println("main started")
	wg.Add(1)
	go f1()
	f2()
	wg.Wait()
	fmt.Println("main finished")

}

func f1() {
	fmt.Println("F1")
	wg.Done()
}
func f2() {
	fmt.Println("F2")
}
