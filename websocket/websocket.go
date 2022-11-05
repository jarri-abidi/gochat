package websocket

import (
	"context"
	"net"

	"github.com/gorilla/websocket"
	"github.com/jarri-abidi/gochat"
	"github.com/jarri-abidi/gochat/messaging"
	"github.com/pkg/errors"
)

type Connections struct {
	users map[string]websocket.Conn
}

func (c *Connections) Relay(ctx context.Context, msg gochat.ReceivedMessage) error {
	conn, ok := c.users[msg.RecipientID()]
	if !ok {
		return errors.New("could not find recipient")
	}

	if err := conn.WriteJSON(msg); err != nil {
		return errors.Wrap(err, "could not write msg to recipient")
	}

	return nil
}

// connections with other chat servers
// connections with actual users
type State map[net.Addr]websocket.Conn

type Publisher struct {
	msgs chan<- messaging.RelayEvent
}

func (p *Publisher) PublishRelayEvent(ctx context.Context, event messaging.RelayEvent) error {
	p.msgs <- event
	return nil
}

type Consumer struct {
	msgs <-chan messaging.RelayEvent
}

func (c *Consumer) ConsumeRelayEvent(ctx context.Context) (*messaging.RelayEvent, error) {
	msg := <-c.msgs
	return &msg, nil
}

// server <---> client

// server can actually call the client
// ^impossible in typical client-server architecture
