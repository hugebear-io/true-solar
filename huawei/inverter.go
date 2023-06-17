package huawei

import (
	"fmt"
	"io"
	"net/http"

	"github.com/hugebear-io/gofiber/errors"
	"github.com/hugebear-io/gofiber/fabric"
	"github.com/hugebear-io/gofiber/requestor"
)

type HuaweiInverterOptions struct{}

type huaweiInverter struct {
	URL     string
	token   string
	headers map[string]string
	options *HuaweiInverterOptions
}

func NewHuaweiInverter(username, password string, options *HuaweiInverterOptions) HuaweiInverter {
	obj := huaweiInverter{
		URL:     URL_VERSION1,
		token:   "",
		options: options,
	}

	token, err := obj.GetToken(username, password)
	if err != nil {
		panic(err)
	}

	obj.token = token
	obj.headers = map[string]string{AUTH_HEADER: token}
	return &obj
}

func (obj huaweiInverter) GetToken(username, password string) (string, error) {
	url := obj.URL + "/login"
	data := map[string]interface{}{
		"userName":   username,
		"systemCode": password,
	}

	req, err := requestor.PrepareHttpRequest(http.MethodPost, url, nil, data)
	if err != nil {
		return "", nil
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	if len(resBody) == 0 {
		return "", nil
	}

	var result APIGetTokenResponse
	if err := fabric.Recast(resBody, &result); err != nil {
		return "", err
	}

	meta := result.APIDefaultResponse.Result()
	if !meta.Success {
		errMsg := fmt.Sprintf("[%v]: %v", meta.FailCode, meta.Message)
		return "", errors.NewInternalError(errMsg)
	}

	mapCookies := make(map[string]string)
	for _, c := range res.Cookies() {
		mapCookies[c.Name] = c.Value
	}

	if val, ok := mapCookies[AUTH_HEADER]; ok {
		return val, nil
	}

	return "", errors.NewInternalError("huawei inverter does not receive token")
}

func (obj huaweiInverter) GetPlantList() ([]PlantItem, error) {
	url := obj.URL + "/getStationList"
	req, err := requestor.PrepareHttpRequest(http.MethodPost, url, obj.headers, nil)
	if err != nil {
		return nil, err
	}

	res, status, err := requestor.PrepareHttpResponse[APIGetPlantListResponse](req)
	if err != nil {
		return nil, err
	}

	if status == http.StatusTooManyRequests {
		return nil, ErrorTooManyRequest
	}

	meta := res.APIDefaultResponse.Result()
	if !meta.Success {
		errMsg := fmt.Sprintf("[%v]: %v", meta.FailCode, meta.Message)
		return nil, errors.NewInternalError(errMsg)
	}

	plantItem := []PlantItem{}
	for _, item := range res.Data {
		tmp := item.Result()
		plantItem = append(plantItem, tmp)
	}

	return plantItem, nil
}

func (obj huaweiInverter) GetPlantRealtimeData(plantCode string) ([]PlantRealtimeData, error) {
	url := URL_VERSION1 + "/getStationRealKpi"
	data := map[string]interface{}{
		"stationCodes": plantCode,
	}

	req, err := requestor.PrepareHttpRequest(http.MethodPost, url, obj.headers, data)
	if err != nil {
		return nil, err
	}

	res, status, err := requestor.PrepareHttpResponse[APIGetPlantRealtimeDataResponse](req)
	if err != nil {
		return nil, err
	}

	if status == http.StatusTooManyRequests {
		return nil, ErrorTooManyRequest
	}

	meta := res.APIDefaultResponse.Result()
	if !meta.Success {
		errMsg := fmt.Sprintf("[%v]: %v", meta.FailCode, meta.Message)
		return nil, errors.NewInternalError(errMsg)
	}

	result := []PlantRealtimeData{}
	for _, item := range res.Data {
		result = append(result, item.Result())
	}

	return result, nil
}

func (obj huaweiInverter) GetDailyPlantData(plantCode string, timestamp int64) ([]HistoricalPlantData, error) {
	url := URL_VERSION1 + "/getKpiStationDay"
	data := map[string]interface{}{
		"stationCodes": plantCode,
		"collectTime":  timestamp,
	}

	req, err := requestor.PrepareHttpRequest(http.MethodPost, url, obj.headers, data)
	if err != nil {
		return nil, err
	}

	res, status, err := requestor.PrepareHttpResponse[APIGetHistoricalPlantDataResponse](req)
	if err != nil {
		return nil, err
	}

	if status == http.StatusTooManyRequests {
		return nil, ErrorTooManyRequest
	}

	meta := res.APIDefaultResponse.Result()
	if !meta.Success {
		errMsg := fmt.Sprintf("[%v]: %v", meta.FailCode, meta.Message)
		return nil, errors.NewInternalError(errMsg)
	}

	result := []HistoricalPlantData{}
	for _, item := range res.Data {
		result = append(result, item.Result())
	}

	return result, nil
}

func (obj huaweiInverter) GetMonthlyPlantData(plantCode string, timestamp int64) ([]HistoricalPlantData, error) {
	url := URL_VERSION1 + "/getKpiStationMonth"
	data := map[string]interface{}{
		"stationCodes": plantCode,
		"collectTime":  timestamp,
	}

	req, err := requestor.PrepareHttpRequest(http.MethodPost, url, obj.headers, data)
	if err != nil {
		return nil, err
	}

	res, status, err := requestor.PrepareHttpResponse[APIGetHistoricalPlantDataResponse](req)
	if err != nil {
		return nil, err
	}

	if status == http.StatusTooManyRequests {
		return nil, ErrorTooManyRequest
	}

	meta := res.APIDefaultResponse.Result()
	if !meta.Success {
		errMsg := fmt.Sprintf("[%v]: %v", meta.FailCode, meta.Message)
		return nil, errors.NewInternalError(errMsg)
	}

	result := []HistoricalPlantData{}
	for _, item := range res.Data {
		result = append(result, item.Result())
	}

	return result, nil
}

func (obj huaweiInverter) GetYearlyPlantData(plantCode string, timestamp int64) ([]HistoricalPlantData, error) {
	url := URL_VERSION1 + "/getKpiStationYear"
	data := map[string]interface{}{
		"stationCodes": plantCode,
		"collectTime":  timestamp,
	}

	req, err := requestor.PrepareHttpRequest(http.MethodPost, url, obj.headers, data)
	if err != nil {
		return nil, err
	}

	res, status, err := requestor.PrepareHttpResponse[APIGetHistoricalPlantDataResponse](req)
	if err != nil {
		return nil, err
	}

	if status == http.StatusTooManyRequests {
		return nil, ErrorTooManyRequest
	}

	meta := res.APIDefaultResponse.Result()
	if !meta.Success {
		errMsg := fmt.Sprintf("[%v]: %v", meta.FailCode, meta.Message)
		return nil, errors.NewInternalError(errMsg)
	}

	result := []HistoricalPlantData{}
	for _, item := range res.Data {
		result = append(result, item.Result())
	}

	return result, nil
}

func (obj huaweiInverter) GetDeviceList(plantCode string) ([]DeviceItem, error) {
	url := URL_VERSION1 + "/getDevList"
	data := map[string]interface{}{
		"stationCodes": plantCode,
	}

	req, err := requestor.PrepareHttpRequest(http.MethodPost, url, obj.headers, data)
	if err != nil {
		return nil, err
	}

	res, status, err := requestor.PrepareHttpResponse[APIGetDeviceListResponse](req)
	if err != nil {
		return nil, err
	}

	if status == http.StatusTooManyRequests {
		return nil, ErrorTooManyRequest
	}

	meta := res.APIDefaultResponse.Result()
	if !meta.Success {
		errMsg := fmt.Sprintf("[%v]: %v", meta.FailCode, meta.Message)
		return nil, errors.NewInternalError(errMsg)
	}

	result := []DeviceItem{}
	for _, item := range res.Data {
		result = append(result, item.Result())
	}

	return result, nil
}

func (obj huaweiInverter) GetDeviceRealtimeData(deviceID, deviceTypeID string) ([]RealtimeDeviceData, error) {
	url := URL_VERSION1 + "/getDevRealKpi"
	data := map[string]interface{}{
		"devIds":    deviceID,
		"devTypeId": deviceTypeID,
	}

	req, err := requestor.PrepareHttpRequest(http.MethodPost, url, obj.headers, data)
	if err != nil {
		return nil, err
	}

	res, status, err := requestor.PrepareHttpResponse[APIGetRealtimeDeviceDataResponse](req)
	if err != nil {
		return nil, err
	}

	if status == http.StatusTooManyRequests {
		return nil, ErrorTooManyRequest
	}

	meta := res.APIDefaultResponse.Result()
	if !meta.Success {
		errMsg := fmt.Sprintf("[%v]: %v", meta.FailCode, meta.Message)
		return nil, errors.NewInternalError(errMsg)
	}

	result := []RealtimeDeviceData{}
	for _, item := range res.Data {
		result = append(result, item.Result())
	}

	return result, nil
}

func (obj huaweiInverter) GetDailyDeviceData(deviceID, deviceTypeID string, timestamp int64) ([]HistoricalDeviceData, error) {
	url := URL_VERSION1 + "/getDevKpiDay"
	data := map[string]interface{}{
		"devIds":      deviceID,
		"devTypeId":   deviceTypeID,
		"collectTime": timestamp,
	}

	req, err := requestor.PrepareHttpRequest(http.MethodPost, url, obj.headers, data)
	if err != nil {
		return nil, err
	}

	res, status, err := requestor.PrepareHttpResponse[APIGetHistoricalDeviceDataResponse](req)
	if err != nil {
		return nil, err
	}

	if status == http.StatusTooManyRequests {
		return nil, ErrorTooManyRequest
	}

	meta := res.APIDefaultResponse.Result()
	if !meta.Success {
		errMsg := fmt.Sprintf("[%v]: %v", meta.FailCode, meta.Message)
		return nil, errors.NewInternalError(errMsg)
	}

	result := []HistoricalDeviceData{}
	for _, item := range res.Data {
		result = append(result, item.Result())
	}

	return result, nil
}

func (obj huaweiInverter) GetMonthlyDeviceData(deviceID, deviceTypeID string, timestamp int64) ([]HistoricalDeviceData, error) {
	url := URL_VERSION1 + "/getDevKpiMonth"
	data := map[string]interface{}{
		"devIds":      deviceID,
		"devTypeId":   deviceTypeID,
		"collectTime": timestamp,
	}

	req, err := requestor.PrepareHttpRequest(http.MethodPost, url, obj.headers, data)
	if err != nil {
		return nil, err
	}

	res, status, err := requestor.PrepareHttpResponse[APIGetHistoricalDeviceDataResponse](req)
	if err != nil {
		return nil, err
	}

	if status == http.StatusTooManyRequests {
		return nil, ErrorTooManyRequest
	}

	meta := res.APIDefaultResponse.Result()
	if !meta.Success {
		errMsg := fmt.Sprintf("[%v]: %v", meta.FailCode, meta.Message)
		return nil, errors.NewInternalError(errMsg)
	}

	result := []HistoricalDeviceData{}
	for _, item := range res.Data {
		result = append(result, item.Result())
	}

	return result, nil
}

func (obj huaweiInverter) GetYearlyDeviceData(deviceID, deviceTypeID string, timestamp int64) ([]HistoricalDeviceData, error) {
	url := URL_VERSION1 + "/getDevKpiYear"
	data := map[string]interface{}{
		"devIds":      deviceID,
		"devTypeId":   deviceTypeID,
		"collectTime": timestamp,
	}

	req, err := requestor.PrepareHttpRequest(http.MethodPost, url, obj.headers, data)
	if err != nil {
		return nil, err
	}

	res, status, err := requestor.PrepareHttpResponse[APIGetHistoricalDeviceDataResponse](req)
	if err != nil {
		return nil, err
	}

	if status == http.StatusTooManyRequests {
		return nil, ErrorTooManyRequest
	}

	meta := res.APIDefaultResponse.Result()
	if !meta.Success {
		errMsg := fmt.Sprintf("[%v]: %v", meta.FailCode, meta.Message)
		return nil, errors.NewInternalError(errMsg)
	}

	result := []HistoricalDeviceData{}
	for _, item := range res.Data {
		result = append(result, item.Result())
	}

	return result, nil
}

func (obj huaweiInverter) GetDeviceAlarmList(plantCode string, from, to int64) ([]DeviceAlarmItem, error) {
	url := URL_VERSION1 + "/getAlarmList"
	data := map[string]interface{}{
		"stationCodes": plantCode,
		"beginTime":    from,
		"endTime":      to,
		"language":     HUAWEI_LANG_ENGLISH,
	}

	req, err := requestor.PrepareHttpRequest(http.MethodPost, url, obj.headers, data)
	if err != nil {
		return nil, err
	}

	res, status, err := requestor.PrepareHttpResponse[APIGetDeviceAlarmListResponse](req)
	if err != nil {
		return nil, err
	}

	if status == http.StatusTooManyRequests {
		return nil, ErrorTooManyRequest
	}

	meta := res.APIDefaultResponse.Result()
	if !meta.Success {
		errMsg := fmt.Sprintf("[%v]: %v", meta.FailCode, meta.Message)
		return nil, errors.NewInternalError(errMsg)
	}

	result := []DeviceAlarmItem{}
	for _, item := range res.Data {
		result = append(result, item.Result())
	}

	return result, nil
}
