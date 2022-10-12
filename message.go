package gochat

import (
	"context"
	"fmt"
	"time"
)

type SentMessage struct {
	id       string
	sender   string
	toGroups []GroupRecipients
	// Before: [G1, G2]
	// After:  [{id: G1, participants: []}]
	toContacts []string
	content    []byte
	createdAt  time.Time
	sentAt     time.Time
	received   []recipient
	seen       []recipient
}

type GroupRecipients struct {
	id           string
	participants []string
}

type recipient struct {
	by, in string
	at     time.Time
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
) (*SentMessage, []ReceivedMessage, error) {
	var (
		id  = "" // TODO: generate uuid
		now = time.Now()
		rms = make(map[string]ReceivedMessage, 0)
	)

	groupIDs := make([]string, 0, len(toGroups))
	for _, g := range toGroups {
		groupIDs = append(groupIDs, g.ID())
		for _, p := range g.participants {
			rm, found := rms[p.ID()]
			if found {
				rm.in = append(rm.in, g.ID())
				continue
			}

			rms[p.ID()] = ReceivedMessage{
				id:        fmt.Sprintf("%s-%s", p.ID(), id), // TODO: need to discuss
				receiver:  p.ID(),
				sender:    sender.ID(),
				in:        []string{g.ID()},
				content:   content,
				createdAt: createdAt,
				sentAt:    now,
			}
		}
	}

	contactIDs := make([]string, 0, len(toContacts))
	for _, u := range toContacts {
		contactIDs = append(contactIDs, u.ID())
		rm, found := rms[u.ID()]
		if found {
			rm.in = append(rm.in, DM)
			continue
		}

		rms[u.ID()] = ReceivedMessage{
			id:        fmt.Sprintf("%s-%s", u.ID(), id), // TODO: need to discuss
			receiver:  u.ID(),
			sender:    sender.ID(),
			in:        []string{DM},
			content:   content,
			createdAt: createdAt,
			sentAt:    now,
		}
	}

	sm := SentMessage{
		id:         fmt.Sprintf("%s-%s", sender.ID(), id),
		sender:     sender.ID(),
		toGroups:   groupIDs,
		toContacts: contactIDs,
		content:    content,
		createdAt:  createdAt,
		sentAt:     now,
		received:   make([]recipient, 0, 0),
		seen:       make([]recipient, 0, 0),
	}

	return &sm, mapToArr(rms), nil
}

func mapToArr(rms map[string]ReceivedMessage) []ReceivedMessage {
	arr := make([]ReceivedMessage, 0, len(rms))
	for _, rm := range rms {
		arr = append(arr, rm)
	}
	return arr
}

type SentMessageRepository interface {
	Insert(context.Context, SentMessage) (*SentMessage, error)
}

type Page[T any] struct {
	Data       []T // TODO: use generics here
	NextCursor string
}

type ReceivedMessageRepository interface {
	Insert(context.Context, ReceivedMessage) (*ReceivedMessage, error)
	FindAll(ctx context.Context, pageSize int, pageCursor string) (*Page, error)
}
