package db

import "time"

type TransferHistory struct {
	ID         int       `db:"id" json:"id"`
	SenderID   int       `db:"sender_id" json:"sender_id"`
	ReceiverID int       `db:"receiver_id" json:"receiver_id"`
	Amount     int       `db:"amount" json:"amount"`
	Message    *string   `db:"message" json:"message,omitempty"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
}

func (c TransferHistory) TableName() string {
	return "transfer_histories"
}
