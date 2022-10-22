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
	recipients Recipients
}

type Recipients map[string][]Status

type Status struct {
	in        string
	delivered time.Time
	seen      time.Time
}

func (sm SentMessage) SenderID() string       { return sm.sender }
func (sm SentMessage) Recipients() Recipients { return sm.recipients }
func (sm SentMessage) Content() []byte        { return sm.content }

type ReceivedMessage struct {
	id         string
	recipient  string
	sender     string
	in         []string
	content    []byte
	createdAt  time.Time
	sentAt     time.Time
	receivedAt time.Time
}

func (rm ReceivedMessage) RecipientID() string { return rm.recipient }
func (rm ReceivedMessage) SenderID() string    { return rm.sender }

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
	recipients := make(Recipients, 0)

	for _, g := range toGroups {
		for _, p := range g.participants {
			statuses, found := recipients[p.ID()]
			if found {
				recipients[p.ID()] = append(statuses, Status{in: g.ID()})
				continue
			}
			recipients[p.ID()] = append(make([]Status, 0), Status{in: g.ID()})
		}
	}

	for _, c := range toContacts {
		statuses, found := recipients[c.ID()]
		if found {
			recipients[c.ID()] = append(statuses, Status{in: DM})
			continue
		}
		recipients[c.ID()] = append(make([]Status, 0), Status{in: DM})
	}

	sm := SentMessage{
		id:         "", // generate uuid
		sender:     sender.ID(),
		content:    content,
		createdAt:  createdAt,
		sentAt:     time.Now(),
		recipients: recipients,
	}

	return &sm, nil
}

func GetReceivedMessages(sm SentMessage) []ReceivedMessage {
	var (
		id  = "" // TODO: generate uuid
		now = time.Now()
		rms = make(map[string]ReceivedMessage, 0)
	)

	for recipientID, statuses := range sm.Recipients() {
		rm := ReceivedMessage{
			id:        fmt.Sprintf("%s-%s", sm.ID(), id), // TODO: need to discuss
			receiver:  p.ID(),
			sender:    sender.ID(),
			in:        []string{g.ID()},
			content:   content,
			createdAt: createdAt,
			sentAt:    now,
		}
	}
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
