package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"strings"
	"time"
	"micro_web/consts"
	"github.com/pkg/errors"
)

func start(sid ServerID) {
	err := os.Chdir(env.MicroWebRoot)
	checkErrorOrQuit(err, "chdir to goworld directory failed")

	ss := detectServerStatus()
	if(ss.IsRun(sid.Name())){
		showMsgAndQuit("server is already running, can not start multiple servers")
	}


	StartService(sid, false)

}


func StartService(sid ServerID,  isRestore bool) {
	showMsg("start service %s ...", sid.Name())

	gameExePath := filepath.Join(env.GetCmdDir(), sid.Name()+BinaryExtension)

	cmd := exec.Command(gameExePath)
	err := runCmdUntilTag(cmd,  env.GetCmdDir()+"\\"+sid.Name()+".log", consts.START_SUCCESS, time.Second*10)
	checkErrorOrQuit(err, "start game failed, see game.log for error")
}

func runCmdUntilTag(cmd *exec.Cmd, logFile string, tag string, timeout time.Duration) (err error) {
	err = cmd.Start()
	if err != nil {
		return
	}

	timeoutTime := time.Now().Add(timeout)
	for time.Now().Before(timeoutTime) {
		time.Sleep(time.Millisecond * 200)
		if isTagInFile(logFile, tag) {
			cmd.Process.Release()
			return
		}
	}
	err = errors.Errorf("wait started tag timeout")
	return
}

func isTagInFile(filename string, tag string) bool {
	data, err := ioutil.ReadFile(filename)
	checkErrorOrQuit(err, "read file error")
	return strings.Contains(string(data), tag)
}
