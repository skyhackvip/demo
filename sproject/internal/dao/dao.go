package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/skyhackvip/geek/sproject/configs"
)

type Dao struct {
	db *sql.DB
}

func New(c *configs.Config) *Dao {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", c.User, c.Pass, c.Host, c.Port, c.Name))
	if err != nil {
		panic(err)
	}
	return &Dao{
		db: db,
	}
}

func (d *Dao) Close() {
	if d.db != nil {
		d.db.Close()
	}
}
