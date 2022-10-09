package gochat

import (
	"fmt"
	"time"
)

type SentMessage struct {
	id        string
	toGroups  []string
	toUsers   []string
	content   []byte
	createdAt time.Time
	sentAt    time.Time
	received  []recipient
	seen      []recipient
}

type recipient struct {
	by, in string
	at     time.Time
}

type ReceivedMessage struct {
	id          string
	in          []string
	messageFrom string
	content     []byte
	createdAt   time.Time
	sentAt      time.Time
	receivedAt  time.Time
}

const DM = "DM"

// NewMessage accepts a list of groups and users to which the message needs to be sent.
// It also accepts the content as a slice of bytes. The value of createdAt should be the time
// when the sender created the message as opposed to when it was received by the server.
func NewMessage(
	sender User,
	toGroups []Group,
	toUsers []User,
	content []byte,
	createdAt time.Time,
) (*SentMessage, map[string]ReceivedMessage, error) {
	var (
		id  = "" // TODO: generate uuid
		now = time.Now()
		rms = make(map[string]ReceivedMessage, 0)
	)

	groupIDs := make([]string, 0, len(toGroups))
	for _, g := range toGroups {
		groupIDs = append(groupIDs, g.id)
		for _, p := range g.participants {
			rm, found := rms[p.id]
			if found {
				rm.in = append(rm.in, g.id)
				continue
			}

			rms[p.id] = ReceivedMessage{
				id:          fmt.Sprintf("%s-%s", p.id, id), // TODO: need to discuss
				in:          []string{g.id},
				messageFrom: sender.id,
				content:     content,
				createdAt:   createdAt,
				sentAt:      now,
			}
		}
	}

	userIDs := make([]string, 0, len(toUsers))
	for _, u := range toUsers {
		userIDs = append(userIDs, u.id)
		rm, found := rms[u.id]
		if found {
			rm.in = append(rm.in, DM)
			continue
		}

		rms[u.id] = ReceivedMessage{
			id:          fmt.Sprintf("%s-%s", u.id, id), // TODO: need to discuss
			in:          []string{DM},
			messageFrom: sender.id,
			content:     content,
			createdAt:   createdAt,
			sentAt:      now,
		}
	}

	sm := SentMessage{
		id:        fmt.Sprintf("%s-%s", sender.ID(), id),
		toGroups:  groupIDs,
		toUsers:   userIDs,
		content:   content,
		createdAt: createdAt,
		sentAt:    now,
		received:  make([]recipient, 0, 0),
		seen:      make([]recipient, 0, 0),
	}

	return &sm, rms, nil
}
