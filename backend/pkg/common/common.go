package common

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func ReadPrivKey() ([]byte, error) {
	b, err := os.ReadFile(os.Getenv("PRIVATE_KEY"))
	return b, err
}

func ReadPubKey() ([]byte, error) {
	b, err := os.ReadFile(os.Getenv("PUBLIC_KEY"))
	return b, err
}

func GetPresignedURL(imageKey string) string {
	bucketName := "tradeups-images"
	endPoint := os.Getenv("S3_ENDPOINT")
	accessKeyId := os.Getenv("S3_ACCESS_KEY_ID")
	accessKey := os.Getenv("S3_ACCESS_KEY")

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyId, accessKey, "")),
		config.WithRegion("auto"),
	)
	if err != nil {
		log.Fatal(err)
	}
	config.WithRequestChecksumCalculation(0)
	config.WithResponseChecksumValidation(0)

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(endPoint)
	})

	presignClient := s3.NewPresignClient(client)

	res, err := presignClient.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: &bucketName,
		Key:    aws.String(imageKey),
	})
	if err != nil {
		log.Fatal("Couldn't get presigned URL for GetObject")
	}

	return res.URL
}
