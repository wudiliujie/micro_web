package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/xiaonanln/goworld/engine/config"
	"io/ioutil"
	"fmt"
)

// Env represents environment variables
type Env struct {
	MicroWebRoot string
	BinRoot string
	Services []string
}

// GetDispatcherDir returns the path to the dispatcher
func (env *Env) GetCmdDir() string {
	return filepath.Join(env.MicroWebRoot, "cmd")
}
// GetDispatcherDir returns the path to the dispatcher
func (env *Env) GetDispatcherDir() string {
	return filepath.Join(env.MicroWebRoot, "components", "dispatcher")
}

// GetGateDir returns the path to the gate
func (env *Env) GetGateDir() string {
	return filepath.Join(env.MicroWebRoot, "components", "gate")
}

// GetDispatcherBinary returns the path to the dispatcher binary
func (env *Env) GetDispatcherBinary() string {
	return filepath.Join(env.GetDispatcherDir(), "dispatcher"+BinaryExtension)
}

// GetGateBinary returns the path to the gate binary
func (env *Env) GetGateBinary() string {
	return filepath.Join(env.GetGateDir(), "gate"+BinaryExtension)
}

var env Env

func getGoSearchPaths() []string {
	var paths []string
	goroot := os.Getenv("GOROOT")
	if goroot != "" {
		paths = append(paths, goroot)
	}

	gopath := os.Getenv("GOPATH")
	for _, p := range strings.Split(gopath, string(os.PathListSeparator)) {
		if p != "" {
			paths = append(paths, p)
		}
	}
	return paths
}

func detectGoWorldPath() {
	searchPaths := getGoSearchPaths()
	showMsg("go search paths: %s", strings.Join(searchPaths, string(os.PathListSeparator)))
	for _, sp := range searchPaths {
		goworldPath := filepath.Join(sp, "src", "micro_web")
		if isdir(goworldPath) {
			env.MicroWebRoot = goworldPath
			env.BinRoot=filepath.Join(sp, "bin")
			break
		}
	}
	if env.MicroWebRoot == "" {
		showMsgAndQuit("micro_web directory is not detected")
	}

	showMsg("microweb directory found: %s", env.MicroWebRoot)
	//查找左右的服务和api
	dir, err := ioutil.ReadDir(env.MicroWebRoot)
	if(err!=nil){
		checkErrorOrQuit(err,"搜索路径错误")
	}
	for _,file:=range dir{
		if file.IsDir() {
			str:=strings.Split(file.Name(),"_");
			if len(str)>=2{
				if str[0]=="api" ||  str[0]=="svr"{
					env.Services= append(env.Services, file.Name())
				}
			}
			fmt.Println( file.Name());
		}
	}
	configFile := filepath.Join(env.MicroWebRoot, "goworld.ini")
	config.SetConfigFile(configFile)
	//config.Get()
}
