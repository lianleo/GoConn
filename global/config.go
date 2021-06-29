package global

import (
	"errors"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/lianleo/GoCommon/db/cache"
	"github.com/lianleo/GoCommon/db/mongo"
	"github.com/lianleo/GoCommon/log"
)

var configPath string
var (
	ERROR_CONFIG_PARSE = errors.New("load config error")
	Config             TOMLFile
)

//TOMLFile 配置
type TOMLFile struct {
	WebAPP      WebAPP       `toml:"WebAPP"`
	MasterMDB   mongo.Config `toml:"MasterMDB"`
	LogConfig   log.Config   `toml:"Log"`
	CommonCache cache.Config `toml:"CommonCache"`
}

//WebAPP WebAPP
type WebAPP struct {
	Title     string `toml:"title"`
	Env       string `toml:"env"`
	Protocol  string `toml:"protocol"`
	Domain    string `toml:"domain"`
	App       string `toml:"app"`
	Port      string `toml:"port"`
	Expires   int    `toml:"expires"`
	StaticDir string `toml:"static_dir"`
}

//Install This func Will init the app config before server start.
func Install(cp string) error {
	if cp == "" {
		configPath = os.Getenv("WEB_SERVER_CONFIG")
	} else {
		configPath = cp
	}
	log.Infof("configPath = %s", configPath)
	if err := configInit(configPath); err != nil {
		return ERROR_CONFIG_PARSE
	}
	log.Init(Config.LogConfig)

	//初始化数据库
	if err := mongo.Init(Config.MasterMDB); err != nil {
		fmt.Println("init failed: ", err)
		return fmt.Errorf("init failed: %v", err)
	}

	return nil
}

func configInit(loc string) error {
	if loc == "" {
		return ERROR_CONFIG_PARSE
	}
	if _, err := toml.DecodeFile(loc, &Config); err != nil {
		log.Fatal("configInit error", err)
		return ERROR_CONFIG_PARSE
	}
	fmt.Println("config loaded.")
	return nil
}
