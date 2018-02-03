package db
import "micro_web/core"
import "micro_web/db"
import "time"
type UserInfo struct {
	Id int
	UserName string
	UserPass string
}
type GameInfo struct {
	Id int
	UserName string
	UserPass string
}
type LogOnlineInfo struct {
	Id int
	AppId int
	LogTime time.Time
	Online int
}
func GetAllUserAccount(pageno int,pagesize int)(ret *[]UserInfo)   {
	var result []UserInfo
	var m UserInfo
	db := db.G_dbmanage.GetDB();
	rows,err := db.Query("select id,username,userpass from user_account limit ?,?",pageno,pagesize)
	core.Check(err);
	defer rows.Close();
	var refs []interface{};
	refs = append(refs, &m.Id)
	refs = append(refs, &m.UserName)
	refs = append(refs, &m.UserPass)
	for rows.Next() {
		err := rows.Scan(refs...);
		core.Check(err);
		result= append(result, m)
	}
	return  &result;
}
func GetUserInfo(id int)(ret UserInfo)   {
	var m UserInfo
	db := db.G_dbmanage.GetDB();
	rows,err := db.Query("select id,username,userpass as UserPass from user_account where id =?",id)
	core.Check(err);
	defer rows.Close();
	var refs []interface{};
	refs = append(refs, &m.Id)
	refs = append(refs, &m.UserName)
	refs = append(refs, &m.UserPass)
	for rows.Next() {
		err := rows.Scan(refs...);
		core.Check(err);
	}
	return  m;
}
func GetLogOnlineInfo(date time.Time)(ret *[]LogOnlineInfo)   {
	var result []LogOnlineInfo
	var m LogOnlineInfo
	db := db.G_dbmanage.GetDB();
	rows,err := db.Query("select appid,logtime,online from log_online where datediff(d,logtime,?)=0",date)
	core.Check(err);
	defer rows.Close();
	var refs []interface{};
	refs = append(refs, &m.AppId)
	refs = append(refs, &m.LogTime)
	refs = append(refs, &m.Online)
	for rows.Next() {
		err := rows.Scan(refs...);
		core.Check(err);
		result= append(result, m)
	}
	return  &result;
}
