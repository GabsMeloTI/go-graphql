package infra

import (
	"github.com/joho/godotenv"
	"os"
)

type ConfigEnv struct {
	ServerName         string
	ServerPort         string
	Environment        string
	DBHost             string
	DBPort             string
	DBUser             string
	DBPassword         string
	DBDatabase         string
	DBSSLMode          string
	DBDriver           string
	SignatureToken     string
	AwsAccessKeyID     string
	AwsSecretAccessKey string
	AwsRegion          string
	AwsQueueUrl        string
}

func NewConfig() ConfigEnv {
	if os.Getenv("ENVIRONMENT") == "" {
		if err := godotenv.Load(".env"); err != nil {
			panic("Error loading env file")
		}
	}

	return ConfigEnv{
		ServerName:         os.Getenv("SERVER_NAME"),
		ServerPort:         os.Getenv("SERVER_PORT"),
		Environment:        os.Getenv("ENVIRONMENT"),
		DBHost:             os.Getenv("DB_HOST"),
		DBPort:             os.Getenv("DB_PORT"),
		DBUser:             os.Getenv("DB_USER"),
		DBPassword:         os.Getenv("DB_PASSWORD"),
		DBDatabase:         os.Getenv("DB_DATABASE"),
		DBSSLMode:          os.Getenv("DB_SSL_MODE"),
		DBDriver:           os.Getenv("DB_DRIVER"),
		SignatureToken:     os.Getenv("SIGNATURE_STRING"),
		AwsAccessKeyID:     os.Getenv("AWS_ACCESS_KEY"),
		AwsSecretAccessKey: os.Getenv("AWS_SECRET_KEY"),
		AwsRegion:          os.Getenv("AWS_REGION"),
		AwsQueueUrl:        os.Getenv("AWS_QUEUE_URL"),
	}
}
