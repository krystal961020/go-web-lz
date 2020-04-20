package conf

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

var EsClient *elastic.Client

const host = "http://127.0.0.1:9200/"

type Conf struct {
	Host       string `yaml:"esHost"`
	Port       string `yaml:"Port"`
	DBUserName string `yaml:"DBUserName"`
	DBPassword string `yaml:"DBPassword"`
	DBIp       string `yaml:"DBIp"`
	DBPort     int    `yaml:"DBPort"`
	DBName     string `yaml:"DBName"`
}

func init() {
	conf := &Conf{}
	if f, err1 := ioutil.ReadFile("./conf.yaml"); err1 != nil {
		err1 := yaml.Unmarshal(f, conf)
		if err1 != nil {
			executeError(err1)
		}
	}
	fmt.Printf("conf %s", host)

	errLog := log.New(os.Stdout, "APP", log.LstdFlags)
	var err error
	EsClient, err := elastic.NewClient(elastic.SetErrorLog(errLog), elastic.SetURL(host))
	executeError(err)
	info, code, err := EsClient.Ping(host).Do(context.Background())
	executeError(err)
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
	esVersion, err := EsClient.ElasticsearchVersion(host)
	executeError(err)
	fmt.Printf("Elasticsearch version %s\n", esVersion)
}

func executeError(err error) {
	if err != nil {
		//panic(err)
	}
}
