package s3client

import (
	"fmt"

	"github.com/alexTenFive/s3-client/helpers"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

// RestoreFromBucket - restore removed item from bucket
func RestoreFromBucket(bucket, obj string) {
	_, err := GetClient().RestoreObject(&s3.RestoreObjectInput{Bucket: aws.String(bucket), Key: aws.String(obj), RestoreRequest: &s3.RestoreRequest{Days: aws.Int64(30)}})
	if err != nil {
		helpers.ExitErrorf("Could not restore %s in bucket %s, %v", obj, bucket, err)
	}

	fmt.Printf("%q should be restored to %q in about 4 hours\n", obj, bucket)
}
