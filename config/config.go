package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"github.com/JIeeiroSst/go-app/repositories/mongo"
)

type WebConfig struct {
	URL string 	`envconfig:"WEB_URL"`
	MongoConfig mongo.Config `envconfig:"WEB_MONGO"`
}

var Config WebConfig

func init(){
	envconfig.Process("",&Config)
	fmt.Println(Config)
}