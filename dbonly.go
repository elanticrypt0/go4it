package go4it

import (
	"github.com/k23dev/go4it/interact"
	"log"
)

type DBOnly struct {
	Config DBOnlyConfig
	Log    *log.Logger
	DB     DB
}

type DBOnlyConfig struct {
	Connection map[string]DatabaseConfig `toml:"db"`
}

func LoadDBOnlyConfig(configFile string) DBOnlyConfig {

	if configFile == "" {
		configFile = "config"
	}
	configFile = configFile + ".toml"

	var config DBOnlyConfig
	interact.ReadAndParseToml(configFile, &config)
	return config
}

func NewAppDBOnly(configFilePath string) DBOnly {
	return DBOnly{
		Config: LoadDBOnlyConfig(configFilePath),
		Log:    newLog(),
	}
}

func (dbo *DBOnly) Connect2Db(connName string) {
	Connect2DBOnly(dbo, connName)
}
