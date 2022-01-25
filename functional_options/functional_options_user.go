package main

import "fmt"

type User struct {
	Name     string
	Address  string
	Password string
	Username string
}

type UserOption func(*User)

func NewUser(opts ...UserOption) *User {
	usr := &User{}
	for _, opt := range opts {
		opt(usr)
	}
	return usr
}

func WithName(name string) UserOption {
	return func(u *User) {
		u.Name = name
	}
}

func WithAddress(address string) UserOption {
	return func(u *User) {
		u.Address = address
	}
}

func WithPassword(pass string) UserOption {
	return func(u *User) {
		u.Password = pass
	}
}

func WithUsername(username string) UserOption {
	return func(u *User) {
		u.Username = username
	}
}

func main() {
	user := NewUser(
		WithName("Lindsay"),
		WithAddress("1303 E Mission"),
		WithPassword("password"),
		WithUsername("LJam"),
	)
	fmt.Println(user)
}
