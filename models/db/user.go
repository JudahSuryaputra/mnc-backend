package db

import "time"

type User struct {
	ID        int        `db:"id" json:"id"`
	FullName  string     `db:"full_name" json:"full_name"`
	Email     string     `db:"email" json:"email"`
	Password  string     `db:"password" json:"-"`
	UserType  string     `db:"user_type" json:"user_type"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"-"`
}

func (c User) TableName() string {
	return "users"
}
