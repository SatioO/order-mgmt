package queue

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/spf13/viper"
)

func Init() *sqs.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(viper.GetString("AWS_REGION")))

	if err != nil {
		log.Fatalln("failed to connect to aws", err)
	}

	return sqs.NewFromConfig(cfg)
}
