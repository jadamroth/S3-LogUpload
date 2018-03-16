# S3-LogUpload
Simple Go program that uploads access &amp; error logs to S3 bucket in set time intervals
#### How to get started:
- Download [AWS Go SDK](https://github.com/aws/aws-sdk-go)
  - go get github.com/aws/aws-sdk-go
- Create AWS account
  - Create new role with access to S3
  - Create new user with AWS programmatic access keys 
  - Create S3 bucket & assign security policies (restrict public access and properly configure ACL, Bucket Policy, CORS, etc)
- Set config.json parameters (make sure config file is in same directory as binary)

#### Configuration Parameters (all are strings unless where noted):
- appName
  - Used for setting S3 Path (s3BucketName/appName/{access || error}/{timestamp}.txt)
- intervalUnit
  - Unit of time measurement (one letter representation)
    - s or S = seconds
    - m or M = minutes
    - h or H = hours
- uploadInteval (int)
  - Number of intervals (used with intervalUnit to create full time duration, e.g. 15 seconds)
- deleteContentAfterUpload (bool)
  - Set true if you want to empty log file content after uploading
- logPaths
  - access
    - path to access log file
  - error
    - path to error log file
- aws
  - access 
    - programmatic access key from created AWS user
  - secret
    - programmatic secret key from created AWS user
  - bucketName
    - S3 bucket Name
  - region
    - [AWS region](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.RegionsAndAvailabilityZones.html)

#### How it works
1) Loops over set time interval (uploadInterval * intervalUnit)
2) Once the timer ends, upload the access & error log to S3
3) Restart timer & repeat #1

####  Tested with:
- [golang 1.9.2](https://github.com/golang/go)
- This has only been tested in development and not in production but will update once it has. 
- In production, I would not place AWS access keys in configuration file. At a minimum, I would set as environment 
variables or, more securely, use a tool such as [Vault](https://github.com/hashicorp/vault).
