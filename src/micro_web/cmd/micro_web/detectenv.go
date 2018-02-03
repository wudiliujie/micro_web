package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/xiaonanln/goworld/engine/config"
)

// Env represents environment variables
type Env struct {
	MicroWebRoot string
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
			break
		}
	}
	if env.MicroWebRoot == "" {
		showMsgAndQuit("goworld directory is not detected")
	}

	showMsg("microweb directory found: %s", env.MicroWebRoot)
	configFile := filepath.Join(env.MicroWebRoot, "goworld.ini")
	config.SetConfigFile(configFile)
	//config.Get()
}
