package User

import (
	"mnc-backend/formatter"
	"mnc-backend/http/utils"
	"mnc-backend/repositories/access_token"
	"mnc-backend/repositories/user"
	"net/http"

	"github.com/gocraft/dbr"
)

type Logout struct {
	DBConn *dbr.Connection
}

func (c Logout) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	email := r.Context().Value("Email").(string)

	sess := c.DBConn.NewSession(nil)

	err := utils.CheckAuthorization(r, sess)
	if err != nil {
		formatter.ERROR(w, http.StatusBadRequest, err)
		return
	}

	currentUser, err := user.GetUserByEmail(sess, email)
	if err != nil {
		formatter.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = access_token.UpdateUserAccessToken(sess, currentUser.Email, nil)
	if err != nil {
		formatter.ERROR(w, http.StatusBadRequest, err)
		return
	}

	return
}
