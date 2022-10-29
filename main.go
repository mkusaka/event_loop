// server.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	go watchQueue()
	for {
		fmt.Println("enter some world")
		input := bufio.NewScanner(os.Stdin)

		input.Scan()

		addQueue(input.Text())
		fmt.Println("input is ", input.Text())
	}
}

var queue []string

func addQueue(task string) {
	queue = append(queue, task)
}

var tick = time.Millisecond

func watchQueue() {
	fmt.Println("start watch")
	for {
		if len(queue) != 0 {
			popper := queue[0]
			queue = queue[1:]

			// task process
			fmt.Println("dequeued", popper)
		}
		time.Sleep(tick)
	}
}
