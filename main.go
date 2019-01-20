package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/m-tree/coordinate"
	"github.com/m-tree/distance"
	"github.com/m-tree/mtree"
)

const maxEntry = 60

func main() {

	coordinateList := make([]*coordinate.Coordinate, 0)
	csvFile, err := os.Open("coordinate.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()
	r := csv.NewReader(csvFile)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		lat, err := strconv.ParseFloat(record[3], 64)
		lon, err := strconv.ParseFloat(record[4], 64)
		coord := &coordinate.Coordinate{
			Latitude:  lat,
			Longitude: lon,
		}
		coordinateList = append(coordinateList, coord)
	}

	distCalc := &distance.HaversineDistance{}
	splitMecha := &mtree.SplitMST{
		DistCalc: distCalc,
		MaxEntry: maxEntry,
	}
	tree := mtree.NewTree(maxEntry, splitMecha, distCalc)

	for idx := range coordinateList {
		tree.Insert(coordinateList[idx], idx)
	}
	start := time.Now()
	for idx := range coordinateList {
		tree.Remove(tree.Root, coordinateList[idx], idx)
	}
	elapsed := time.Since(start)
	fmt.Println(elapsed)

}
