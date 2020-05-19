package s3client

import (
	"github.com/alexTenFive/s3-client/helpers"
	"github.com/aws/aws-sdk-go/service/s3"
)

// ListBuckets s
func ListBuckets() []*s3.Bucket {
	result, err := GetClient().ListBuckets(nil)
	if err != nil {
		helpers.ExitErrorf("error when retrieving buckets list: %v", err)
	}

	return result.Buckets
}
