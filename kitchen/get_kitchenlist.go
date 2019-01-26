package kitchen

import (
	"encoding/csv"
	"io"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/vroup/mo-iwd-sa/config"
	"github.com/vroup/mo-iwd-sa/coordinate"
)

// GetKitchenList return kitchen list from prepared data csv
func GetKitchenList(conf *config.Config) []*Kitchen {
	var kitchenList []*Kitchen
	csvPath := "Kitchen " + conf.DeliveryDate + ".csv"
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

		id := record[0]

		latitude, err := strconv.ParseFloat(record[1], 64)
		longitude, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			log.Fatal(err)
		}

		coordinate := &coordinate.Coordinate{
			Latitude:  latitude,
			Longitude: longitude,
		}

		minCap, err := strconv.Atoi(record[3])
		optCap, err := strconv.Atoi(record[4])
		maxCap, err := strconv.Atoi(record[5])
		if err != nil {
			log.Fatal(err)
		}

		capacity := &Capacity{
			Minimum: minCap,
			Optimum: optCap,
			Maximum: maxCap,
		}
		kitchen := &Kitchen{
			ID:         id,
			Coordinate: coordinate,
			Capacity:   capacity,
		}

		kitchenList = append(kitchenList, kitchen)

	}
	return kitchenList
}
