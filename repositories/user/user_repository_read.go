package user

import (
	"mnc-backend/models/db"

	"github.com/gocraft/dbr"
)

func GetUserByEmail(sess *dbr.Session, email string) (*db.User, error) {
	var user *db.User

	query := sess.Select("*").
		From(db.User{}.TableName()).
		Where("email = ?", email)

	err := query.LoadOne(&user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
