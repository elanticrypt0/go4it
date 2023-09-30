package go4it

func LoadAppConfig(configFile string) *AppConfig {

	if configFile == "" {
		configFile = "appconfig"
	}
	configFile = configFile + ".toml"

	var appconfig AppConfig
	ReadOrParseToml(configFile, &appconfig)
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
