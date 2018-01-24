package prj

import (
	"github.com/17bixin/gobase/log"
	"github.com/17bixin/mola/code/apps/front"
	"github.com/17bixin/mola/code/config"
	"github.com/17bixin/mola/code/server"
	"io/ioutil"
	"os"
	"os/exec"
)

func Initprj(name, path string, servicetype int) error {
	targetdir := path + "/" + name
	_, err := os.Stat(targetdir)
	if err == nil {
		// no such file or dir
		log.Error(log.Fields{
			"err":           "target folder already exist",
			"target folder": targetdir,
		})
		return err
	}
	os.MkdirAll(targetdir, os.ModePerm)
	os.Mkdir(targetdir+"/bin", os.ModePerm)

	os.MkdirAll(targetdir+"/src/apps/front", os.ModePerm)
	ioutil.WriteFile(targetdir+"/src/apps/front/main.go", []byte(front.Front2_file), os.ModePerm)
	os.MkdirAll(targetdir+"/src/common/config", os.ModePerm)
	ioutil.WriteFile(targetdir+"/src/common/config/config.go", []byte(config.Config2_file), os.ModePerm)
	os.MkdirAll(targetdir+"/src/common/server", os.ModePerm)
	ioutil.WriteFile(targetdir+"/src/common/server/server.go", []byte(server.Server1_file), os.ModePerm)

	os.Setenv("GOPATH", targetdir)
	goget := exec.Command("go", "get", "github.com/17bixin/gobase")
	goget.Stdout = os.Stdout
	exr := goget.Run()
	if exr != nil {
		return exr
	}
	return nil
}
