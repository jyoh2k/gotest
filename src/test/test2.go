package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wait sync.WaitGroup
	defer func() {
		fmt.Println("main end!")
	}()

	myChannel_1 := make(chan int)
	myChannel_2 := make(chan string)

	wait.Add(1)
	go func() {
		defer wait.Done()
		fmt.Println("Func1")
		time.Sleep(time.Second)
		myChannel_1 <- 1
		close(myChannel_1)
	}()

	wait.Add(1)
	go func() {
		defer wait.Done()
		val, flag := <-myChannel_2
		if !flag {
			fmt.Println("channel close")
		} else {
			fmt.Println("func :", val, flag)
		}
	}()

	val, flag := <-myChannel_1
	if !flag {
		fmt.Println("channel close")
	} else {
		fmt.Println("main :", val, flag)
	}
	myChannel_2 <- "hi"
	close(myChannel_2)

	wait.Wait()
}
