package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func (conf *Config) ReadMaxValue() {
	maxVFile, err := os.Open("prepared data/max_value_" + conf.DataSize + ".json")
	if err != nil {
		log.Fatal(err)
	}
	defer maxVFile.Close()
	jsonBytes, err := ioutil.ReadAll(maxVFile)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(jsonBytes, &conf.MaxValue)
}
