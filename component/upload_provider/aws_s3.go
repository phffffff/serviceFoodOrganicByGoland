package uploadProvider

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"go_service_food_organic/common"
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

func (provider *s3Provider) SaveFileUploaded(c context.Context, data []byte, dst string) (*common.Image, error) {
	fileBytes := bytes.NewReader(data)
	fileType := http.DetectContentType(data)

	svc := s3.New(provider.session)

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

	img := &common.Image{
		Url:       fmt.Sprintf("%s/%s", provider.domain, dst),
		CloudName: "s3",
	}

	return img, nil
}

//_, err := svc.HeadObject(&s3.HeadObjectInput{
//Bucket: aws.String(provider.bucketName),
//Key:    aws.String(dst),
//})
//if err == nil {
//return nil, err
//}
