package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/alexTenFive/s3-client/helpers"
	"github.com/alexTenFive/s3-client/s3client"
)

const (
	CreateBucketCommand     = "create"
	ListBucketsCommand      = "buckets"
	BucketCommand           = "bucket"
	BucketUploadItemCommand = "upload"
	BucketRemoveItemCommand = "remove"
)

func main() {
	if len(os.Args) < 2 {
		helpers.ExitErrorf("command name missing!\nUsage: %s <command> <param?>", os.Args[0])
		os.Exit(1)
	}

	cmd := os.Args[1]
	switch cmd {
	case CreateBucketCommand:
		if len(os.Args) != 3 {
			helpers.ExitErrorf("bucket name missing!\nUsage: %s create <bucket_name>", os.Args[0])
			os.Exit(1)
		}
		s3client.CreateBucket(os.Args[2])
	case ListBucketsCommand:
		list := s3client.ListBuckets()
		fmt.Println("Buckets list:")
		for _, b := range list {
			fmt.Printf("* %s created on %s\n", aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
		}
	case BucketCommand:
		if len(os.Args) < 3 {
			helpers.ExitErrorf("bucket name missing!\nUsage: %s bucket <bucket_name>", os.Args[0])
			os.Exit(1)
		}

		bucket := os.Args[2]

		if len(os.Args) == 3 {
			list := s3client.BucketList(bucket)
			fmt.Printf("List items in bucket <%s>:\n", bucket)
			for _, item := range list {
				fmt.Printf("%+v\n", item)
			}
			break
		}

		subcmd := os.Args[3]
		if len(os.Args) != 5 {
			helpers.ExitErrorf("filepath missing!\nUsage: %s bucket %s %s <itempath>", os.Args[0], bucket, subcmd)
			os.Exit(1)
		}

		switch subcmd {
		case BucketUploadItemCommand:
			filepath := os.Args[4]
			s3client.UploadToBucket(bucket, filepath)
		case BucketRemoveItemCommand:
		}

	}
	os.Exit(0)
}
