package main

import (
	"time"

	"github.com/gorilla/websocket"
)

const (
    // Time allowed to write message to peer
    writeWait = 10 * time.Second

    // Time allowed to read the next message from peer
    readWait = 60 * time.Second

    // Send pings to peer with this period. Must be less than readWait
    pingPeriod = (readWait * 9) / 10

    // Max message size allowed from peer
    maxMessageSize = 512
)

var (
    newline = []byte{'\n'}
    space = []byte{' '}
)

var upgrader = websocket.Upgrader{
    ReadBufferSize: 1024,
    WriteBufferSize: 1024,
}

type Client struct {
    hub *Hub

    conn *websocket.Conn

    send chan []byte
}
