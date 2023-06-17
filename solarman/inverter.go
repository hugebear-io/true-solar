package solarman

import (
	"net/http"

	"github.com/hugebear-io/gofiber/errors"
	"github.com/hugebear-io/gofiber/requestor"
)

type SolarmanInverterOptions struct{}

type solarmanInverter struct {
	URL       string
	username  string
	password  string
	appID     string
	appSecret string
	headers   map[string]string
	options   *SolarmanInverterOptions
}

func NewSolarmanInverter(username, password, appID, appSecret string, options *SolarmanInverterOptions) SolarmanInverter {
	obj := solarmanInverter{
		URL:       URL_VERSION1,
		username:  username,
		password:  DecodePassword(password),
		appID:     appID,
		appSecret: appSecret,
		options:   options,
	}

	tokenResp, err := obj.GetBasicToken()
	if err != nil {
		panic(err)
	}

	obj.headers = map[string]string{
		AUTHORIZATION_HEADER: buildAuthorizationToken(tokenResp.AccessToken),
	}

	return &obj
}

func (obj solarmanInverter) GetBasicToken() (*GetTokenResponse, error) {
	body := GetTokenBody{
		Username:  obj.username,
		Password:  obj.password,
		AppSecret: obj.appSecret,
	}

	url := obj.URL + "/account/v1.0/token?appId=" + obj.appID
	req, err := requestor.PrepareHttpRequest(http.MethodPost, url, nil, body)
	if err != nil {
		return nil, err
	}

	data, status, err := prepareHttpResponse[APIGetTokenResponse](req)
	if err != nil {
		return nil, err
	}

	result := data.Result()
	if status != http.StatusOK {
		return nil, errors.NewError(status, result.Message)
	}

	return &result, nil
}

func (obj solarmanInverter) GetBusinessToken(orgID int) (*GetTokenResponse, error) {
	body := GetTokenBody{
		Username:       obj.username,
		Password:       obj.password,
		OrganizationID: orgID,
		AppSecret:      obj.appSecret,
	}

	url := obj.URL + "/account/v1.0/token?appId=" + obj.appID
	req, err := requestor.PrepareHttpRequest(http.MethodPost, url, nil, body)
	if err != nil {
		return nil, err
	}

	data, status, err := prepareHttpResponse[APIGetTokenResponse](req)
	if err != nil {
		return nil, err
	}

	result := data.Result()
	if status != http.StatusOK {
		return nil, errors.NewError(status, result.Message)
	}

	return &result, nil
}

func (obj solarmanInverter) GetUserInfo() (*GetUserInfoResponse, error) {
	url := obj.URL + "/account/v1.0/info?language=en"
	req, err := requestor.PrepareHttpRequest(http.MethodPost, url, obj.headers, nil)
	if err != nil {
		return nil, err
	}

	data, status, err := prepareHttpResponse[APIGetUserInfoResponse](req)
	if err != nil {
		return nil, err
	}

	result := data.Result()
	if status != http.StatusOK {
		return nil, errors.NewError(status, result.Message)
	}

	return &result, nil
}

func (obj solarmanInverter) GetPlantListWithPagination(businessToken string, page, size int) (*GetPlantListResponse, error) {
	body := map[string]interface{}{
		"page": page,
		"size": size,
	}

	headers := map[string]string{AUTHORIZATION_HEADER: buildAuthorizationToken(businessToken)}
	url := obj.URL + "/station/v1.0/list?language=en"
	req, err := requestor.PrepareHttpRequest(http.MethodPost, url, headers, body)
	if err != nil {
		return nil, err
	}

	data, status, err := prepareHttpResponse[APIGetPlantListResponse](req)
	if err != nil {
		return nil, err
	}

	result := data.Result()
	if status != http.StatusOK {
		return nil, errors.NewError(status, result.Message)
	}

	return &result, nil
}

func (obj solarmanInverter) GetPlantList(businessToken string) ([]PlantItem, error) {
	plantList := []PlantItem{}
	page := 1
	for {
		res, err := obj.GetPlantListWithPagination(businessToken, page, MAX_PAGE_SIZE)
		if err != nil {
			return nil, err
		}

		plantList = append(plantList, res.PlantList...)
		page += 1

		if len(plantList) >= res.Total {
			break
		}
	}

	return plantList, nil
}

func (obj solarmanInverter) GetPlantBaseInfo(businessToken string, plantID int) (*GetPlantBaseInfoResponse, error) {
	body := StationBody{
		StationID: plantID,
	}

	headers := map[string]string{
		AUTHORIZATION_HEADER: buildAuthorizationToken(businessToken),
	}

	url := obj.URL + "/station/v1.0/base?language=en"
	req, err := requestor.PrepareHttpRequest(http.MethodPost, url, headers, body)
	if err != nil {
		return nil, err
	}

	data, status, err := prepareHttpResponse[APIGetPlantBaseInfoResponse](req)
	if err != nil {
		return nil, err
	}

	result := data.Result()
	if status != http.StatusOK {
		return nil, errors.NewError(status, result.Message)
	}

	return &result, nil
}

func (obj solarmanInverter) GetPlantRealtimeData(businessToken string, plantID int) (*GetRealtimePlantDataResponse, error) {
	url := obj.URL + "/station/v1.0/realTime?language=en"
	body := StationBody{
		StationID: plantID,
	}

	headers := map[string]string{
		AUTHORIZATION_HEADER: buildAuthorizationToken(businessToken),
	}

	req, err := requestor.PrepareHttpRequest(http.MethodPost, url, headers, body)
	if err != nil {
		return nil, err
	}

	data, status, err := prepareHttpResponse[APIGetRealtimePlantDataResponse](req)
	if err != nil {
		return nil, err
	}

	result := data.Result()
	if status != http.StatusOK {
		return nil, errors.NewError(status, result.Message)
	}

	return &result, nil
}

func (obj solarmanInverter) GetHistoricalPlantData(businessToken string, plantID int, timeType int, from, to int64) (*GetHistoricalPlantDataResponse, error) {
	startTime := buildDateFromTimestamp(from, timeType)
	endTime := buildDateFromTimestamp(to, timeType)

	body := GetHistoricalPlantDataBody{
		StationID: plantID,
		StartTime: startTime,
		EndTime:   endTime,
		TimeType:  timeType,
	}

	headers := map[string]string{
		AUTHORIZATION_HEADER: buildAuthorizationToken(businessToken),
	}

	url := obj.URL + "/station/v1.0/history?language=en"
	req, err := requestor.PrepareHttpRequest(http.MethodPost, url, headers, body)
	if err != nil {
		return nil, err
	}

	data, status, err := prepareHttpResponse[APIGetHistoricalPlantDataResponse](req)
	if err != nil {
		return nil, err
	}

	result := data.Result()
	if status != http.StatusOK {
		return nil, errors.NewError(status, result.Message)
	}

	return &result, nil
}

func (obj solarmanInverter) GetPlantDeviceListWithPagination(businessToken string, plantID int, page, size int) (*GetPlantDeviceListResponse, error) {
	body := GetPlantDeviceListBody{
		StationID: plantID,
		Page:      page,
		Size:      size,
	}

	headers := map[string]string{
		AUTHORIZATION_HEADER: buildAuthorizationToken(businessToken),
	}

	url := obj.URL + "/station/v1.0/device?language=en"
	req, err := requestor.PrepareHttpRequest(http.MethodPost, url, headers, body)
	if err != nil {
		return nil, err
	}

	data, status, err := prepareHttpResponse[APIGetPlantDeviceListResponse](req)
	if err != nil {
		return nil, err
	}

	result := data.Result()
	if status != http.StatusOK {
		return nil, errors.NewError(status, result.Message)
	}

	return &result, nil
}

func (obj solarmanInverter) GetPlantDeviceList(businessToken string, plantID int) ([]PlantDeviceItem, error) {
	devices := []PlantDeviceItem{}
	page := 1
	for {
		res, err := obj.GetPlantDeviceListWithPagination(businessToken, plantID, page, MAX_PAGE_SIZE)
		if err != nil {
			return nil, err
		}

		devices = append(devices, res.DeviceListItems...)
		page += 1

		if len(devices) >= res.Total {
			break
		}
	}

	return devices, nil
}

func (obj solarmanInverter) GetDeviceRealtimeData(businessToken string, deviceSN string) (*GetRealtimeDeviceDataResponse, error) {
	body := DeviceBody{
		DeviceSerialNumber: deviceSN,
	}

	headers := map[string]string{
		AUTHORIZATION_HEADER: buildAuthorizationToken(businessToken),
	}

	url := obj.URL + "/device/v1.0/currentData?language=en"
	req, err := requestor.PrepareHttpRequest(http.MethodPost, url, headers, body)
	if err != nil {
		return nil, err
	}

	data, status, err := prepareHttpResponse[APIGetRealtimeDeviceDataResponse](req)
	if err != nil {
		return nil, err
	}

	result := data.Result()
	if status != http.StatusOK {
		return nil, errors.NewError(status, result.Message)
	}

	return &result, nil
}

func (obj solarmanInverter) GetHistoricalDeviceData(businessToken string, deviceSN string, timeType int, from, to int64) (*GetHistoricalDeviceDataResponse, error) {
	startTime := buildDateFromTimestamp(from, timeType)
	endTime := buildDateFromTimestamp(to, timeType)

	body := GetHistoricalDeviceDataBody{
		DeviceSN:  deviceSN,
		StartTime: startTime,
		EndTime:   endTime,
		TimeType:  timeType,
	}

	headers := map[string]string{
		AUTHORIZATION_HEADER: buildAuthorizationToken(businessToken),
	}

	url := obj.URL + "/device/v1.0/historical?language=en"
	req, err := requestor.PrepareHttpRequest(http.MethodPost, url, headers, body)
	if err != nil {
		return nil, err
	}

	data, status, err := prepareHttpResponse[APIGetHistoricalDeviceDataResponse](req)
	if err != nil {
		return nil, err
	}

	result := data.Result()
	if status != http.StatusOK {
		return nil, errors.NewError(status, result.Message)
	}

	return &result, nil
}

func (obj solarmanInverter) GetDeviceAlertListWithPagination(businessToken string, deviceSN string, from, to int64, page, size int) (*GetDeviceAlertListResponse, error) {
	body := GetDeviceAlertListBody{
		DeviceSN:       deviceSN,
		StartTimestamp: from,
		EndTimestamp:   to,
		Page:           page,
		Size:           size,
	}

	headers := map[string]string{
		AUTHORIZATION_HEADER: buildAuthorizationToken(businessToken),
	}

	url := obj.URL + "/device/v1.0/alertList?language=en"
	req, err := requestor.PrepareHttpRequest(http.MethodPost, url, headers, body)
	if err != nil {
		return nil, err
	}

	data, status, err := prepareHttpResponse[APIGetDeviceAlertListResponse](req)
	if err != nil {
		return nil, err
	}

	result := data.Result()
	if status != http.StatusOK {
		return nil, errors.NewError(status, result.Message)
	}

	return &result, nil
}

func (obj solarmanInverter) GetDeviceAlertList(businessToken string, deviceSN string, from, to int64) ([]DeviceAlertItem, error) {
	result := []DeviceAlertItem{}
	page := 1
	for {
		res, err := obj.GetDeviceAlertListWithPagination(businessToken, deviceSN, from, to, page, MAX_PAGE_SIZE)
		if err != nil {
			return nil, err
		}

		result = append(result, res.AlertList...)
		page += 1

		if len(result) >= res.Total {
			break
		}
	}

	return result, nil
}
