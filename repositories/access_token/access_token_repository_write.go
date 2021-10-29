package access_token

import (
	"mnc-backend/models/db"
	"mnc-backend/models/requests"

	"github.com/gocraft/dbr"
)

func InsertAccessToken(sess *dbr.Session, r requests.InsertAccessTokenRequest) error {
	userAccessToken := db.AccessToken{
		Email:       r.Email,
		AccessToken: &r.AccessToken,
	}

	columns := []string{
		"email",
		"access_token",
	}

	_, err := sess.InsertInto(db.AccessToken{}.TableName()).
		Columns(columns...).
		Record(userAccessToken).
		Exec()
	if err != nil {
		return err
	}

	return err
}

func UpdateUserAccessToken(sess *dbr.Session, currentEmail string, token *string) error {
	data := make(map[string]interface{})

	data["access_token"] = token

	_, err := sess.Update(db.AccessToken{}.TableName()).
		SetMap(data).
		Where("email = ?", currentEmail).
		Exec()
	if err != nil {
		return err
	}

	return nil
}
