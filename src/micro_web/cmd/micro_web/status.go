package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/keybase/go-ps"

)

// ServerStatus represents the status of a server
type ServerStatus struct {
	ServiceProcs            []ps.Process
	ServiceName            []string
	ServerID             ServerID
}

func detectServerStatus() *ServerStatus {
	ss := &ServerStatus{}
	procs, err := ps.Processes()
	checkErrorOrQuit(err, "list processes failed")
	for _, proc := range procs {
		path, err := proc.Path()
		if err != nil {
			if pathErr, ok := err.(*os.PathError); ok {
				path = pathErr.Path
				if strings.HasSuffix(path, " (deleted)") {
					path = path[:len(path)-10]
				}
			} else {
				continue
			}
		}

		relpath, err := filepath.Rel(env.GetCmdDir(), path)
		if err != nil || strings.HasPrefix(relpath, "..") {
			continue
		}

		dir, _ := filepath.Split(relpath)

			if strings.HasSuffix(dir, string(filepath.Separator)) {
				dir = dir[:len(dir)-1]
			}
			serverName :=strings.Replace(relpath,BinaryExtension,"",-1)
			ss.ServiceProcs = append(ss.ServiceProcs, proc)
			ss.ServiceName= append(ss.ServiceName,serverName)
		}
	return ss
}

func status() {
	ss := detectServerStatus()
	showServerStatus(ss)
}

func showServerStatus(ss *ServerStatus) {
	for _,p :=range  ss.ServiceName{
		showMsg("%s",p)
	}

}
func (ss* ServerStatus)IsRun(servername string)(ret bool){
	for _,v:=range ss.ServiceName{
		if(v==servername){
			return  true
		}
	}
	return  false
}

