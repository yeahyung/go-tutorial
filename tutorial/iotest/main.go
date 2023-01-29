package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	currentDirectory, _ := os.Getwd()
	metastoreConfigurationFile := currentDirectory + "/iotest/metastore-site.xml"
	fmt.Println(metastoreConfigurationFile)

	bytes, _ := ioutil.ReadFile(metastoreConfigurationFile)

	configurationContent := string(bytes)

	if _, err := os.Stat(currentDirectory + "/iotest/temp_dir/"); os.IsNotExist(err) {
		fmt.Println(err)
		mkdirErr := os.Mkdir(currentDirectory+"/iotest/temp_dir/", os.ModePerm)
		if mkdirErr != nil {
			fmt.Println(mkdirErr)
		}
	}
	if deleteErr := os.Remove(currentDirectory + "/iotest/temp_dir/metastore-site.xml"); deleteErr != nil {
		fmt.Println(deleteErr)
	}
	err := ioutil.WriteFile(currentDirectory+"/iotest/temp_dir/metastore-site.xml", []byte(strings.Replace(configurationContent, "${WAREHOUSE_DIR}", "s3a://bucket", 1)), os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
}
