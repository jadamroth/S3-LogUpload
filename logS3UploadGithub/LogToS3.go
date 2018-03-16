package main

import (
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws"
	"logS3UploadGithub/Init"
	"logS3UploadGithub/Aws"
	"time"
	"fmt"
)

// parameters from config.json file (place as same directory as go binary)
var appName *string // name of service/application
var uploadInterval time.Duration // time interval for how often we upload to S3
var accessLogPath *string
var errorLogPath *string

func init(){

	// get parameters from configuration file
	config := Init.ParseConfiguration()


	// set AWS credentials
	creds := credentials.NewStaticCredentials(config.Aws.Access, config.Aws.Secret, "")

	// check for AWS credentials error
	_, err := creds.Get()
	if err != nil{
		fmt.Printf("Bad AWS credentials: %s", err)
	}

	Aws.Cfg = aws.NewConfig().WithRegion(config.Aws.Region).WithCredentials(creds) // create new AWS configuration


	// create map for converting representing string to time interval (i.e. s = second, m = minute, h = hour)
	stringToDuration := map[string] time.Duration{
		"s": time.Second,
		"S": time.Second,

		"m": time.Minute,
		"M": time.Minute,

		"h": time.Hour,
		"H": time.Hour,
	}

	// set parameters from config.json
	uploadInterval = time.Duration(config.UploadInterval * stringToDuration[config.IntervalUnit] ) // multiply time (as a number) by interval (i.e. 15 * seconds)
	appName = &config.AppName // name of service

	Aws.DeleteFileContent = config.DeleteFileContent // if we should delete log file content after uploading to S3
	Aws.S3BucketName = &config.Aws.BucketName

	accessLogPath = &config.Path.AccessLogPath
	errorLogPath = &config.Path.ErrorLogPath
}


func main() {

	// loop in intervals for S3 uploads
	for{
		fmt.Println("Timer started.")
		uploadTimer := time.NewTimer(uploadInterval) // set timer

		<- uploadTimer.C // receive notice that timer expired

		fmt.Println("Timer expired and uploading files to S3.")
		uploadTimer.Reset(uploadInterval) // reset timer

		go Aws.GetContentFromFile(*accessLogPath, "access", *appName) // upload access.log
		go Aws.GetContentFromFile(*errorLogPath, "error", *appName) // upload error.log
	}
}