package url

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/erobx/tradeups/backend/pkg/common"
)

type PresignedUrlManager struct {
    cache map[string]*UrlEntry
    mu sync.RWMutex
    client *s3.PresignClient
    bucket string
}

type UrlEntry struct {
    Url string
    ExpiresAt time.Time
}

func NewPresignedUrlManager(bucket string) *PresignedUrlManager {
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

    return &PresignedUrlManager{
        cache: make(map[string]*UrlEntry),
        client: s3.NewPresignClient(client),
        bucket: bucket,
    }
}

func (m *PresignedUrlManager) GetUrls(imageKeys []string) map[string]string {
    result := make(map[string]string)
    keysToGenerate := []string{}

    m.mu.RLock()
    now := time.Now()
    for _, key := range imageKeys {
        if entry, exists := m.cache[key]; exists && now.Before(entry.ExpiresAt) {
            result[key] = entry.Url
        } else {
            keysToGenerate = append(keysToGenerate, key)
        }
    }
    m.mu.RUnlock()

    if len(keysToGenerate) > 0 {
        var wg sync.WaitGroup
        var mu sync.Mutex

        for _, key := range keysToGenerate {
            wg.Add(1)
            go func(key string) {
                defer wg.Done()
                out, err := m.client.PresignGetObject(context.Background(), &s3.GetObjectInput{
                    Bucket: &m.bucket,
                    Key: aws.String(common.PrefixKey(key)),
                }, func(opts *s3.PresignOptions) {
                        opts.Expires = time.Hour * 24
                    })
                //log.Println("Generated new presigned url for:", key)
                if err == nil {
                    mu.Lock()
                    result[key] = out.URL
                    m.mu.Lock()
                    m.cache[key] = &UrlEntry{
                        Url: out.URL,
                        ExpiresAt: time.Now().Add(23 * time.Hour),
                    }
                    m.mu.Unlock()
                    mu.Unlock()
                }
            }(key)
        }
        wg.Wait()
    }

    return result
}

