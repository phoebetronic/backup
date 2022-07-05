package apicliaws

import (
	"context"
	"io"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/xh3b4sd/tracer"
)

func (a *AWS) Upload(buc string, key string, fil *os.File) error {
	var err error

	var inf os.FileInfo
	{
		inf, err = fil.Stat()
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var rea io.Reader
	{
		rea = &Reader{
			fil: fil,
			siz: inf.Size(),
		}
	}

	{
		inp := &s3.PutObjectInput{
			Bucket: aws.String(buc),
			Key:    aws.String(key),
			Body:   rea,
		}

		_, err := manager.NewUploader(a.S3, par).Upload(context.Background(), inp)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}

func par(u *manager.Uploader) {
	u.PartSize = 5 * 1024 * 1024
	u.LeavePartsOnError = true
}
