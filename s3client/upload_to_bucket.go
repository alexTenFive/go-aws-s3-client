package s3client

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/alexTenFive/go-aws-s3-client/helpers"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

const tmpFilesPath = "C:/Users/flash/tmp"

// UploadToBucket uploads file to bucket by path
func UploadToBucket(bucket, filename string) {

	file, err := os.Open(filename)
	if err != nil {
		// Print the error and exit.
		helpers.ExitErrorf("Unable to open file %q: %v", filename, err)
	}
	defer file.Close()

	stat, _ := file.Stat()
	if stat.IsDir() {
		tarFileName := filepath.Clean(tmpFilesPath) + string(filepath.Separator) + filepath.Base(filename) + ".tar.gz"
		tarFile, err := os.Create(tarFileName)
		if err != nil {
			helpers.ExitErrorf("Cannot create temp archive %q: %v", tarFileName, err)
		}
		defer os.Remove(tarFileName) // remove archive after upload to s3

		if err = helpers.Tar(filename, tarFile); err != nil {
			helpers.ExitErrorf("Cannot tar archive %q: %v", tarFileName, err)
		}
		tarFile.Close() // confirm write to it

		tarFile, err = os.Open(tarFileName) // open for upload to s3
		if err != nil {
			helpers.ExitErrorf("Cannot open archive %q: %err\n", tarFileName)
		}
		defer tarFile.Close()

		err = upload(bucket, tarFile, tarFileName)
		if err != nil {
			helpers.ExitErrorf("Unable to upload %q to %q, %v", tarFileName, bucket, err)
		}
		return
	}
	err = upload(bucket, file, filename)
	if err != nil {
		helpers.ExitErrorf("Unable to upload %q to %q, %v", filename, bucket, err)
	}
}

func upload(bucket string, file *os.File, filename string) error {
	uploader := GetUploader()
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filepath.Base(filename)),
		Body:   file,
	})
	if err != nil {
		return err
	}

	fmt.Printf("Successfully uploaded %q to %q\n", filename, bucket)
	return nil
}
