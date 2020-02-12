package load

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/naoina/toml"
)

// Conf 全域的設定變數
var Conf *Config

// LoadConfig 讀取設定檔
func LoadConfig() *Config {
	env := os.Getenv("PROJECT_ENV")

	root, getWdErr := os.Getwd()
	if getWdErr != nil {
		log.Println("getWdErr :", getWdErr)
		return nil
	}

	// 讀取預設的設定檔
	defaultPath := root + "/config/default.toml"
	defaultToml, readErr := ioutil.ReadFile(defaultPath)
	if readErr != nil {
		log.Println("readErr :", readErr)
		return nil
	}
	unmarshalErr := toml.Unmarshal(defaultToml, &Conf)
	if unmarshalErr != nil {
		log.Println("unmarshalErr :", unmarshalErr)
		return nil
	}

	// 讀取環境的設定檔以覆蓋
	configFile := root + "/config/" + env + ".toml"
	_, statErr := os.Stat(configFile)
	if statErr == nil {
		tomlData, readFileErr := ioutil.ReadFile(configFile)
		if readFileErr != nil {
			log.Println("readFileErr :", readFileErr)
			return nil
		}

		unmarshalErr := toml.Unmarshal(tomlData, &Conf)
		if unmarshalErr != nil {
			log.Println("unmarshalErr :", unmarshalErr)
			return nil
		}
	} else if !os.IsNotExist(statErr) {
		log.Println("讀取設定檔錯誤 Err:", statErr)
		return nil
	}

	return Conf
}
