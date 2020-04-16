package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/endpoints"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func init() {
	log.SetFlags(log.Lmicroseconds)
}

func run(ctx context.Context) error {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return fmt.Errorf("unable to load SDK config: %w", err)
	}
	cfg.Region = endpoints.UsWest2RegionID
	s3c := s3.New(cfg)
	resp, err := s3c.ListBucketsRequest(nil).Send(ctx)
	if err != nil {
		return fmt.Errorf("failed to list buckets: %w", err)
	}
	for i, bucket := range resp.Buckets {
		log.Printf("#%d: %s", i, aws.StringValue(bucket.Name))
	}
	return nil
}

func main() {
	if err := run(context.Background()); err != nil {
		log.Fatalf("error: %s", err)
	}
}
