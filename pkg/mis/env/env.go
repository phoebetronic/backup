package env

const (
	AwsDefaultRegion = "AWS_DEFAULT_REGION"
	AwsAccessKey     = "AWS_ACCESS_KEY"
	AwsSecretKey     = "AWS_SECRET_KEY"
	FtxApiKey        = "FTX_API_KEY"
	FtxApiSecret     = "FTX_API_SECRET"
)

type Env struct {
	AWS EnvAWS
	FTX EnvFTX
}

type EnvAWS struct {
	Region    string
	AccessKey string
	SecretKey string
}

type EnvFTX struct {
	ApiKey    string
	ApiSecret string
}
