package uploadProvider

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"go_service_food_organic/module/upload/image/model"
	"log"
	"net/http"
)

type s3Provider struct {
	bucketName string
	region     string
	apiKey     string
	secret     string
	domain     string
	session    *session.Session
}

func NewS3Provider(bucketName string, region string, apiKey string, secret string, domain string) *s3Provider {
	provider := &s3Provider{
		bucketName: bucketName,
		region:     region,
		apiKey:     apiKey,
		secret:     secret,
		domain:     domain,
	}

	s3Session, err := session.NewSession(&aws.Config{
		Region: aws.String(provider.region),
		Credentials: credentials.NewStaticCredentials(
			provider.apiKey,
			provider.secret,
			"",
		),
	})
	if err != nil {
		log.Fatalln(err)
	}

	provider.session = s3Session
	return provider
}

func (provider *s3Provider) SaveFileUploaded(c context.Context, data []byte, dst string) (*imageModel.Image, error) {
	fileBytes := bytes.NewReader(data)
	fileType := http.DetectContentType(data)

	svc := s3.New(provider.session)

	flag, _ := provider.isObjectExists(dst, svc)
	if flag {
		return nil, imageModel.ErrorFileExists()
	}

	_, err := svc.PutObject(&s3.PutObjectInput{
		ACL:         aws.String("private"),
		Body:        fileBytes,
		Bucket:      aws.String(provider.bucketName),
		ContentType: aws.String(fileType),
		Key:         aws.String(dst),
	})
	if err != nil {
		return nil, err
	}

	img := &imageModel.Image{
		Url:       fmt.Sprintf("%s/%s", provider.domain, dst),
		CloudName: "s3",
	}

	return img, nil
}

func (provider *s3Provider) isObjectExists(dst string, svc *s3.S3) (bool, error) {

	// List all objects in the bucket
	resp, err := svc.ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String(provider.bucketName),
	})

	if err != nil {
		return false, err
	}

	for _, obj := range resp.Contents {
		if aws.StringValue(obj.Key) == dst {
			return true, nil
		}
	}
	return false, nil
}
