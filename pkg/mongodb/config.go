package mongodb

// Конфигурация mongodb
type MongoDBConfig struct {
	Uri    string `mapstructure:"URI"`
	DbName string `mapstructure:"DBNAME"`
}
