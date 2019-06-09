package userService

import(
	"blog/models/user"
	"blog/util"
	mydb"blog/db/mysql"
)
func CreateNewUser(username string, loginId string, password string) error {
	var sqlConn = mydb.Conn()
	md5Password := util.EncryptionPassword(password)
	_, err := sqlConn.Execute("insert into users(username, password, loginId, id) values(?,?,?,?)", username, md5Password, loginId, util.GetRowId())
	if (err != nil) {
		return err
	}
	return nil
}
func UserLogin(loginId string, password string) (bool, string, error) {
	var userList []model.User
	password = util.EncryptionPassword(password)
	err := mydb.Conn().Table(&userList).Where("loginId",loginId).Where("password", password).Select()
	if err != nil {
		return false, "", err
	}
	// 没有查询到该用户
	if len(userList) == 0 {
		return false, "", nil
	}
	// 登陆成功
	return true, userList[0].Id, nil
}