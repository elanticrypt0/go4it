package go4it

import "gorm.io/gorm"

type App struct {
	Config    *AppConfig
	DB        []DBActive
	DBPrimary *gorm.DB
}

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

func (a *App) SetPrimaryDB(index uint8) {
	a.DBPrimary = a.DB[index].Conn
}

func (a *App) GetPrimaryDB() *gorm.DB {
	return a.DBPrimary
}
