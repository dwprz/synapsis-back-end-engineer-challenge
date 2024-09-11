package config

import (
	"user-service/src/common/log"
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

	userServiceSecrets, err := client.KVv2(mountPath).Get(context.Background(), "user-service")
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "config.setUpForNonDevelopment", "section": "KVv2.Get"}).Fatal(err)
	}

	jwtSecrets, err := client.KVv2(mountPath).Get(context.Background(), "jwt")
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "config.setUpForNonDevelopment", "section": "KVv2.Get"}).Fatal(err)
	}

	currentAppConf := new(currentApp)
	currentAppConf.RestfulAddress = userServiceSecrets.Data["RESTFUL_ADDRESS"].(string)

	postgresConf := new(postgres)
	postgresConf.Dsn = userServiceSecrets.Data["POSTGRES_DSN"].(string)
	postgresConf.User = userServiceSecrets.Data["POSTGRES_USER"].(string)
	postgresConf.Password = userServiceSecrets.Data["POSTGRES_PASSWORD"].(string)

	redisConf := new(redis)
	redisConf.AddrNode1 = userServiceSecrets.Data["REDIS_ADDR_NODE_1"].(string)
	redisConf.AddrNode2 = userServiceSecrets.Data["REDIS_ADDR_NODE_2"].(string)
	redisConf.AddrNode3 = userServiceSecrets.Data["REDIS_ADDR_NODE_3"].(string)
	redisConf.AddrNode4 = userServiceSecrets.Data["REDIS_ADDR_NODE_4"].(string)
	redisConf.AddrNode5 = userServiceSecrets.Data["REDIS_ADDR_NODE_5"].(string)
	redisConf.AddrNode6 = userServiceSecrets.Data["REDIS_ADDR_NODE_6"].(string)
	redisConf.Password = userServiceSecrets.Data["REDIS_PASSWORD"].(string)

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
		CurrentApp: currentAppConf,
		Postgres:   postgresConf,
		Redis:      redisConf,
		Jwt:        jwtConf,
	}
}
