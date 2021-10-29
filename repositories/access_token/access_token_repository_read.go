package access_token

import (
	"mnc-backend/models/db"

	"github.com/gocraft/dbr"
)

func GetUserAccessToken(sess *dbr.Session, accessToken string) (*db.AccessToken, error) {
	var currentToken *db.AccessToken

	query := sess.Select("*").
		From(db.AccessToken{}.TableName()).
		Where("access_token = ?", accessToken)

	err := query.LoadOne(&currentToken)
	if err != nil {
		return nil, err
	}

	return currentToken, nil
}
