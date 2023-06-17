package main

import (
	"log"
	"time"

	"github.com/hugebear-io/gofiber/fabric"
	"github.com/hugebear-io/true-solar/solarman"
)

var (
	username           = "u1.invt.th@gmail.com"
	password           = "Truec[8mugiup18"
	appID              = "202010143565002"
	appSecret          = "222c202135013aee622c71cdf8c47757"
	companyID          = 5751
	plantID            = 1765633
	deviceSN_inverter  = "F12207002767"
	deviceSN_collector = "2004160425"
)

func main() {
	loc := fabric.SetTimeZone("Asia/Bangkok")
	from := time.Date(2023, 1, 1, 0, 0, 0, 0, loc).Unix()
	_ = from
	to := time.Date(2023, 1, 15, 0, 0, 0, 0, loc).Unix()
	_ = to

	obj := solarman.NewSolarmanInverter(username, password, appID, appSecret, nil)
	business, _ := obj.GetBusinessToken(companyID)
	result, err := obj.GetDeviceAlertList(business.AccessToken, deviceSN_collector, from, to)
	if err != nil {
		log.Fatal(err)
	}

	fabric.PrintJSON(map[string]interface{}{"result": result})
}
