package main

import (
	"bufio"
	"encoding/json"
	"net"
	"strconv"
	"strings"

	"github.com/andlabs/ui"
	"github.com/gorilla/websocket"
)

var conn net.Conn
var stream *websocket.Conn
var room string
var nameModel *NameModel
var bindsModel *BindsModel

type HostParams struct {
	RoomName string            `json:"room_name"`
	Binds    map[string]string `json:"binds"`
}

type EventKind string

const (
	EventKindJoin  EventKind = "join"
	EventKindLeave EventKind = "leave"
	EventKindKey   EventKind = "key"
)

type Event struct {
	Kind  EventKind `json:"kind"`
	Value string    `json:"value"`
}

func listen() {
	// Connect to robot
	var err error
	conn, err = net.Dial("tcp", hostname.Text()+":"+strconv.Itoa(port.Value()))
	handle(err)

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
	bindsModel.SetBinds(binds)

	// Connect to server
	stream, _, err = websocket.DefaultDialer.Dial("wss://vdrive.nv7haven.com/host", nil)
	handle(err)

	// Send params
	pars := HostParams{
		RoomName: room,
		Binds:    binds,
	}
	parsJson, err := json.Marshal(pars)
	handle(err)
	err = stream.WriteMessage(websocket.TextMessage, parsJson)
	handle(err)

	// Listen
	go func() {
		for isHosting {
			_, msg, err := stream.ReadMessage()
			if err != nil {
				if isHosting {
					cleanup()
				}
				return
			}

			var ev Event
			err = json.Unmarshal(msg, &ev)
			handle(err)

			switch ev.Kind {
			case EventKindJoin:
				nameModel.AddName(ev.Value)

			case EventKindLeave:
				nameModel.RemoveName(ev.Value)

			case EventKindKey:
				_, err = conn.Write([]byte(ev.Value))
				handle(err)
			}
		}
	}()

	isHosting = true
}

func cleanup() {
	ui.QueueMain(func() {
		btn.SetText("Stopping...")
		btn.Disable()
	})

	go func() {
		isHosting = false

		err := conn.Close()
		handle(err)
		err = stream.Close()
		handle(err)

		nameModel.Clear()
		bindsModel.Clear()

		ui.QueueMain(func() {
			hostname.SetReadOnly(false)
			roomname.SetReadOnly(false)
			btn.SetText("Host")
			btn.Enable()
		})
	}()
}
