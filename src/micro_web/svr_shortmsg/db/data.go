package db
import "micro_web/core"
import "micro_web/db"
import "time"
type RegisterShortMsgInfo struct {
	Phone int64
	Code int
	SendTime time.Time
	Status int
}
func GetRegisterShortMsgInfo(phone int64)(ret RegisterShortMsgInfo)   {
	var m RegisterShortMsgInfo
	db := db.G_dbmanage.GetDB();
	rows,err := db.Query("select phone,code,sendtime,status from tab_register_shortmsg where id =?",phone)
	core.Check(err);
	defer rows.Close();
	var refs []interface{};
	refs = append(refs, &m.Phone)
	refs = append(refs, &m.Code)
	refs = append(refs, &m.SendTime)
	refs = append(refs, &m.Status)
	refs = append(refs, &m.SendTime)
	refs = append(refs, &m.Status)
	for rows.Next() {
		err := rows.Scan(refs...);
		core.Check(err);
	}
	return  m;
}
func InsertRegisterShortMsgInfo(Phone int64,Code int,SendTime time.Time,Status int)(lastid int64,rows int64)   {
	db := db.G_dbmanage.GetDB();
	ret,err := db.Exec("insert into tab_register_shortmsg (phone,code,sendtime,status) values(?,?,?,?) ",Phone,Code,SendTime,Status)
	core.Check(err)
	lastid ,_=ret.LastInsertId()
	rows ,_=ret.RowsAffected()
	return  lastid,rows
}
func UpdateRegisterShortMsgInfo(Code int,SendTime time.Time,Status int,Phone int64)(lastid int64,rows int64)   {
	db := db.G_dbmanage.GetDB();
	ret,err := db.Exec("update tab_register_shortmsg  set code=?,sendtime=?,status=? where phone=?  ",Code,SendTime,Status,Phone)
	core.Check(err)
	lastid ,_=ret.LastInsertId()
	rows ,_=ret.RowsAffected()
	return  lastid,rows
}
