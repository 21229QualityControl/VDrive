package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	// Get binds
	read := bufio.NewReader(conn)
	line, err := read.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = line[:len(line)-1] // Remove newline

	// Parse binds
	binds := make(map[string]string)
	for _, bind := range strings.Split(line, ";") {
		if len(bind) == 0 {
			continue
		}
		val := strings.Split(bind, ": ")
		binds[val[0]] = val[1]
	}
	fmt.Println(binds)

	// Send key
	_, err = conn.Write([]byte("W:true\n"))
	if err != nil {
		panic(err)
	}
}
