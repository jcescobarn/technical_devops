package utils

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strconv"

	"koronet_web_app/entities"
)

type Functions struct {
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

func (f *Functions) LoadEnvData() (*entities.DatabaseConnectionInfo, error) {
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
	RedisPassword := os.Getenv("REDIS_PASSWORD")
	RedisDB := os.Getenv("REDIS_DB")

	if MysqlDBUser == "" || MysqlDBPassword == "" || MysqlDBName == "" || MysqlDBHost == "" || MysqlDBPort == "" ||
		MongoDBUser == "" || MongoDBPassword == "" || MongoDBName == "" || MongoDBHost == "" || MongoDBPort == "" ||
		RedisHost == "" || RedisPassword == "" || RedisDB == "" {
		return nil, errors.New("some required environment variables are missing")
	}

	mysql := &entities.MySQLConnectionInfo{
		DBUser:     MysqlDBUser,
		DBPassword: MysqlDBPassword,
		DBName:     MysqlDBName,
		DBHost:     MysqlDBHost,
		DBPort:     MysqlDBPort,
	}

	mongo := &entities.MongoConnectionInfo{
		DBUser:     MongoDBUser,
		DBPassword: MongoDBPassword,
		DBName:     MongoDBName,
		DBHost:     MongoDBHost,
		DBPort:     MongoDBPort,
	}

	redisDBToInt, err := strconv.Atoi(RedisDB)
	if err != nil {
		return nil, err
	}

	redis := &entities.RedisConnectionInfo{
		Address:  RedisHost,
		Password: RedisPassword,
		DB:       redisDBToInt,
	}

	return &entities.DatabaseConnectionInfo{
		MongoInfo: mongo,
		MysqlInfo: mysql,
		RedisInfo: redis,
	}, nil
}
