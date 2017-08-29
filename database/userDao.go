package database

import (
	user "ss-server/models"
)

func QueryUser(account string, password string) (mUser user.User, err error) {
	row := Db.QueryRow("SELECT * FROM vpn_user WHERE account = ? AND password = ?", account, password)
	err = row.Scan(&mUser.ID, &mUser.Name, &mUser.Account, &mUser.Password)
	return
}
