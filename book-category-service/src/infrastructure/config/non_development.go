package config

import (
	"book-category-service/src/common/log"
	"context"
	"encoding/base64"
	"os"
	"strings"

	vault "github.com/hashicorp/vault/api"
	"github.com/sirupsen/logrus"
)

func setUpForNonDevelopment(appStatus string) *Config {
	defaultConf := vault.DefaultConfig()
	defaultConf.Address = os.Getenv("SYNOPSIS_BOOK_MANAGEMENT_CONFIG_ADDRESS")

	client, err := vault.NewClient(defaultConf)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "config.setUpForNonDevelopment", "section": "vault.NewClient"}).Fatal(err)
	}

	client.SetToken(os.Getenv("SYNOPSIS_BOOK_MANAGEMENT_CONFIG_TOKEN"))

	mountPath := "synopsis-book-management-secrets" + "-" + strings.ToLower(appStatus)

	bookCategoryServiceSecrets, err := client.KVv2(mountPath).Get(context.Background(), "book-category-service")
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "config.setUpForNonDevelopment", "section": "KVv2.Get"}).Fatal(err)
	}

	bookServiceSecrets, err := client.KVv2(mountPath).Get(context.Background(), "book-service")
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "config.setUpForNonDevelopment", "section": "KVv2.Get"}).Fatal(err)
	}

	jwtSecrets, err := client.KVv2(mountPath).Get(context.Background(), "jwt")
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "config.setUpForNonDevelopment", "section": "KVv2.Get"}).Fatal(err)
	}

	currentAppConf := new(currentApp)
	currentAppConf.RestfulAddress = bookCategoryServiceSecrets.Data["RESTFUL_ADDRESS"].(string)
	currentAppConf.RestfulAddress = bookCategoryServiceSecrets.Data["GRPC_ADDRESS"].(string)

	bookServiceConf := new(bookService)
	bookServiceConf.GrpcAddress = bookServiceSecrets.Data["GRPC_ADDRESS"].(string)

	postgresConf := new(postgres)
	postgresConf.Dsn = bookCategoryServiceSecrets.Data["POSTGRES_DSN"].(string)
	postgresConf.User = bookCategoryServiceSecrets.Data["POSTGRES_USER"].(string)
	postgresConf.Password = bookCategoryServiceSecrets.Data["POSTGRES_PASSWORD"].(string)

	jwtConf := new(jwt)

	jwtPrivateKey := jwtSecrets.Data["PRIVATE_KEY"].(string)
	base64Byte, err := base64.StdEncoding.DecodeString(jwtPrivateKey)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "config.setUpForNonDevelopment", "section": "base64.StdEncoding.DecodeString"}).Fatal(err)
	}
	jwtPrivateKey = string(base64Byte)

	jwtPublicKey := jwtSecrets.Data["Public_KEY"].(string)
	base64Byte, err = base64.StdEncoding.DecodeString(jwtPublicKey)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "config.setUpForNonDevelopment", "section": "base64.StdEncoding.DecodeString"}).Fatal(err)
	}
	jwtPublicKey = string(base64Byte)

	jwtConf.PrivateKey = loadRSAPrivateKey(jwtPrivateKey)
	jwtConf.PublicKey = loadRSAPublicKey(jwtPublicKey)

	return &Config{
		CurrentApp:  currentAppConf,
		BookService: bookServiceConf,
		Postgres:    postgresConf,
		Jwt:         jwtConf,
	}
}
