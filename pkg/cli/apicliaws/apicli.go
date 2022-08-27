package apicliaws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// AWS is a container for AWS SDK specific clients. Environment variables as
// stated below MUST be provided when using this client implementation.
//
//     AWS_ACCESS_KEY
//     AWS_SECRET_KEY
//     AWS_REGION
//
//     https://github.com/aws/aws-sdk-go-v2/blob/386724971857987a5d2a50f506f23df615765ac7/config/env_config.go
//
type AWS struct {
	S3 *s3.Client
}

func New() *AWS {
	var err error

	var c aws.Config
	{
		c, err = config.LoadDefaultConfig(context.Background())
		if err != nil {
			panic(err)
		}
	}

	return &AWS{
		S3: s3.NewFromConfig(c),
	}
}
