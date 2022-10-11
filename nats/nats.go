package nats

import (
	"github.com/jarri-abidi/gochat/messaging"
	"github.com/jarri-abidi/gochat/notification"
)

// nats.NewMessagingPublisher
// messaging.Service needs a messaging.EventPublisher
func NewMessagingPublisher() messaging.EventPublisher {
	return nil
}

// nats.NewNotificationPublisher
// notification.Service needs a notification.EventPublisher
func NewNotificationPublisher() notification.EventPublisher {
	return nil
}
