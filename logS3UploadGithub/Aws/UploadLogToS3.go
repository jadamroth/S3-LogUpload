package Aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"fmt"
	"bytes"
	"io/ioutil"
	"time"
)

var Cfg *aws.Config // create new AWS configuration
var DeleteFileContent bool // if we should delete log file content after uploading
var S3BucketName *string

func GetContentFromFile(logPath string, logType string, appName string){

	// get content from file
	fileContents, rErr := ioutil.ReadFile(logPath)

	if rErr != nil{ // error reading content from file

		fmt.Println("Error reading " + logType + " file.")

	} else{ // if no errors reading file, proceed with upload

		// upload log content to S3 - return response & error
		uploadResponse, err := uploadLogToS3(fileContents, appName + "/" + logType + "/" + time.Now().String() + ".txt")

		if err != nil{ // s3 upload error

			fmt.Println("S3 upload Error -> " + err.Error())

		} else {

			// if we should delete file content after uploading
			if DeleteFileContent == true {

				// empty log file contents, so we're not re-uploading same logs
				wErr := ioutil.WriteFile(logPath, []byte{}, 0644)

				if wErr != nil { // error emptying log file content

					fmt.Println("Error writing to " + logType + " file.")
				}
			}

			// do something with S3 response
			fmt.Printf("response %s", uploadResponse)
		}
	}
}


func uploadLogToS3(fileContent []byte, s3path string) (string, error){

	s3Session, err := session.NewSession(Cfg) // create new S3 session

	// convert to byte reader for S3 Body
	fileBytes := bytes.NewReader(fileContent)

	// set S3 parameters
	params := &s3.PutObjectInput{
		Bucket: aws.String(*S3BucketName),
		Key: aws.String(s3path),
		Body: fileBytes,
		ContentLength: aws.Int64(fileBytes.Size()),
	}

	resp, err := s3.New(s3Session).PutObject(params) // put object in S3 bucket

	if err != nil {
		return "", err
	}

	return awsutil.StringValue(resp), nil
}