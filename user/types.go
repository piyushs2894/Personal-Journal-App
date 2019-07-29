package user

import "time"

type UserData struct {
	ID        int
	UserName  string
	Password  string
	Name      string
	Email     string
	Mobile    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
