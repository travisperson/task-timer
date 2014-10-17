package main

import (
	"fmt"
	"net"
	"encoding/gob"
)

// comment

const (
	_ = iota
	CREATE_TASK
	SWITCH_TASK
)

type JobPackage struct {
	Action int
}

func main () {
	fmt.Println("Hello")
	conn, err := net.Dial("tcp", "localhost:9876")
	if err != nil {
		fmt.Println("Error")
	}
	encoder := gob.NewEncoder(conn)
	p := &JobPackage{CREATE_TASK}
	encoder.Encode(p)
	conn.Close()
	fmt.Println("done");
}
