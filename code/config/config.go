package config

var Config2_file = `

package config

import (
	"encoding/json"
	"github.com/17bixin/gobase/log"
	"io/ioutil"
)

type Config struct {
	HttpPort        int

}

var (
	DefaultConfig Config = Config{
		HttpPort:        8000,
	}
)

func Init(configfilepath string) error {
	data, err := ioutil.ReadFile(configfilepath)
	if err != nil {
		log.Fatal(log.Fields{"error": err, "app": "config file read "})
		return err
	}

	err = json.Unmarshal([]byte(data), &DefaultConfig)
	if err != nil {
		log.Fatal(log.Fields{"error": err, "app": "config file parse "})
		return err
	}
	log.Info(log.Fields{"app": "config file", "config": DefaultConfig})
	return nil
}
`
