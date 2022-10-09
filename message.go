package gochat

import (
	"time"
)

type SentMessage struct {
	id        string
	toGroups  []string
	toUsers   []string
	content   []byte
	createdAt time.Time
	sentAt    time.Time
	received  []stuff
	seen      []stuff
}

type stuff struct {
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

// NewMessage accepts a list of group IDs and user IDs to which the message needs to be sent.
// It also accepts the content as a slice of bytes. The value of createdAt should be the time
// when the sender created the message as opposed to when it was received by the server.
func NewMessage(
	sender User,
	toGroups []Group,
	toUsers []User,
	content []byte,
	createdAt time.Time,
) (*SentMessage, []ReceivedMessage, error) {
	// id := "" // TODO: generate uuid

	groupIDs := make([]string, 0, len(toGroups))
	for _, g := range toGroups {
		groupIDs = append(groupIDs, g.id)
	}

	userIDs := make([]string, 0, len(toUsers))
	for _, u := range toUsers {
		userIDs = append(userIDs, u.id)
	}

	// sentAt := time.Now()
	// sm := SentMessage{
	// 	id:        fmt.Sprintf("%s-%s", sender.ID(), id),
	// 	toGroups:  groupIDs,
	// 	toUsers:   userIDs,
	// 	content:   content,
	// 	createdAt: createdAt,
	// 	sentAt:    sentAt,
	// 	received:  make([]stuff, 0, 0),
	// 	seen:      make([]stuff, 0, 0),
	// }

	// var rms []ReceivedMessage
	// for _, g := range toGroups {

	// 	rm := ReceivedMessage{
	// 		id:          "",
	// 		in:          nil,
	// 		messageFrom: sender.id,
	// 		content:     content,
	// 		createdAt:   createdAt,
	// 		sentAt:      sentAt,
	// 	}
	// }

	return nil, nil, nil
}
