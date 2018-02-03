package main

import (
	"strings"
	"path/filepath"
	"os"
	"fmt"
	"os/exec"
)

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

func main()  {
	args := os.Args //获取用户输入的所有参数
	if args==nil || len(args)<1{
		fmt.Println("请输入路径?")
		return
	}
	fmt.Println(args[1])
	files ,_:=WalkDir(args[1],".proto");
	for _,file := range  files{
		rs:=[]rune(file)
		idx:=strings.LastIndex(file,"\\")
		//len:=len(rs)
		path:=string(rs[0:idx])
		//file =string(rs[idx+1:len])
		fmt.Println(path)
		fmt.Println(file)
		//c := exec.Command("protoc", "--go_out=plugins=micro:.", file)

		//c := exec.Command("cmd", "/C", "del", "D:\\a.txt")
		c := exec.Command("cmd", "/C","protoc","--proto_path="+path,"--go_out=plugins=micro:.",file)
		c.Dir=path;
		output, err := c.CombinedOutput()
		if err != nil {
			fmt.Println("Error:", err)
		}
		fmt.Printf("callEXE结果:\n%v", string(output))


	}
}
