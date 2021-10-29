package user

import (
	"mnc-backend/models/db"
	"mnc-backend/models/requests"

	"github.com/gocraft/dbr"
)

func CreateUser(sess *dbr.Session, r requests.RegisterRequest) (int, error) {
	tx, err := sess.Begin()
	if err != nil {
		return 0, err
	}

	user := db.User{
		FullName: r.FullName,
		Email:    r.Email,
		Password: r.Password,
	}

	columns := []string{
		"full_name",
		"email",
		"password",
	}

	var id int

	err = sess.InsertInto(db.User{}.TableName()).
		Columns(columns...).
		Record(user).
		Returning("id").
		Load(id)

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	tx.Commit()
	return id, nil
}
