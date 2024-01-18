package main

import (
	"context"
	"log"

	"github.com/quic-go/quic-go"
)

const no_sub_message = "No subscribers are connected"

func (c *Connections) handlePublisher(connection quic.Connection) {
	log.Printf("New publisher connected: %v\n", connection.RemoteAddr())
	for {
		stream, err := connection.AcceptStream(context.Background())
		if err != nil {
			log.Println(err)
			return
		}
		c.addConnection(stream, c.publishers)

		go c.handlePubStream(stream, connection)
	}
}

func (c *Connections) handlePubStream(stream quic.Stream, connection quic.Connection) {

	buf := make([]byte, 1024)

	for {
		if len(c.subscribers) == 0 {
			_, err := stream.Write([]byte(no_sub_message))
			if err != nil {
				log.Println(err)
				return
			}
			break
		}

		n, err := stream.Read(buf)
		if err != nil {
			log.Println(err)
			return
		}

		c.broadcastMessage(buf[:n])
	}
}

// TODO: looks the same as function in subscriber.go
func (c *Connections) broadcastMessage(message []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for subscriber := range c.subscribers {
		go func(sub quic.Stream) {
			_, err := sub.Write(message)
			log.Printf("Broadcasting message from publisher: %v\n", string(message))
			if err != nil {
				log.Println(err)
			}
		}(subscriber)
	}
}
