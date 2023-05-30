package viperconf

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/viper"
)

var conf Configuration

//unexportedgo
const CONFIGFILEPATH = "./config"

//ViperEcho prints "Hello from ViperEcho"
func ViperEcho() string {

	out, err := json.Marshal(conf)
	if err != nil {
		panic(err)
	}

	return string(out)

}

//InitConfig InitConfig
func InitConfig() Configuration {

	initViper()
	initConfiguration()
	return conf
}

//InitConfigWithfile InitConfig for injection
func InitConfigWithfile(configfile string) Configuration {

	initViperConfigfile(configfile)
	initConfiguration()
	return conf
}

func initViper() {
	initViperConfigfile(CONFIGFILEPATH)
}

func initViperConfigfile(configfile string) {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")

	// ชื่อ config file
	//viper.SetConfigName(configfile)

	// ระบุ path ของ config file
	//viper.AddConfigPath(".") ระบุจาก package main
	viper.AddConfigPath(configfile)

	// อ่าน value จาก ENV variable & Configmap
	viper.AutomaticEnv()

	// แปลง _ underscore ใน env เป็น . dot notation ใน viper เพื่อรองรับ การรับค่าจาก ENV VAR
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// อ่าน config
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

}

func initConfiguration() {

	if err := viper.Unmarshal(&conf); err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

}
