package gochat

import (
	"strings"

	"github.com/pkg/errors"
)

type User struct {
	id       string
	userName string
	fullName string
	groups   []Group
	contacts []Contact
}

var (
	ErrEmptyID       = errors.New("id cannot be empty")
	ErrEmptyUserName = errors.New("user name cannot be empty")
	ErrEmptyFullName = errors.New("full name cannot be empty")
)

func NewUser(userName, fullName string, contacts ...Contact) (*User, error) {
	if strings.TrimSpace(userName) == "" {
		return nil, ErrEmptyUserName
	}

	if strings.TrimSpace(fullName) == "" {
		return nil, ErrEmptyFullName
	}

	return &User{
		id:       "", // TODO: generate uuid
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

type Contact struct {
	id       string
	fullName string
	userName string
}

func NewContact(id, fullName, userName string) (*Contact, error) {
	if strings.TrimSpace(id) == "" {
		return nil, ErrEmptyID
	}

	if strings.TrimSpace(userName) == "" {
		return nil, ErrEmptyUserName
	}

	if strings.TrimSpace(fullName) == "" {
		return nil, ErrEmptyFullName
	}

	return &Contact{id: id, fullName: fullName, userName: userName}, nil
}

func (c Contact) ID() string       { return c.id }
func (c Contact) UserName() string { return c.userName }
func (c Contact) FullName() string { return c.fullName }
