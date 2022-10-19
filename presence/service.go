package presence

import "context"

type Service interface {
	FindUser(context.Context, FindUserRequest) (*FindUserResponse, error)
}

type FindUserRequest struct{}

type FindUserResponse struct {
	IsOnline bool
	Server   string
}
