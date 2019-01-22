package order

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

// GetOrderList return order list on the delivery date. The order retrieved from the csv first is checked for one that has quantity more than driver's capacity. That particular one will be split to individual orders less equal than the driver's max capacity
func GetOrderList(conf config.Config) []*Order {
	var orderList []*Order

	csvPath := "Order " + conf.DeliveryDate + ".csv"
	file, err := ioutil.ReadFile("../prepared data/" + csvPath)
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
		userID := record[1]

		qty, err := strconv.Atoi(record[2])
		if err != nil {
			log.Fatal(err)
		}

		latitude, err := strconv.ParseFloat(record[3], 64)
		longitude, err := strconv.ParseFloat(record[4], 64)
		if err != nil {
			log.Fatal(err)
		}

		coordinate := &coordinate.Coordinate{
			Latitude:  latitude,
			Longitude: longitude,
		}

		order := &Order{
			ID:         id,
			UserID:     userID,
			Quantity:   qty,
			Coordinate: coordinate,
		}

		newOrderList := splitExcessOrder(order, conf)
		orderList = append(orderList, newOrderList...)

	}

	return orderList
}

// add idx trailing the excessive order to differentiate them on the edgesoilmap of the iwd
func splitExcessOrder(order *Order, conf config.Config) []*Order {
	var orderList []*Order
	for idx := 1; order.Quantity > conf.MaxDriverCapacity; idx++ {
		excessiveID := strconv.Itoa(idx)
		newOrder := &Order{
			ID:         order.ID + "-" + excessiveID,
			UserID:     order.UserID,
			Quantity:   conf.MaxDriverCapacity,
			Coordinate: order.Coordinate,
		}
		orderList = append(orderList, newOrder)
		order.Quantity -= conf.MaxDriverCapacity
	}
	orderList = append(orderList, order)
	return orderList
}
