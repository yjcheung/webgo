package model

import (
	"fmt"
	"webgo/utils"
)

//User结构体
type User struct {
	ID       int
	Username string
	Password string
	Phone    string
}

//AddUser 添加User的方法一
func (user *User) AddUser() error {
	//sql语句
	sqlStr := "insert into users(username,password,phone) values(?,?,?)"
	//预编译
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译出现异常：", err)
		return err
	}

	_, err2 := inStmt.Exec("admin", "123456", "13344556677")
	if err2 != nil {
		fmt.Println("执行出现异常：", err2)
		return err2
	}

	return nil

}
