package rating

import (
	"encoding/csv"
	"io"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/vroup/mo-iwd-sa/config"
)

// GetRatingMap read rating csv on this date and return a map out of it
func GetRatingMap(conf *config.Config) Map {
	ratingMap := newMap()
	csvPath := "Rating " + conf.DataSize + ".csv"
	file, err := ioutil.ReadFile("prepared data/" + csvPath)
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(strings.NewReader(string(file)))

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		userID := record[0]
		kitchenID := record[1]
		key := userID + "+" + kitchenID

		rate, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			log.Fatal(err)
		}
		ratingMap[key] = rate
	}
	return ratingMap
}
