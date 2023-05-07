package main

import (
	"fmt"
	"strconv"
	"time"
)

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func counter(c chan string) {
	for i := 0; ; i++ {
		c <- strconv.Itoa(i)
	}
}

func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func routineFunc(c chan string, chanId int) {
	for {
		c <- strconv.Itoa(chanId)
	}
}

func main() {
	var c chan string = make(chan string)

	go pinger(c)
	for i := 0; i < 10; i++ {
		go routineFunc(c, i)
	}
	//go counter(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
	fmt.Println(input)
}
