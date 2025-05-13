package main

import (
    "context"
    "fmt"
    "log"

    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
    cfg, err := config.LoadDefaultConfig(context.TODO())
    if err != nil {
        log.Fatalf("unable to load SDK config, %v", err)
    }

    client := s3.NewFromConfig(cfg)

    result, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
    if err != nil {
        log.Fatalf("unable to list buckets, %v", err)
    }

    fmt.Println("Buckets:")
    for _, b := range result.Buckets {
        fmt.Printf("* %s\n", *b.Name)
    }
}
