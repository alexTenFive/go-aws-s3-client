package s3client

import (
	"fmt"
	"os"

	"github.com/alexTenFive/s3-client/helpers"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

// DownloadFromBucket downloads item from bucket
func DownloadFromBucket(bucket, item string) {
	file, err := os.Create(item)
	if err != nil {
		helpers.ExitErrorf("Unable to create file destination %q: %v", item, err)
	}
	defer file.Close()

	downloader := GetDownloader()
	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(item),
		})
	if err != nil {
		defer func() {
			file.Close()
			os.Remove(item)
		}()
		fmt.Printf("Unable to download item %q, %v", item, err)
		return
	}

	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")
}
