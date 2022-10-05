package gochat

import (
	"errors"
	"strings"
)

type User struct {
	id       string
	userName string
	fullName string
	groups   []Group
	contacts []Contact
}

var (
	ErrEmptyUserName = errors.New("username cannot be empty")
	ErrEmptyFullName = errors.New("fullname cannot be empty")
)

func NewUser(userName, fullName string, contacts ...Contact) (*User, error) {
	if strings.TrimSpace(userName) == "" {
		return nil, ErrEmptyUserName
	}

	if strings.TrimSpace(fullName) == "" {
		return nil, ErrEmptyFullName
	}

	return &User{
		id:       "", // generate uuid
		userName: userName,
		fullName: fullName,
		groups:   make([]Group, 0),
		contacts: contacts,
	}, nil
}

func (u User) ID() string          { return u.id }
func (u User) UserName() string    { return u.userName }
func (u User) FullName() string    { return u.fullName }
func (u User) Groups() []Group     { return append([]Group{}, u.groups...) }
func (u User) Contacts() []Contact { return append([]Contact{}, u.contacts...) }
