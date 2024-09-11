package config

import (
	"book-service/src/common/log"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func setUpForDevelopment() *Config {
	err := os.Chdir(os.Getenv("SYNOPSIS_BOOK_SERVICE_WORKSPACE"))
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "config.setUpForDevelopment", "section": "os.Chdir"}).Fatal(err)
	}

	viper := viper.New()
	viper.SetConfigFile(".env")
	viper.AddConfigPath(".")

	err = viper.ReadInConfig()
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "config.setUpForDevelopment", "section": "viper.ReadInConfig"}).Fatal(err)
	}

	currentAppConf := new(currentApp)
	currentAppConf.RestfulAddress = viper.GetString("CURRENT_APP_RESTFUL_ADDRESS")
	currentAppConf.GrpcAddress = viper.GetString("CURRENT_APP_GRPC_ADDRESS")

	bookCategoryServiceConf := new(bookCategoryService)
	bookCategoryServiceConf.GrpcAddress = viper.GetString("GRPC_ADDRESS_BOOK_CATEGORY_SERVICE ")

	postgresConf := new(postgres)
	postgresConf.Dsn = viper.GetString("POSTGRES_DSN")
	postgresConf.User = viper.GetString("POSTGRES_USER")
	postgresConf.Password = viper.GetString("POSTGRES_PASSWORD")

	jwtConf := new(jwt)
	jwtConf.PrivateKey = loadRSAPrivateKey(viper.GetString("JWT_PRIVATE_KEY"))
	jwtConf.PublicKey = loadRSAPublicKey(viper.GetString("JWT_PUBLIC_KEY"))

	return &Config{
		CurrentApp:          currentAppConf,
		BookCategoryService: bookCategoryServiceConf,
		Postgres:            postgresConf,
		Jwt:                 jwtConf,
	}
}
