package app

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/spf13/viper"
)

var configData = []byte(`
go_version: 1.9.1
version: 1.0.0

redis:
  host: 127.0.0.1
  port: 6379
  password: 646689
  db: 0
  idle_timeout: 120
  max_idle_conns: 100
  max_open_conns: 1000
  initial_open_conns: 10
`)

const (
	pass = "\u2713"
	fail = "\u2717"
)

func TestInitConfig(t *testing.T) {
	t.Log("start test read config...")

	filename := "test_config.yml"
	ioutil.WriteFile(filename, configData, 0664)

	viper.SetConfigName("test_config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		t.Fatal("\tinit viper error: ", err, "\033[0;31m", fail, "\033[0m")
	} else {
		t.Log("\tinit viper success", "\033[0;32m", pass, "\033[0m")
	}

	if viper.GetString("version") != "1.0.0" {
		t.Error("\tread one stage config information error", "\033[0;31m", fail, "\033[0m")
	} else {
		t.Log("\tread one stage config information success", "\033[0;32m", pass, "\033[0m")
	}

	if viper.GetInt("redis.port") != 6379 {
		t.Error("\tread two stage config information error", "\033[0;31m", fail, "\033[0m")
	} else {
		t.Log("\tread two stage config information success", "\033[0;32m", pass, "\033[0m")
	}

	os.Remove(filename)
}
