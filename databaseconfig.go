package go4it

type DatabaseConfig struct {
	ConnectionName string `toml:"connName"`
	Engine         string `toml:"engine"`
	Host           string `toml:"host"`
	Port           string `toml:"port"`
	User           string `toml:"user"`
	Password       string `toml:"password"`
	DBName         string `toml:"dbname"`
}
