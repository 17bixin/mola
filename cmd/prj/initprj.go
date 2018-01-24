package prj

import (
	"fmt"
	"github.com/17bixin/gobase/log"
	"os"
	"os/exec"
)

func Initprj(name, path string, servicetype int) {
	targetdir := path + "/" + name
	if os.IsExist(targetdir) {
		log.Error(log.Fields{
			"err":           "target folder already exist",
			"target folder": targetdir,
		})
		return
	}
	os.Mkdir(targetdir, os.ModePerm)
	os.Mkdir(targetdir+"/src", os.ModePerm)
	os.Mkdir(targetdir+"/bin", os.ModePerm)
	os.Setenv("GOPATH", targetdir)
	goget := exec.Command("go", "get", "github.com/17bixin/gobase")
	out, err := goget.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}
