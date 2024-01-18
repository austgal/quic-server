package main

import (
	"sync"

	"github.com/quic-go/quic-go"
)

type Connections struct {
	subscribers map[quic.Stream]struct{}
	publishers  map[quic.Stream]struct{}
	mu          sync.Mutex
}

func (c *Connections) addConnection(connection quic.Stream, connectionMap map[quic.Stream]struct{}) {
	c.mu.Lock()
	connectionMap[connection] = struct{}{}
	c.mu.Unlock()
}

func (c *Connections) removeConnection(connection quic.Stream, connectionMap map[quic.Stream]struct{}) {
	c.mu.Lock()
	delete(connectionMap, connection)
	c.mu.Unlock()
}
