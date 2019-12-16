/* Package config provider global config information */
package config

import (
	"encoding/json"
	"io/ioutil"
)

//config about database
type DB struct {
	User string `json:"user"`
	Host string `json:"host"`
	Port int `json:"port"`
	Password string `json:"password"`
	Database string `json:"database"`

}

// config about web server
type Web struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type Config struct {
	DB  DB  `json:"db"`
	Web Web `json:"web"`
}

//GlobalConfig store the config from config.json
var globalConfig *Config

//Get globalConfig
func Get() *Config {
	return globalConfig
}

//read config data from json file
func ReadConfig() error {
	data, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		return err
	}
	globalConfig = &Config{}
	err = json.Unmarshal(data, globalConfig)
	if err != nil {
		return err
	}
	return nil
}
