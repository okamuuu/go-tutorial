package xorm

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type User struct {
	Id      int64
	Name    string
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

var engine *xorm.Engine

func init() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:@/go_tutorial?charset=utf8")
	if err != nil {
		panic(err)
	}
	engine.ShowSQL(true)

	err = engine.CreateTables(&User{})
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = engine.Exec("truncate user")
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = engine.Insert(&User{Name: "xlw"})
	if err != nil {
		fmt.Println(err)
		return
	}
}

func FindUsers() ([]User, error) {
	users := make([]User, 0)
	err := engine.Find(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}
