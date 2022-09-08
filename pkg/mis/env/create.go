package env

import "os"

func Create() Env {
	var e Env

	{
		e.AWS.Region = os.Getenv(AwsDefaultRegion)
		e.AWS.AccessKey = os.Getenv(AwsAccessKey)
		e.AWS.SecretKey = os.Getenv(AwsSecretKey)
	}

	{
		verify(e)
	}

	return e
}
