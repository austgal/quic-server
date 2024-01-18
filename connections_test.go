package main

import (
	"testing"

	"github.com/quic-go/quic-go"
	"github.com/stretchr/testify/assert"
)

func TestAddConnectionSub(t *testing.T) {
	connections := &Connections{
		subscribers: make(map[quic.Stream]struct{}),
		publishers:  make(map[quic.Stream]struct{}),
	}

	connection := &MockStream{}
	connections.addConnection(connection, connections.subscribers)

	assert.Contains(t, connections.subscribers, connection)
}

func TestAddConnectionPub(t *testing.T) {
	connections := &Connections{
		subscribers: make(map[quic.Stream]struct{}),
		publishers:  make(map[quic.Stream]struct{}),
	}

	connection := &MockStream{}
	connections.addConnection(connection, connections.publishers)

	assert.Contains(t, connections.publishers, connection)
}

func TestRemoveSubscriber(t *testing.T) {
	connections := &Connections{
		subscribers: make(map[quic.Stream]struct{}),
		publishers:  make(map[quic.Stream]struct{}),
	}

	connection := &MockStream{}
	connections.subscribers[connection] = struct{}{}
	connections.removeConnection(connection, connections.subscribers)

	assert.NotContains(t, connections.subscribers, connection)
}

func TestRemovePublisher(t *testing.T) {
	connections := &Connections{
		subscribers: make(map[quic.Stream]struct{}),
		publishers:  make(map[quic.Stream]struct{}),
	}

	connection := &MockStream{}
	connections.publishers[connection] = struct{}{}
	connections.removeConnection(connection, connections.publishers)

	assert.NotContains(t, connections.publishers, connection)
}
