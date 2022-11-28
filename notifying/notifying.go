package notifying

import "context"

type Service interface {
	PushNotification(context.Context, PushNotificationRequest) error
}

type PushNotificationRequest struct {
	Header  []byte
	Content []byte
	UserID string
}

type EventPublisher interface{}
