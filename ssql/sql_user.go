package ssql

import (
	"main/model"
	"main/utils"
)

func CheckUserMobileAndPassword(mobile string, password string) (*model.User, error) {

	sqlStr := "select id, mobile, name, password from userinfo where mobile = ? and password = ?"
	row := utils.Db.QueryRow(sqlStr, mobile, password)
	user := &model.User{}
	err := row.Scan(&user.ID, &user.Mobile, &user.Name, &user.Password)
	if err != nil{
		return user,err
	}
	return user, nil
}

func AddUser(user *model.User) error {
	sqlStr := "insert into userinfo(mobile, name, password) values(?,?,?)"
	_, err := utils.Db.Exec(sqlStr, user.Mobile, user.Name, user.Password)
	if err != nil {
		return err
	}
	utils.Db.QueryRow("SELECT id FROM userinfo WHERE mobile=?", user.Mobile).Scan(&user.ID)
	return nil

}
