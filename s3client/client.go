package s3client

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

const (
	defaultRegion = "eu-central-1"
	defaultUser   = "dev-user"
	// uploader & downloader
	defaultPartSize = 128 * 1024 * 1024 // 32 MB (32 * 1024 * 1024) bytes
)

var client *s3.S3
var uploader *s3manager.Uploader
var downloader *s3manager.Downloader
var sess *session.Session

func init() {
	var err error
	sess, err = session.NewSession(&aws.Config{
		Region:      aws.String(defaultRegion),
		Credentials: credentials.NewSharedCredentials("", defaultUser)})
	if err != nil {
		panic(fmt.Sprintf("Unable to create session: %v", err))
	}
}

// GetClient returns s3 aws client
func GetClient() *s3.S3 {
	if client == nil {
		// Create S3 service client
		client = s3.New(sess)
	}
	return client
}

// GetUploader returns s3manager aws uploader
func GetUploader() *s3manager.Uploader {
	if uploader == nil {
		uploader = s3manager.NewUploader(sess, func(u *s3manager.Uploader) {
			u.PartSize = defaultPartSize
		})
	}
	return uploader
}

// GetDownloader returns s3manager aws downloader
func GetDownloader() *s3manager.Downloader {
	if downloader == nil {
		downloader = s3manager.NewDownloader(sess, func(d *s3manager.Downloader) {
			d.PartSize = defaultPartSize
		})
	}
	return downloader
}
