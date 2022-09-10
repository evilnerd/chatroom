package data

import (
	"sync"
	"time"
)

var (
	users sync.Map
)

type User struct {
	Name   string    `json:"name"`
	Joined time.Time `json:"joined"`
	Exists bool      `json:"-"`
}

func NewUser(name string) User {
	u := User{
		Name:   name,
		Joined: time.Now(),
		Exists: true,
	}
	users.Store(name, u)
	return u
}

func FindUser(name string) User {
	u, exists := users.Load(name)
	if !exists {
		return User{
			Exists: false,
		}
	}
	return u.(User)
}
