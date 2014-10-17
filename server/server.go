package main

import (
	"fmt"
	"net"
	"encoding/gob"
	"time"
)

const (
	_ = iota
	TASK_CREATE
	TASK_SWITCH
)

type JobPackage struct {
	Action int
}

type Task struct {
	name string
	start, end time.Time
	session_start time.Time
	elapsed time.Duration
}
/*
func handleConnection(conn net.Conn) {
}
*/

func main () {
	fmt.Println("Hello")

	ln, err := net.Listen("tcp", ":9876")

	if err != nil {
		// handle error
	}

	task_list = make([]*Task, 1000)

	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			continue
		}

		dec := gob.NewDecoder(conn)
		p := &JobPackage{}
		dec.Decode(p)
		fmt.Printf("Received : %+v", p);

		switch (p.Action) {
			case TASK_CREATE:


			case TASK_SWITCH:
		}
	}
}
