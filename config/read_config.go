package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// ReadConfig return config from config.json file
func ReadConfig() *Config {
	var conf Config
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	jsonBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(jsonBytes, &conf)

	maxVFile, err := os.Open("prepared data/max_value_" + conf.DataSize + ".json")
	if err != nil {
		log.Fatal(err)
	}
	defer maxVFile.Close()
	jsonBytes, err = ioutil.ReadAll(maxVFile)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(jsonBytes, &conf.MaxValue)
	return &conf
}
