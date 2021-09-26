package dao

import (
	"github.com/skyhackvip/geek/sproject/internal/model"
)

func (d *Dao) CheckLogin(name, pass string) bool {
	u := &model.User{}
	err := d.db.QueryRow("select id from user where name=?", name).Scan(&u.Pass)
	if err != nil {
		return false
	}
	if u.Pass == pass {
		return true
	}
	return false
}
