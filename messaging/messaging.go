package messaging

import (
	"context"
	"time"

	"github.com/jarri-abidi/gochat"
	"github.com/pkg/errors"
)

type Service interface {
	Send(context.Context, SendRequest) (*SendResponse, error)
	Receive(context.Context) []gochat.ReceivedMessage
}

// messaging.EventPublisher
type EventPublisher interface {
	PublishMessageCreatedEvent(context.Context, []gochat.ReceivedMessage) error
}

type EventConsumer interface {
	ConsumeMessageCreatedEvent(context.Context) ([]gochat.ReceivedMessage, error)
}

type inmemQueue struct {
	msgs map[string]chan interface{}
}

// func (q *inmemQueue) Publish(ctx context.Context, event interface{}) error {
// 	q.msgs <- event
// 	return nil
// }

// func (q *inmemQueue) Consume(ctx context.Context) (interface{}, error) {
// 	return <-q.msgs, nil
// }

const MessageCreatedEvent = "MESSAGE_CREATED_EVENT"

func (q *inmemQueue) ConsumeMessageCreatedEvent(ctx context.Context) (interface{}, error) {
	return <-q.msgs[MessageCreatedEvent], nil
}

func (q *inmemQueue) PublishMessageCreatedEvent(ctx context.Context, rms []gochat.ReceivedMessage) error {
	q.msgs[MessageCreatedEvent] <- rms
	return nil
}

type SendRequest struct {
	Sender     gochat.User
	Recipients []gochat.Contact
	Groups     []gochat.Group
	Content    []byte
	CreatedAt  time.Time
}

type SendResponse struct{}

func NewService() Service {
	return &service{}
}

type service struct {
	sentMessages     gochat.SentMessageRepository
	receivedMessages gochat.ReceivedMessageRepository
	publisher        EventPublisher
	consumer         EventConsumer
}

func (s *service) Send(ctx context.Context, req SendRequest) (*SendResponse, error) {
	sm, rms, err := gochat.NewMessage(req.Sender, req.Groups, req.Recipients, req.Content, req.CreatedAt)
	if err != nil {
		return nil, errors.Wrap(err, "could not create new message")
	}

	if _, err := s.sentMessages.Insert(ctx, *sm); err != nil {
		return nil, errors.Wrapf(err, "could not save message for sender", sm.SenderID())
	}

	for _, rm := range rms {
		if _, err := s.receivedMessages.Insert(ctx, rm); err != nil {
			// TODO: don't fail unless all failed
			return nil, errors.Wrapf(err, "could not save message for recipient %s", rm.ReceiverID())
		}
	}

	// TODO: publish some event/notification? the consumer c would check
	// each msg and find the server to which the recipient is connected.
	// c would then send the message to that server otherwise if recipient
	// is offline then perhaps send a push notification.
	s.publisher.PublishMessageCreatedEvent(ctx, rms)

	return &SendResponse{}, nil
}

func (s *service) Receive(ctx context.Context) []gochat.ReceivedMessage {
	// rms, _ := s.consumer.ConsumeMessageCreatedEvent(ctx)

	// for _, rm := range rms {
	// 	s.presenceService
	// }

	var cursor string
	// TODO: run the following in a loop until there are unread messages
	s.receivedMessages.FindAll(ctx, 10, cursor)
	return nil
}

// faisal, _ := gochat.NewUser("markhaur", "Faisal Nisar")
// jarri, _ := gochat.NewUser("jarri-abidi", "Jarri Abidi")
// sarim, _ := gochat.NewUser("sa41m", "sarim")

// fContact, _ := gochat.NewContact(faisal.ID(), faisal.FullName(), faisal.UserName())
// jContact, _ := gochat.NewContact(jarri.ID(), jarri.FullName(), jarri.UserName())
// sContact, _ := gochat.NewContact(sarim.ID(), sarim.FullName(), sarim.UserName())
// jarri.AddContacts(*fContact, *sContact)
// faisal.AddContacts(*jContact)
// sarim.AddContacts(*jContact)

// group, _ := gochat.NewGroup("gochat", *fContact, *jContact, *sContact)

// sm, rms, _ := gochat.NewMessage(*faisal, []gochat.Group{*group}, []gochat.User{*jarri}, []byte("message"), time.Now())

// fmt.Printf("sm: %+v\n", sm)
// fmt.Printf("rms: %+v\n", rms)
// TODO: fetch jarri-abidi messages using Repository
