package person

type Config struct {
	DatabaseType string
	DatabaseName string
	DatabaseUrl  string
}

func NewConfig(dbType string, dbName string, dbUrl string) *Config {
	return &Config{
		DatabaseType: dbType,
		DatabaseName: dbName,
		DatabaseUrl:  dbUrl,
	}
}

const (
	Mongo    = "MONGODB"
	MySQL    = "MYSQL"
	PostGres = "POSTGRES"
)

type IConfig interface {
	GetDBType() string
	GetDBName() string
	GetDBUrl() string
}

func (config *Config) GetDBType() string {
	return config.DatabaseType
}
func (config *Config) GetDBName() string {
	return config.DatabaseName
}
func (config *Config) GetDBUrl() string {
	return config.DatabaseUrl
}
