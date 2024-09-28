package utils

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"os"
)

type Functions struct {
}

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
	mongoInfo *MongoConnectionInfo
	mysqlInfo *MySQLConnectionInfo
	redisInfo *RedisConnectionInfo
}

func NewFunctions() *Functions {
	return &Functions{}
}

func (f *Functions) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (f *Functions) LoadEnvData() (*DatabaseConnectionInfo, error) {
	MysqlDBUser := os.Getenv("MYSQL_USER")
	MysqlDBPassword := os.Getenv("MYSQL_PASSWORD")
	MysqlDBName := os.Getenv("MYSQL_DB_NAME")
	MysqlDBHost := os.Getenv("MYSQL_HOST")
	MysqlDBPort := os.Getenv("MYSQL_PORT")

	MongoDBUser := os.Getenv("MONGO_USER")
	MongoDBPassword := os.Getenv("MONGO_PASSWORD")
	MongoDBName := os.Getenv("MONGO_DB_NAME")
	MongoDBHost := os.Getenv("MONGO_HOST")
	MongoDBPort := os.Getenv("MONGO_PORT")

	RedisHost := os.Getenv("REDIS_HOST")
	RedisPassword := os.Getenv("REDIS_PASSWORD") // Corrección de nombre
	RedisDB := os.Getenv("REDIS_DB")

	// Validar que las variables de entorno necesarias están cargadas
	if MysqlDBUser == "" || MysqlDBPassword == "" || MysqlDBName == "" || MysqlDBHost == "" || MysqlDBPort == "" ||
		MongoDBUser == "" || MongoDBPassword == "" || MongoDBName == "" || MongoDBHost == "" || MongoDBPort == "" ||
		RedisHost == "" || RedisPassword == "" || RedisDB == "" {
		return nil, errors.New("some required environment variables are missing")
	}

	mysql := &MySQLConnectionInfo{
		DBUser:     MysqlDBUser,
		DBPassword: MysqlDBPassword,
		DBName:     MysqlDBName,
		DBHost:     MysqlDBHost,
		DBPort:     MysqlDBPort,
	}

	mongo := &MongoConnectionInfo{
		DBUser:     MongoDBUser,
		DBPassword: MongoDBPassword,
		DBName:     MongoDBName,
		DBHost:     MongoDBHost,
		DBPort:     MongoDBPort,
	}

	redis := &RedisConnectionInfo{
		Address:  RedisHost,
		Password: RedisPassword,
		DB:       RedisDB,
	}

	return &DatabaseConnectionInfo{
		mongoInfo: mongo,
		mysqlInfo: mysql,
		redisInfo: redis,
	}, nil
}
