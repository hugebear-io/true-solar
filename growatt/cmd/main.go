package main

import (
	"log"
	"time"

	"github.com/hugebear-io/gofiber/fabric"
	"github.com/hugebear-io/true-solar/growatt"
)

var (
	username = "Trueupc1"
	token    = "33bfe28df3f24cd42a8de64de0e7036e"
	plantID  = 353799
	deviceSN = "MKE2A130BC"
)

func main() {
	now := time.Now()
	growatt := growatt.NewGrowattInverter(username, token)
	result, err := growatt.GetPbdAlertListWithPagination(deviceSN, now.Unix(), 1, 10)
	if err != nil {
		log.Fatal(err)
	}

	fabric.PrintJSON(result)
	// fabric.PrintJSON(map[string]interface{}{"result": result})
}
