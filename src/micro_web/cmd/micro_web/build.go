package main

import (
	"os"
	"os/exec"
)

func build(sid ServerID) {
	showMsg("building server %s ...", sid)

	buildServer(sid)
	//buildDispatcher()
	//buildGate()
}

func buildServer(sid ServerID) {
	serverPath :=sid.Path()
	showMsg("server directory is %s ...", serverPath)
	if !isdir(serverPath) {
		showMsgAndQuit("wrong server id: %s, using '\\' instead of '/'?", sid)
	}
	targetName:=env.GetCmdDir()+"\\"+sid.Name()+BinaryExtension
	showMsg("go build %s ...", sid)
	buildDirectory(serverPath,targetName)
}

func buildDispatcher() {
	showMsg("go build dispatcher ...")
	//buildDirectory(filepath.Join(env.MicroWebRoot, "components", "dispatcher"))
}

func buildGate() {
	showMsg("go build gate ...")
	//buildDirectory(filepath.Join(env.MicroWebRoot, "components", "gate"))
}

func buildDirectory(dir string,targetname string) {
	var err error
	var curdir string
	curdir, err = os.Getwd()
	checkErrorOrQuit(err, "")

	err = os.Chdir(dir)
	checkErrorOrQuit(err, "")

	defer os.Chdir(curdir)

	cmd := exec.Command("go", "build", "-o",targetname,".")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	err = cmd.Run()
	checkErrorOrQuit(err, "")
	return
}
