package prj

import (
	"github.com/17bixin/gobase/log"
	"os"
	"os/exec"
)

func Initprj(name, path string, servicetype int) {
	if os.IsExist(path + "/" + name) {
		log.Error(log.Fields{
			"err":           "target folder already exist",
			"target folder": path + "/" + name,
		})
		return
	}
	os.Mkdir(path+"/"+name, os.ModePerm)
	exec.Command("git", "clone ")
}
