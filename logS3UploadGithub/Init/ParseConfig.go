package Init

import (
	"io/ioutil"
	"fmt"
	"os"
	"encoding/json"
	"time"
)

type Config struct{
	AppName				string			`json:"appName"`
	IntervalUnit		string			`json:"intervalUnit"`
	UploadInterval		time.Duration 	`json:"uploadInterval"`
	DeleteFileContent 	bool			`json:"deleteContentAfterUpload"`
	Path 				*LogPaths		`json:"logPaths"`
	Aws					*Aws			`json:"aws"`
}
type LogPaths struct{
	AccessLogPath	string	`json:"access"`
	ErrorLogPath	string	`json:"error"`
}
type Aws struct{
	Access		string	`json:"access"`
	Secret		string	`json:"secret"`
	BucketName 	string 	`json:"bucketName"`
	Region	 	string 	`json:"region"`
}


func ParseConfiguration() *Config{

	raw, err := ioutil.ReadFile("./config.json")

	if err != nil { // error reading configuration file, print error & exit program
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var config Config

	json.Unmarshal(raw, &config)

	return &config
}