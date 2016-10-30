package xorm

import (
	"fmt"
	"testing"

	myxorm "go-tutorial/xorm"
)

func TestFind(t *testing.T) {
	users, err := myxorm.FindUsers()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(users)
}
