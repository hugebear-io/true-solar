package kstar

import (
	"crypto/sha1"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

type KStarInverterOptions struct {
}

type kstarInverter struct {
	URL      string
	username string
	password string
	options  *KStarInverterOptions
}

func NewKStarInverter(username, password string, options *KStarInverterOptions) KStarInverter {
	obj := kstarInverter{
		URL:      URL_VERSION1,
		username: username,
		password: password,
		options:  options,
	}

	// encrypt password
	obj.password = DecodePassword(obj.password)

	return &obj
}

func (obj kstarInverter) EncryptParameter(data map[string]string) string {
	data["userCode"] = obj.username
	data["password"] = obj.password

	var keys []string
	for k := range data {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	var params []string
	for _, k := range keys {
		params = append(params, fmt.Sprintf("%s=%s", k, data[k]))
	}

	paramStr := strings.Join(params, "&")
	hash := sha1.New()
	hash.Write([]byte(paramStr))
	result := fmt.Sprintf("%x", hash.Sum(nil))
	return result
}

func (obj kstarInverter) GetPlantList() ([]PlantItem, error) {
	sign := obj.EncryptParameter(make(map[string]string))

	data := url.Values{}
	data.Set("userCode", obj.username)
	data.Set("password", obj.password)
	data.Set("sign", sign)

	url := obj.URL + "/power/info"
	req, err := prepareHttpRequest(http.MethodPost, url, data)
	if err != nil {
		return nil, err
	}

	res, _, err := prepareHttpResponse[APIGetPlantListResponse](req)
	if err != nil {
		return nil, err
	}

	result := []PlantItem{}
	for _, item := range res.Data {
		result = append(result, item.Result())
	}

	return result, nil
}

func (obj kstarInverter) GetDeviceListWithPagination(page, size int) (*DeviceData, error) {
	pageStr := strconv.Itoa(page)
	sizeStr := strconv.Itoa(size)
	params := map[string]string{
		"PageNum":  pageStr,
		"PageSize": sizeStr,
	}

	sign := obj.EncryptParameter(params)
	data := url.Values{}
	data.Set("userCode", obj.username)
	data.Set("password", obj.password)
	data.Set("PageNum", pageStr)
	data.Set("PageSize", sizeStr)
	data.Set("sign", sign)

	url := obj.URL + "/inverter/list"
	req, err := prepareHttpRequest(http.MethodPost, url, data)
	if err != nil {
		return nil, err
	}

	res, _, err := prepareHttpResponse[APIGetDeviceListResponse](req)
	if err != nil {
		return nil, err
	}

	result := res.Data.Result()
	return &result, nil
}

func (obj kstarInverter) GetDeviceList() ([]DeviceItem, error) {
	result := []DeviceItem{}
	var page int = 1
	for {
		data, err := obj.GetDeviceListWithPagination(page, MAX_PAGE_SIZE)
		if err != nil {
			return nil, err
		}

		if len(data.List) == 0 {
			break
		}

		result = append(result, data.List...)
		if len(result) >= data.Count {
			break
		}

		page += 1
	}
	return result, nil
}

func (obj kstarInverter) GetRealtimeDeviceData(deviceID string) (*DeviceInfo, error) {
	params := map[string]string{"deviceId": deviceID}
	sign := obj.EncryptParameter(params)

	data := url.Values{}
	data.Set("userCode", obj.username)
	data.Set("password", obj.password)
	data.Set("deviceId", deviceID)
	data.Set("sign", sign)

	url := obj.URL + "/device/real"
	req, err := prepareHttpRequest(http.MethodPost, url, data)
	if err != nil {
		return nil, err
	}

	res, _, err := prepareHttpResponse[APIGetRealtimeDeviceDataResponse](req)
	if err != nil {
		return nil, err
	}

	result := res.Data.Result()
	return &result, nil
}

func (obj kstarInverter) GetRealtimeAlarmListOfDevice(deviceID string) ([]DeviceAlarmItem, error) {
	params := map[string]string{"deviceId": deviceID}
	sign := obj.EncryptParameter(params)

	data := url.Values{}
	data.Set("userCode", obj.username)
	data.Set("password", obj.password)
	data.Set("deviceId", deviceID)
	data.Set("sign", sign)

	url := obj.URL + "/alarm/device/list"
	req, err := prepareHttpRequest(http.MethodPost, url, data)
	if err != nil {
		return nil, err
	}

	res, _, err := prepareHttpResponse[APIGetRealtimeAlarmListOfDeviceResponse](req)
	if err != nil {
		return nil, err
	}

	result := []DeviceAlarmItem{}
	for _, item := range res.Data {
		result = append(result, item.Result())
	}

	return result, nil
}

func (obj kstarInverter) GetRealtimeAlarmListOfPlantWithPagination(from, to int64, page, size int) (*DeviceAlarmData, error) {
	fromTime := time.Unix(from, 0)
	fromDate := fromTime.Format("2006-01-02")
	toTime := time.Unix(to, 0)
	toDate := toTime.Format("2006-01-02")

	pageStr := strconv.Itoa(page)
	sizeStr := strconv.Itoa(size)
	params := map[string]string{
		"stime":    fromDate,
		"etime":    toDate,
		"PageNum":  pageStr,
		"PageSize": sizeStr,
	}

	sign := obj.EncryptParameter(params)
	data := url.Values{}
	data.Set("userCode", obj.username)
	data.Set("password", obj.password)
	data.Set("PageNum", pageStr)
	data.Set("PageSize", sizeStr)
	data.Set("stime", fromDate)
	data.Set("etime", toDate)
	data.Set("sign", sign)

	url := obj.URL + "/real/alarm/list"
	req, err := prepareHttpRequest(http.MethodPost, url, data)
	if err != nil {
		return nil, err
	}

	res, _, err := prepareHttpResponse[APIGetDeviceAlarmListResponse](req)
	if err != nil {
		return nil, err
	}

	result := res.Data.Result()
	return &result, nil
}

func (obj kstarInverter) GetRealtimeAlarmListOfPlant(from, to int64) ([]DeviceAlarmItem, error) {
	result := []DeviceAlarmItem{}
	var page int = 1
	for {
		data, err := obj.GetRealtimeAlarmListOfPlantWithPagination(from, to, page, MAX_PAGE_SIZE)
		if err != nil {
			return nil, err
		}

		if len(data.Data) == 0 {
			break
		}

		result = append(result, data.Data...)
		if len(result) >= data.Count {
			break
		}
	}

	return result, nil
}
