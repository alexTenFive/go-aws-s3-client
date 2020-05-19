package s3client

import (
	"github.com/alexTenFive/s3-client/helpers"
	"github.com/aws/aws-sdk-go/service/s3"
)

// ListBuckets s
func ListBuckets() []*s3.Bucket {
	svc := GetClient()

	result, err := svc.ListBuckets(nil)
	if err != nil {
		helpers.ExitErrorf("error when retrieving buckets list: %v", err)
		return nil
	}

	return result.Buckets
}
