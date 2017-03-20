package utils

import (
	"io/ioutil"
	"crypto/x509"
	"log"
	"flag"
	"github.com/golang/glog"
	"strings"
	myconfig "github.com/larspensjo/config"
)

var (
	configFile = flag.String("configfile", "config/server.toml", "General configuration file")
)

//topic list
var topic = make(map[string]string)


func LoadCA(caFile string) *x509.CertPool {
	pool := x509.NewCertPool()

	if ca, e := ioutil.ReadFile(caFile); e != nil {
		log.Fatal("ReadFile: ", e)
	} else {
		pool.AppendCertsFromPEM(ca)
	}
	return pool
}

func ReadConfig(config string) string {

	configs := strings.Split(config, ".")

	cfg, err := myconfig.ReadDefault(*configFile)

	if err != nil {
		glog.Fatalf("Fail to find", *configFile, err)
	}

	//Initialized topic from the configuration
	if cfg.HasSection(configs[0]) {
		section, err := cfg.SectionOptions(configs[0])
		if err == nil {
			for _, v := range section {
				options, err := cfg.String(configs[0], v)
				if err == nil {
					topic[v] = options
				}
			}
		}
	}
	//Initialized topic from the configuration END

	return topic[configs[1]]

}

