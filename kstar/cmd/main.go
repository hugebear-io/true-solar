package main

import (
	"fmt"
	"time"

	"github.com/hugebear-io/gofiber/fabric"
	"github.com/hugebear-io/true-solar/kstar"
)

var (
	username string = "u6.kst"
	password string = "Truec[8mugiup18"
	deviceId string = "I1102610779C083140020721"
)

func main() {
	loc := fabric.SetTimeZone("Asia/Bangkok")
	from := time.Date(2023, 1, 1, 0, 0, 0, 0, loc)
	to := time.Date(2023, 5, 1, 0, 0, 0, 0, loc)

	kstar := kstar.NewKStarInverter(username, password, nil)
	result, err := kstar.GetRealtimeAlarmListOfPlantWithPagination(from.Unix(), to.Unix(), 1, 200)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", result)
}
