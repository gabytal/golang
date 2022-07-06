package main

import (
	"amithamanyak/packages/s3"
	"amithamanyak/packages/utils"
)

func main() {
	bucket := utils.GetCliArg(1)
	fileName := utils.GetCliArg(2)

	s3.DownloadS3(bucket, fileName)
}
