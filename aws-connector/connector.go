package connector

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/smithy-go"
)

const BUCKET_NAME string = "booking-bucket-2024"

type S3PersistentAgent struct {
	sdkClient *s3.Client
}

func NewS3PersistentAgent(ctx context.Context) (*S3PersistentAgent, error) {
	config, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("cannot load config from file ~/.aws/credentials")
	}
	sdkClient := s3.NewFromConfig(config)
	agent := &S3PersistentAgent{sdkClient: sdkClient}
	return agent, nil
}

func (agent *S3PersistentAgent) BucketExists(ctx context.Context) (bool, error) {
	_, err := agent.sdkClient.HeadBucket(ctx, &s3.HeadBucketInput{
		Bucket: aws.String(BUCKET_NAME),
	})
	exists := true
	if err != nil {
		var apiError smithy.APIError
		if errors.As(err, &apiError) {
			switch apiError.ErrorCode() {
			case "NotFound":
				exists = false
				err = fmt.Errorf("bucket: %s does not exist", BUCKET_NAME)
			default:
				err = fmt.Errorf("unexpected error while retrieving bucket info")
			}
		}
	}
	return exists, err
}

func (agent *S3PersistentAgent) DirExists(dirname string) (bool, error) {
	result, err := agent.sdkClient.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(BUCKET_NAME),
	})

	if err != nil {
		return false, fmt.Errorf("could not list objects from bucket due to: %v", err)
	}
	if !strings.HasSuffix("/", dirname) {
		dirname = dirname + "/"
	}
	for _, element := range result.Contents {
		if *element.Key == dirname {
			return true, nil
		}
	}
	return false, nil
}

func (agent *S3PersistentAgent) FileExists(fileName string) (bool, error) {
	_, err := agent.sdkClient.HeadObject(context.TODO(), &s3.HeadObjectInput{
		Bucket: aws.String(BUCKET_NAME),
		Key:    aws.String(fileName),
	})
	if err != nil {
		var apiError smithy.APIError
		if errors.As(err, &apiError) {
			switch apiError.ErrorCode() {
			case "NotFound":
				return false, nil
			default:
				return false, fmt.Errorf("unexpected error while retrieving bucket info")
			}
		}
	}
	return true, nil
}

func (agent *S3PersistentAgent) CreateDir(dirName string) error {
	_, err := agent.sdkClient.HeadObject(context.TODO(), &s3.HeadObjectInput{
		Bucket: aws.String(BUCKET_NAME),
		Key:    aws.String(dirName),
	})
	if err != nil {
		return fmt.Errorf("could not create directory \"%v\" due to:\n%v", dirName, err)
	}
	return nil
}

func (agent *S3PersistentAgent) UploadFile(dirs []string, fileName string, file io.Reader) error {
	filePath := strings.Join(dirs, "/") + "/" + fileName
	_, err := agent.sdkClient.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(BUCKET_NAME),
		Key:    aws.String(filePath),
		Body:   file,
	})
	if err != nil {
		return fmt.Errorf("could not upload file \"%v\" due to: %v", fileName, err)
	}
	return nil
}
