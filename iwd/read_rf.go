package iwd

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/vroup/mo-iwd-sa/config"
)

// ReadRF from file for IGD calculation
func ReadRF(conf *config.Config) []*Score {
	rf := make([]*Score, 0)
	rfPath := "rf_" + conf.DataSize + ".json"
	f, err := os.Open("prepared data/" + rfPath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	byteF, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(byteF, &rf)
	return rf
}
