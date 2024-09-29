package entities

type MongoConnectionInfo struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
}

type MySQLConnectionInfo struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
}

type RedisConnectionInfo struct {
	Address  string
	Password string
	DB       int
}

type DatabaseConnectionInfo struct {
	MongoInfo *MongoConnectionInfo
	MysqlInfo *MySQLConnectionInfo
	RedisInfo *RedisConnectionInfo
}
