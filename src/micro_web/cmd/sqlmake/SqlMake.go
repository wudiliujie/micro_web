package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	strings "strings"
	"os"
	"fmt"
	"io"
	"micro_web/core"
	"github.com/pkg/errors"


	"time"
	"path/filepath"
	"apiserveer/DB"
)
type Result struct {
	Models[] ModelInfo `xml:"model"`
	Sqls []SqlInfo `xml:"sql"`
}
type SqlInfo struct {
	Name      string    `xml:"n,attr"`
	Sql       string       `xml:"s,attr"`
	Desc         string       `xml:"d,attr"`
	Out    string    `xml:"out,attr"`
	Array  bool      `xml:"array,attr"`
	Type  string      `xml:"t,attr"`
	Param    []SqlProp    `xml:"p"`
}
type SqlProp struct {
	Name string `xml:"n,attr"`
	Type string `xml:"t,attr"`
	Desc string `xml:"d,attr"`
}
type ModelInfo struct {
	Name      string    `xml:"n,attr"`
	Desc         string       `xml:"d,attr"`
	Prop    []ModelProp    `xml:"p"`
}
type ModelProp struct {
	Name string `xml:"n,attr"`
	Type string `xml:"t,attr"`
	Desc string `xml:"d,attr"`
	Maping string `xml:"m,attr"`
}

type Writer struct {
	Content string
	PrevCount int
}
func (writer *Writer)AddLine(msg string){
	if(strings.HasSuffix(msg,"}")){
		writer.PrevCount--;

	}
	if(writer.PrevCount>0){
		writer.Content+=strings.Repeat("\t",writer.PrevCount);
	}

	writer.Content+=msg;
	writer.Content+="\r\n";
	if(strings.HasSuffix(msg,"{")){
		writer.PrevCount++;
	}
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
func  test()  {
	/*	users:= DB.GetAllUserAccount(1,10)
	for _,v := range *users{
		fmt.Println(v.UserName)
	}
	user:=DB.GetUserInfo(1);
	fmt.Println(user.UserName)
	count:=DB.GetUserCount();
	fmt.Println(count)*/
	logonline:=DB.GetLogOnlineInfo(time.Date(2018,1,30,0,0,0,0,time.UTC))
	for _,v := range *logonline{
		fmt.Println(v)
	}

}
//获取指定目录及所有子目录下的所有文件，可以匹配后缀过滤。
func WalkDir(dirPth, suffix string) (files []string, err error) {
	files = make([]string, 0, 30)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		//if err != nil { //忽略错误
		// return err
		//}

		if fi.IsDir() { // 忽略目录
			return nil
		}

		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, filename)
		}

		return nil
	})

	return files, err
}
func main() {
	args := os.Args //获取用户输入的所有参数
	if args==nil || len(args)<1{
		fmt.Println("请输入路径?")
		return
	}
	fmt.Println(args[1])
	files ,_:=WalkDir(args[1],"SQL.xml");
	for _,file := range  files{
		rs:=[]rune(file)
		idx:=strings.LastIndex(file,"\\")
		//len:=len(rs)
		path:=string(rs[0:idx])
		//file =string(rs[idx+1:len])
		fmt.Println(path)
		fmt.Println(file)
		MakeGo(file,path+"\\data.go");
	}
}
func  MakeGo(file string,target string)  {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	var result Result
	err = xml.Unmarshal(content, &result)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(result)
	var writer Writer;
	writer.AddLine("package db");
	writer.AddLine("import \"micro_web/core\"")
	writer.AddLine("import \"micro_web/db\"")
	writer.AddLine("import \"time\"")
	for _, model  := range result.Models {
		writer.AddLine("type "+model.Name+" struct {")
		for _,p :=range  model.Prop{
			writer.AddLine(p.Name+" "+p.Type);
		}
		writer.AddLine("}")
	}
	for _,sql := range result.Sqls{
		//查找对应的out
		for _,mm := range result.Models{
			if mm.Name==sql.Out {
				MakeFunc(&writer,&sql,&mm);
			}
			if(sql.Type=="E"){
				MakeExec(&writer,&sql)
			}
		}
	}
	fmt.Println(writer.Content);
	fmt.Println("aaa");
	//保存文件
	Save(target,writer.Content);
	
}

func MakeFunc(writer *Writer,info * SqlInfo,model* ModelInfo){

	param :=""
	inparam:=""
	//写入参数
	if len(info.Param)>0 {
		for _,v := range info.Param{
			param+=v.Name+","
			inparam+=v.Name+" "+ v.Type+","
		}
		param= strings.TrimSuffix(param,",");
		inparam= strings.TrimSuffix(inparam,",");
	}


	if(info.Array){
		writer.AddLine("func "+info.Name+"("+inparam+")(ret *[]"+info.Out+")   {")
		writer.AddLine("var result []"+info.Out)
	}else {
		writer.AddLine("func "+info.Name+"("+inparam+")(ret "+info.Out+")   {")
	}


	writer.AddLine("var m "+info.Out)
	writer.AddLine("db := db.G_dbmanage.GetDB();")
	if param!=""{
		writer.AddLine("rows,err := db.Query(\""+info.Sql+"\","+param+")")
	}else {
		writer.AddLine("rows,err := db.Query(\""+info.Sql+"\")")
	}

	writer.AddLine("core.Check(err);")
	writer.AddLine("defer rows.Close();")
	if info.Type=="Q"{
		writer.AddLine("var refs []interface{};")
		//获取顺序scan
		fields:=GetFieldOrder(&info.Sql, model);
		for _,v :=range  fields{
			writer.AddLine("refs = append(refs, &m."+v+")")
		}
	}
	writer.AddLine("for rows.Next() {")

	if info.Type=="Q"{
		writer.AddLine("err := rows.Scan(refs...);")
	}else if info.Type=="S"{
		writer.AddLine("err := rows.Scan(&m);")
	}




	writer.AddLine("core.Check(err);")
	if(info.Array){
		writer.AddLine("result= append(result, m)")
	}
	writer.AddLine("}")



	if(info.Array){
		writer.AddLine("return  &result;")
	}else {
		writer.AddLine("return  m;")
	}
	writer.AddLine("}")

}
func MakeExec(writer *Writer,info * SqlInfo){

	param :=""
	inparam:=""
	//写入参数
	if len(info.Param)>0 {
		for _,v := range info.Param{
			param+=v.Name+","
			inparam+=v.Name+" "+ v.Type+","
		}
		param= strings.TrimSuffix(param,",");
		inparam= strings.TrimSuffix(inparam,",");
	}
	writer.AddLine("func "+info.Name+"("+inparam+")(lastid int64,rows int64)   {")
	writer.AddLine("db := db.G_dbmanage.GetDB();")
	if param!=""{
		writer.AddLine("ret,err := db.Exec(\""+info.Sql+"\","+param+")")
	}else {
		writer.AddLine("ret,err := db.Exec(\""+info.Sql+"\")")
	}
	writer.AddLine("core.Check(err)")
	writer.AddLine("lastid ,_=ret.LastInsertId()");
	writer.AddLine("rows ,_=ret.RowsAffected()");


	writer.AddLine("return  lastid,rows")
	writer.AddLine("}")
}


func GetFieldOrder( sql *string, info *ModelInfo)(ret []string){
	idx:= strings.Index(*sql,"select")
	if(idx<0){
		errors.New("不是查询语句");
	}
	end:=strings.Index(*sql,"from")
	if(end<=0){
		errors.New("缺少from")
	}
	var result []string;
	//field:=strings.sub
	rs := []rune(*sql)
	fields:= strings.Split(string(rs[idx+6:end]),",")
	for i,v := range fields{
		field:=strings.ToLower(strings.Trim(v," "));
		if strings.Index(field," as ")>-1{
			str:=strings.Split(field," as ")
			field=strings.Trim(str[1]," ")
		}
		fields[i]= field;
	}


	for _,v := range fields{
		field:=strings.Trim(v," ");
		for _,p := range  info.Prop{
			if field== strings.ToLower(p.Name) || field==strings.ToLower(p.Maping) {
				result= append(result,p.Name);
			}
		}
	}
	return  result;

}



func Save(filename string, data string){
	f, err1 := os.Create(filename) //创建文件
	core.Check(err1)
	n, err1 := io.WriteString(f, data) //写入文件(字符串)
	core.Check(err1)
	fmt.Println("写入 %d 个字节n", n)
}