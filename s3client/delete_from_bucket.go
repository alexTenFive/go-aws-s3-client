package s3client

import (
	"fmt"

	"github.com/alexTenFive/s3-client/helpers"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

// DeleteFromBucket delete item from bucket
func DeleteFromBucket(bucket, obj string) {
	_, err := GetClient().DeleteObject(&s3.DeleteObjectInput{Bucket: aws.String(bucket), Key: aws.String(obj)})
	if err != nil {
		helpers.ExitErrorf("Unable to delete object %q from bucket %q, %v", obj, bucket, err)
	}

	err = GetClient().WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(obj),
	})
	if err != nil {
		helpers.ExitErrorf("Object %q not exists in bucket %q: err", obj, bucket, err)
	}

	fmt.Printf("Object %q successfully deleted\n", obj)
}
