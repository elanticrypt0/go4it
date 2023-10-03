package go4it

import "fmt"

type App struct {
	Config *AppConfig
	DB     DB
}

func LoadAppConfig(configFile string) *AppConfig {

	if configFile == "" {
		configFile = "appconfig"
	}
	configFile = configFile + ".toml"

	var appconfig AppConfig
	ReadAndParseToml(configFile, &appconfig)
	// set default app url
	appconfig.App_url = fmt.Sprintf("%s:%d", appconfig.App_server_host, appconfig.App_server_port)
	return &appconfig
}

func NewApp(configFile string) App {
	return App{
		Config: LoadAppConfig(configFile),
	}
}

func (a *App) Connect2Db(connName string) {
	Connect2DB(a, connName)
}
