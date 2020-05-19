package s3client

import (
	"github.com/alexTenFive/go-aws-s3-client/helpers"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

// BucketList return bucket list items
func BucketList(bucket string) []*s3.Object {
	resp, err := GetClient().ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String(bucket)})
	if err != nil {
		helpers.ExitErrorf("Unable to list items in bucket %q, %v", bucket, err)
	}
	return resp.Contents
}
