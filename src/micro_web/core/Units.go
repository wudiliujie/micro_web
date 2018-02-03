package core

import (
	"os"
	"strings"
	"os/exec"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetCurrentPath() string {
	s, err := exec.LookPath(os.Args[0])
	Check(err)
	i := strings.LastIndex(s, "\\")
	path := string(s[0 : i+1])
	return path
}