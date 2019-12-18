package config

import (
	"favorite-jobs/types"
	"favorite-jobs/utils"
	"github.com/BurntSushi/toml"
)

var Config *types.Config

func LoadConfig() {
	if Config != nil {
		return
	}
	Config = &types.Config{}
	fileName := utils.GetConfigPath("./config/database.toml")
	if _, err := toml.DecodeFile(fileName, Config); err != nil {
		panic(err)
	}
}
