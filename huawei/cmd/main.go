package main

import (
	"fmt"
	"time"

	"github.com/hugebear-io/gofiber/fabric"
	"github.com/hugebear-io/true-solar/huawei"
)

const (
	username     = "trueapi"
	password     = "Trueapi12@"
	plantCode    = "2F3365B1D745457F8EF49F0E23248875"
	deviceID     = "-49678782321252"
	deviceTypeID = "1"
)

func main() {
	loc := fabric.SetTimeZone("Asia/Bangkok")
	now := time.Now()

	from := time.Date(2022, 5, 1, 0, 0, 0, 0, loc)
	to := time.Date(2023, 5, 2, 0, 0, 0, 0, loc)

	collectTime := now.UnixMilli()
	_ = collectTime

	huawei := huawei.NewHuaweiInverter(username, password, nil)
	result, err := huawei.GetDeviceAlarmList(plantCode, from.UnixMilli(), to.UnixMilli())
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", result)
}
