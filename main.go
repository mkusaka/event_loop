// server.go
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	go watchQueue()
	for {
		fmt.Println("enter some world")
		input := bufio.NewScanner(os.Stdin)

		input.Scan()

		event := NewEvent(func(e *Event) {
			e.start()
			fmt.Println(input.Text())
			rand.Seed(time.Now().UnixNano())
			i := rand.Intn(2)
			if i == 1 {
				fmt.Println("random 1 == 1, end")
				e.done()
			} else {
				fmt.Println("random 1 != 1")
			}
		})
		addQueue(event)
		fmt.Println("input is ", input.Text())
	}
}

type EventType = string

const (
	wait EventType = "wait"
)

type State = string

const (
	waiting State = "waiting"
	running State = "running"
	done    State = "done"
)

type Event struct {
	eventType string
	processor func(e *Event)
	state     State
}

func NewEvent(processor func(e *Event)) *Event {
	return &Event{eventType: wait, state: waiting, processor: processor}
}

func (e *Event) start() {
	e.state = running
}

func (e *Event) done() {
	e.state = done
}

var queue []*Event

func addQueue(task *Event) {
	queue = append(queue, task)
}

var tick = time.Millisecond

func watchQueue() {
	fmt.Println("start watch")
	for {
		if len(queue) != 0 {
			popper := queue[0]

			// task process
			fmt.Println("dequeued", popper)
			popper.processor(popper)

			if popper.state == done {
				fmt.Println("event is done.")
				queue = queue[1:]
			} else {
				fmt.Printf("event is %s\n", popper.state)
			}
		}
		time.Sleep(tick)
	}
}
