package main

import "github.com/mkevac/yaegitesting/user"

type UserDB struct {
	users []*user.User
}

func NewDBWithRandomUsers(count uint) *UserDB {
	userID := uint32(1)
	db := &UserDB{
		users: make([]*user.User, 0, count),
	}
	for i := uint(0); i < count; i++ {
		randomUser := user.NewRandomUser()
		randomUser.UserID = userID
		userID++
		db.users = append(db.users, randomUser)
	}
	return db
}
