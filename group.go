package gochat

import (
	"context"
	"strings"

	"github.com/pkg/errors"
)

type Group struct {
	id           string
	name         string
	participants []Contact
}

var ErrEmptyGroupName = errors.New("group name cannot be empty")

func NewGroup(name string, participants ...Contact) (*Group, error) {
	if strings.TrimSpace(name) == "" {
		return nil, ErrEmptyGroupName
	}

	return &Group{
		id:           "", // TODO: generate uuid
		name:         name,
		participants: participants,
	}, nil
}

func (g Group) ID() string { return g.id }

type GroupRepository interface {
	Insert(context.Context, Group) (*Group, error)
	FindAll(context.Context) ([]Group, error)
	FindByID(ctx context.Context, id string) (*Group, error)
	Update(context.Context, Group) error
}
