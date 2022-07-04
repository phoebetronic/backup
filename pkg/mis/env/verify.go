package env

import "fmt"

func verify(e Env) {
	{
		if e.AWS.Region == "" {
			panic(fmt.Sprintf("${%s} must not be empty", AwsDefaultRegion))
		}
		if e.AWS.AccessKey == "" {
			panic(fmt.Sprintf("${%s} must not be empty", AwsAccessKey))
		}
		if e.AWS.SecretKey == "" {
			panic(fmt.Sprintf("${%s} must not be empty", AwsSecretKey))
		}
	}

	{
		if e.FTX.ApiKey == "" {
			panic(fmt.Sprintf("${%s} must not be empty", FtxApiKey))
		}
		if e.FTX.ApiSecret == "" {
			panic(fmt.Sprintf("${%s} must not be empty", FtxApiSecret))
		}
	}
}
