package config

import (
	"applyUpLoadFile/utils"
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"os"
	"strings"
)

var (
	Cfg               = new(Config)
	DefaultConfigFile = "config-dev.toml"
)

type Config struct {
	Log       Log
	Databases map[string]Databases
	UpLoad    Upload
	Server    Server
	Email     Email
}

type Server struct {
	Host string `toml:"host" json:"host"`
	Port string `toml:"port" json:"port"`
}

type Databases struct {
	Name     string `toml:"name"`
	Type     string `toml:"type"`
	Url      string `toml:"url"`
	User     string `toml:"user"`
	PassWord string `toml:"password"`
	Mode     string `toml:"mode"`
}

type Log struct {
	Level        string  `toml:"level"       json:"level"`
	Formatter    string  `toml:"formatter"   json:"formatter"`
	OutFile      string  `toml:"out_file"    json:"out_file"`
	ErrFile      string  `toml:"err_file"    json:"err_file"`
	Release      float64 `toml:"release"    json:"release"`
	Mode         string  `toml:"mode"    json:"mode"`
	LogPath      string  `toml:"log_path"    json:"log_path"`
	LogName      string  `toml:"log_name"    json:"log_name"`
	MaxAge       int     `toml:"max_age"    json:"max_age"`
	RotationTime int     `toml:"rotation_time"    json:"rotation_time"`
}

type Upload struct {
	Url     string `toml:"url"`
	MaxFile int64  `toml:"max_file"`
}

type Email struct {
	Recipient    []string `toml:"recipient"`
	SmtpPassword string   `toml:"smtp_password"`
	SmtpUsername string   `toml:"smtp_username"`
	IamUserName  string   `toml:"iam_user_name"`
	Host         string   `toml:"host"`
}

func init() {

	if evn, err := ioutil.ReadFile("../environment"); err != nil {
		fmt.Println("获取环境文件失败：../environment,defualt \"dev\" env(dev,pre)")
	} else {
		DefaultConfigFile = strings.Replace("config-"+string(evn)+".toml", "\n", "", -1)
	}

	if err := InitConfig(DefaultConfigFile); err != nil {
		panic(err)
	}

	utils.CreateDateDir(Cfg.UpLoad.Url)
}

func InitConfig(configFile string) error {

	if configFile == "" {
		configFile = DefaultConfigFile
	}
	if _, err := os.Stat(configFile); err != nil {
		return errors.New("config file err:" + err.Error())
	} else {
		configBytes, er := ioutil.ReadFile(configFile)
		if er != nil {
			return errors.New("config load err:" + er.Error())
		}
		_, err = toml.Decode(string(configBytes), &Cfg)
		if err != nil {
			return errors.New("config decode err:" + err.Error())
		}
	}
	return nil
}

func GetConfig() *Config {
	return Cfg
}

func GetLogConfig() *Log {
	return &Cfg.Log
}
