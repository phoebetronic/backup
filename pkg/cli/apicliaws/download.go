package apicliaws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/xh3b4sd/tracer"
)

func (a *AWS) Download(buc string, key string) ([]byte, error) {
	var siz int64
	{
		inp := &s3.HeadObjectInput{
			Bucket: aws.String(buc),
			Key:    aws.String(key),
		}

		out, err := a.S3.HeadObject(context.Background(), inp)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		siz = out.ContentLength
	}

	{
		fmt.Printf("fetching %s\n", a.siz(siz))
	}

	var wri *Writer
	{
		wri = &Writer{
			wri: manager.NewWriteAtBuffer([]byte{}),
			siz: siz,
		}
	}

	{
		inp := &s3.GetObjectInput{
			Bucket: aws.String(buc),
			Key:    aws.String(key),
		}

		_, err := manager.NewDownloader(a.S3).Download(context.Background(), wri, inp)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	{
		fmt.Printf("\n")
	}

	return wri.wri.Bytes(), nil
}
