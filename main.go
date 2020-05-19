package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/alexTenFive/go-aws-s3-client/helpers"
	"github.com/alexTenFive/go-aws-s3-client/s3client"
)

const (
	// CreateBucketCommand - creates bucket in storage
	CreateBucketCommand = "create"
	// DeleteBucketCommand - deletes bucket from storage
	DeleteBucketCommand = "delete"
	// ListBucketsCommand - display buckets list
	ListBucketsCommand = "buckets"
	// BucketCommand - for bucket operations
	BucketCommand = "bucket"
	// BucketUploadItemCommand - subcommand for <bucket> - upload item to bucket
	BucketUploadItemCommand = "upload"
	// BucketDownloadItemCommand - subcommand for <bucket> - download item from bucket
	BucketDownloadItemCommand = "download"
	// BucketRemoveItemCommand - subcommand for <bucket> - remove item from bucket
	BucketRemoveItemCommand = "remove"
	// BucketRestoreItemCommand - subcommand for <bucket> - restore removed item to bucket
	BucketRestoreItemCommand = "restore"
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
			helpers.ExitErrorf("bucket name missing!\nUsage: %s %s <bucket_name>", os.Args[0], CreateBucketCommand)
			os.Exit(1)
		}
		s3client.CreateBucket(os.Args[2])
	case DeleteBucketCommand:
		if len(os.Args) != 3 {
			helpers.ExitErrorf("bucket name missing!\nUsage: %s %s <bucket_name>", os.Args[0], DeleteBucketCommand)
			os.Exit(1)
		}
		s3client.DeleteBucket(os.Args[2])
	case ListBucketsCommand:
		list := s3client.ListBuckets()
		fmt.Println("Buckets list:")
		for _, b := range list {
			fmt.Printf("* %s created on %s\n", aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
		}
	case BucketCommand:
		if len(os.Args) < 3 {
			helpers.ExitErrorf("bucket name missing!\nUsage: %s %s <bucket_name>", os.Args[0], BucketCommand)
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
			helpers.ExitErrorf("filepath missing!\nUsage: %s %s %s %s <itempath>", os.Args[0], BucketCommand, bucket, subcmd)
			os.Exit(1)
		}

		item := os.Args[4]
		switch subcmd {
		case BucketUploadItemCommand:
			s3client.UploadToBucket(bucket, item)
		case BucketDownloadItemCommand:
			s3client.DownloadFromBucket(bucket, item)
		case BucketRemoveItemCommand:
			s3client.DeleteFromBucket(bucket, item)
		case BucketRestoreItemCommand:
			s3client.RestoreFromBucket(bucket, item)
		}

	}
	os.Exit(0)
}
