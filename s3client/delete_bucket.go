package s3client

import (
	"fmt"

	"github.com/alexTenFive/go-aws-s3-client/helpers"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

// DeleteBucket deletes s3 bucket
func DeleteBucket(bucket string) {
	_, err := GetClient().DeleteBucket(&s3.DeleteBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		helpers.ExitErrorf("Unable to delete bucket %q, %v", bucket, err)
	}

	// Wait until bucket is created before finishing
	fmt.Printf("Waiting for bucket %q to be deleted...\n", bucket)

	err = GetClient().WaitUntilBucketNotExists(&s3.HeadBucketInput{
		Bucket: aws.String(bucket),
	})

	if err != nil {
		helpers.ExitErrorf("Error occurred while waiting for bucket to be deleted, %v", bucket)
	}

	fmt.Printf("Bucket %q successfully deleted\n", bucket)
}
