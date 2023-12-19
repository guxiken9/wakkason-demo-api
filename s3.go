package main

import (
	"bytes"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
)

const AWS_REGION = "ap-northeast-1"
const S3_BUCKET_NAME = "wakkason-demo-bucket-team-e"

func NewS3Client() (*s3.S3, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(AWS_REGION),
	})
	if err != nil {
		return nil, err
	}

	client := s3.New(sess)
	return client, nil
}

type Result struct {
	Key          string
	PreSignedURL string
}

func PutImage(image []byte) (*Result, error) {

	c, err := NewS3Client()
	if err != nil {
		return nil, err
	}
	key := uuid.NewString()
	if err != nil {
		return nil, err
	}

	size := len(image)
	fileType := http.DetectContentType(image)

	var result Result
	_, err = c.PutObject(&s3.PutObjectInput{
		Bucket:        aws.String(S3_BUCKET_NAME),
		Key:           aws.String(key),
		Body:          aws.ReadSeekCloser(bytes.NewReader(image)),
		ContentLength: aws.Int64(int64(size)),
		ContentType:   aws.String(fileType),
	})
	if err != nil {
		return nil, err
	}
	result.Key = S3_BUCKET_NAME + "/" + key

	req, _ := c.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(S3_BUCKET_NAME),
		Key:    aws.String(key),
	})
	url, err := req.Presign(24 * time.Hour)
	if err != nil {
		return nil, err
	}
	result.PreSignedURL = url
	return &result, nil
}
