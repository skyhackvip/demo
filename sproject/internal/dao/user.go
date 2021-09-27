package dao

import (
	"database/sql"
	"github.com/skyhackvip/geek/sproject/internal/model"
)

func (d *Dao) CheckLogin(name, pass string) (bool, error) {
	u := &model.User{}
	err := d.db.QueryRow("select id from user where name=?", name).Scan(&u.Pass)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		} else {
			return false, err
		}
	}
	if u.Pass == pass {
		return true, nil
	}
	return false, nil
}
