package main

import (
	"context"
	"log"

	"github.com/quic-go/quic-go"
)

const new_sub_msg = "Subscriber connected"

func (c *Connections) handleSubscriber(connection quic.Connection) {
	log.Printf("New subscriber connected: %v\n", connection.RemoteAddr())

	for {
		stream, err := connection.AcceptStream(context.Background())
		if err != nil {
			log.Println(err)
			return
		}
		c.addConnection(stream, c.subscribers)
		go c.informPublishers([]byte(new_sub_msg))
	}
}

func (c *Connections) informPublishers(message []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for publisher := range c.publishers {
		go func(pub quic.Stream) {
			_, err := pub.Write([]byte(message))
			if err != nil {
				log.Println(err)
			}
			log.Printf("informing publisher : %v\n", string(message))
		}(publisher)
	}
}
