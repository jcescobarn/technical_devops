package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"koronet_web_app/config"
	"koronet_web_app/handlers"
	"koronet_web_app/repositories"
	"koronet_web_app/routes"
	"koronet_web_app/utils"
)

func main() {

	var utilsFunctions *utils.Functions = utils.NewFunctions()
	configData, err := utilsFunctions.LoadEnvData()

	var configRedis *config.RedisConfig = config.NewRedisConfig(configData.RedisInfo.Address, configData.RedisInfo.Password, configData.RedisInfo.DB)

	var configMysql *config.MySQLConfig = config.NewMysqlConfig(configData.MysqlInfo.DBUser, configData.MysqlInfo.DBPassword, configData.MysqlInfo.DBName, configData.MysqlInfo.DBHost, configData.MysqlInfo.DBPort)
	var configMongo *config.MongoDBConfig = config.NewMongoConfig(configData.MongoInfo.DBUser, configData.MongoInfo.DBPassword, configData.MongoInfo.DBName, configData.MongoInfo.DBHost, configData.MongoInfo.DBPort)

	mysqlConnection, err := configMysql.Connect()
	if err != nil {
		log.Fatalf("Mysql connection error")
	}

	mongoConnection, err := configMongo.Connect()
	if err != nil {
		log.Fatalf("MongoDB connection error")
	}

	redisConnection, err := configRedis.Connect()
	if err != nil {
		log.Fatalf("Redis connection error")
	}

	var userRepository *repositories.UserRepository = repositories.NewUserRepository(mysqlConnection)
	userRepository.EnsureDatabaseExists()
	var postRepository *repositories.PostRepository = repositories.NewPostRepository(mongoConnection, "koronet", "post")
	err = postRepository.EnsureCollectionExists()
	if err != nil {
		log.Fatalf("problem")
	}
	var sessionRepository *repositories.SessionRepository = repositories.NewSessionRepository(redisConnection)
	var userHandler *handlers.UserHandler = handlers.NewUserHandler(userRepository, utilsFunctions)
	var postHandler *handlers.PostHandler = handlers.NewPostHandler(postRepository)
	var sessionHandler *handlers.SessionHandler = handlers.NewSessionHandler(sessionRepository)
	var mainHandler *handlers.MainHandler = handlers.NewMainHandler()
	var userRoutes *routes.UserRoutes = routes.NewUserRoutes(userHandler)
	var postRoutes *routes.PostRoutes = routes.NewPostRoutes(postHandler)
	var sessionRoutes *routes.SessionRoutes = routes.NewSessionRoutes(sessionHandler)
	var mainRoutes *routes.GeneralRoutes = routes.NewGeneralRoutes(mainHandler)

	var webApp *gin.Engine = gin.Default()

	userRoutes.GetRoutes(webApp)
	postRoutes.GetRoutes(webApp)
	sessionRoutes.GetRoutes(webApp)
	mainRoutes.GetRoutes(webApp)

	webApp.Run()
}
