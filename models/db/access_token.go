package db

import "time"

type AccessToken struct {
	ID          int       `db:"id" json:"id"`
	Email       string    `db:"email" json:"email"`
	AccessToken *string   `db:"access_token" json:"access_token,omitempty"`
	CreatedAt   time.Time `db:"created_at" json:"-"`
}

func (c AccessToken) TableName() string {
	return "access_tokens"
}
