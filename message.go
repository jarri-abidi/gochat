package gochat

import (
	"context"
	"fmt"
	"time"
)

type SentMessage struct {
	id         string
	sender     string
	content    []byte
	createdAt  time.Time
	sentAt     time.Time
	recipients recipients
}

type recipients map[string][]status

type status struct {
	in        string
	delivered time.Time
	seen      time.Time
}

func (sm SentMessage) SenderID() string { return sm.sender }

type ReceivedMessage struct {
	id         string
	receiver   string
	sender     string
	in         []string
	content    []byte
	createdAt  time.Time
	sentAt     time.Time
	receivedAt time.Time
}

func (rm ReceivedMessage) ReceiverID() string { return rm.receiver }

const DM = "DM"

// NewMessage accepts a list of groups and users to which the message needs to be sent.
// It also accepts the content as a slice of bytes. The value of createdAt should be the time
// when the sender created the message as opposed to when it was received by the server.
func NewMessage(
	sender User,
	toGroups []Group,
	toContacts []Contact,
	content []byte,
	createdAt time.Time,
) (*SentMessage, error) {
	var (
		id         = "" // TODO: generate uuid
		now        = time.Now()
		recipients = make(recipients, 0)
	)

	for _, g := range toGroups {
		for _, p := range g.participants {
			statuses, found := recipients[p.ID()]
			if found {
				recipients[p.ID()] = append(statuses, status{in: g.ID()})

				continue
			}

			recipients[p.ID()] = append(make([]status, 0), status{in: g.ID()})
		}
	}

	for _, c := range toContacts {
		statuses, found := recipients[c.ID()]
		if found {
			recipients[c.ID()] = append(statuses, status{in: DM})
			continue
		}

		recipients[c.ID()] = append(make([]status, 0), status{in: DM})
	}

	sm := SentMessage{
		id:         fmt.Sprintf("%s-%s", sender.ID(), id),
		sender:     sender.ID(),
		content:    content,
		createdAt:  createdAt,
		sentAt:     now,
		recipients: recipients,
	}

	return &sm, nil
}

type SentMessageRepository interface {
	Insert(context.Context, SentMessage) (*SentMessage, error)
}

type Page[T any] struct {
	Data       []T
	NextCursor string
}

type ReceivedMessageRepository interface {
	Insert(context.Context, ReceivedMessage) (*ReceivedMessage, error)
	FindAll(ctx context.Context, pageSize int, pageCursor string) (*Page[ReceivedMessage], error)
}
