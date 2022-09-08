package env

const (
	AwsDefaultRegion = "AWS_DEFAULT_REGION"
	AwsAccessKey     = "AWS_ACCESS_KEY"
	AwsSecretKey     = "AWS_SECRET_KEY"
)

type Env struct {
	AWS EnvAWS
}

type EnvAWS struct {
	Region    string
	AccessKey string
	SecretKey string
}
