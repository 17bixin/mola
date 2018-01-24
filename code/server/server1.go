package server

var Server1_file = `

package server

import (
	"common/config"
	"github.com/17bixin/gobase/httpserver"
	"net/http"
)

var (
	s = httpserver.New()
)

func Run() {
	s.SetPort(config.GWConfig.HttpPort)
	s.Run()
}

func Stop() {
	s.Stop()
}
func HandleFunc(path, method string, f func(http.ResponseWriter, *http.Request)) {
	s.Mux.Methods(method).Path(path).HandlerFunc(f)
}
`
