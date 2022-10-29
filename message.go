package gochat

import (
	"context"
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

func (sm SentMessage) ID() string             { return sm.id }
func (sm SentMessage) SenderID() string       { return sm.sender }
func (sm SentMessage) Recipients() Recipients { return sm.recipients }
func (sm SentMessage) Content() []byte        { return sm.content }
func (sm SentMessage) CreatedAt() time.Time   { return sm.createdAt }
func (sm SentMessage) SentAt() time.Time      { return sm.sentAt }

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

func (rm ReceivedMessage) ID() string          { return rm.id }
func (rm ReceivedMessage) RecipientID() string { return rm.recipient }
func (rm ReceivedMessage) SenderID() string    { return rm.sender }

const DM = "DM"

// NewSentMessage accepts a list of groups and contacts to which the message needs to be sent.
// It also accepts the content as a slice of bytes. The value of createdAt should be the time
// when the sender created the message as opposed to when it was received by the server.
func NewSentMessage(
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

func NewReceivedMessages(sm SentMessage) []ReceivedMessage {
	rms := make([]ReceivedMessage, 0)

	for recipientID, statuses := range sm.Recipients() {
		in := make([]string, 0, len(statuses))
		for _, status := range statuses {
			in = append(in, status.in)
		}

		rms = append(rms, ReceivedMessage{
			id:        sm.ID(),
			recipient: recipientID,
			sender:    sm.SenderID(),
			in:        in,
			content:   sm.Content(),
			createdAt: sm.CreatedAt(),
			sentAt:    sm.SentAt(),
		})
	}

	return rms
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
