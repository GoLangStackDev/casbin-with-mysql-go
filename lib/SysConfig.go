package lib

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Mysql struct {
	DNS string `yaml:"dns"`
}
type SysConfig struct {
	Mysql Mysql
}

func newSysConfig() *SysConfig {
	return &SysConfig{}
}

func InitConfig() *SysConfig {
	config := newSysConfig()
	if b:= LoadConfigFile();b!=nil {
		if err:=yaml.Unmarshal(b,config);err!=nil {
			log.Fatal(err)
		}
	}
	return config
}
//读取配置文件
func LoadConfigFile() []byte {
	file := GetWorkDir() + "/application.yaml"
	b,err := ioutil.ReadFile(file)
	if err!=nil {
		log.Println(err)
		return nil
	}
	return b
}
//获取当前工作目录
func GetWorkDir() string {
	wd,_:=os.Getwd()
	return strings.Replace(wd,"\\","/",-1)
}