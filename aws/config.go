package aws

import (
	"commons/logger"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var _logger = logger.New()

func newConfig() aws.Config{
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		_logger.Fatalf("unable to load SDK config: %v", err)
	}

	return cfg
}