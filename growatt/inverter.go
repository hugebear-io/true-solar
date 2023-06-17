package growatt

import (
	"net/http"
	"strings"
	"time"

	"github.com/hugebear-io/gofiber/requestor"
)

type growattInverter struct {
	URL      string
	token    string
	headers  map[string]string
	username string
}

func NewGrowattInverter(username, token string) *growattInverter {
	obj := &growattInverter{
		URL:      URL_VERSION1,
		token:    token,
		username: username,
		headers:  map[string]string{AUTH_HEADER: token},
	}

	return obj
}

func (obj growattInverter) GetPlantListWithPagination(page, size int) (*GetPlantListResponse, error) {
	queryMap := map[string]interface{}{
		"user_name": obj.username,
		"page":      page,
		"perpage":   size,
	}

	query := requestor.BuildQueryParams(queryMap)
	base := obj.URL + "/plant/user_plant_list" + query

	req, err := requestor.PrepareHttpRequest(http.MethodPost, base, obj.headers, nil)
	if err != nil {
		return nil, err
	}

	data, _, err := prepareHttpResponse[APIGetPlantListResponse](req)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	result := data.Result()
	return &result, nil
}

func (obj growattInverter) GetPlantList() ([]PlantItem, error) {
	plants := []PlantItem{}
	page := 1
	for {
		data, err := obj.GetPlantListWithPagination(page, MAX_PAGE_SIZE)
		if err != nil {
			return nil, err
		}

		plants = append(plants, data.Data.Plants...)

		if len(plants) >= data.Data.Count {
			break
		}

		page += 1
	}

	return plants, nil
}

func (obj growattInverter) GetPlantOverviewInfo(plantID int) (*PlantOverviewInfoData, error) {
	queryMap := map[string]interface{}{
		"plant_id": plantID,
	}

	query := requestor.BuildQueryParams(queryMap)
	base := obj.URL + "/plant/data" + query

	req, err := requestor.PrepareHttpRequest(http.MethodGet, base, obj.headers, nil)
	if err != nil {
		return nil, err
	}

	data, _, err := prepareHttpResponse[APIGetPlantOverviewInfoResponse](req)
	if err != nil {
		return nil, err
	}

	result := data.Result().Data
	return &result, nil
}

func (obj growattInverter) GetPlantDataLoggerInfo(plantID int) (*PlantDataLoggerInfoData, error) {
	queryMap := map[string]interface{}{
		"plant_id": plantID,
	}

	query := requestor.BuildQueryParams(queryMap)
	base := obj.URL + "/device/datalogger/list" + query

	req, err := requestor.PrepareHttpRequest(http.MethodGet, base, obj.headers, nil)
	if err != nil {
		return nil, err
	}

	data, _, err := prepareHttpResponse[APIGetPlantDataLoggerInfoResponse](req)
	if err != nil {
		return nil, err
	}

	result := data.Result().Data
	return &result, nil
}

func (obj growattInverter) GetPlantDeviceListWithPagination(plantID, page, size int) (*GetPlantDeviceListResponse, error) {
	queryMap := map[string]interface{}{
		"plant_id": plantID,
		"page":     page,
		"perpage":  size,
	}

	query := requestor.BuildQueryParams(queryMap)
	base := obj.URL + "/device/list" + query

	req, err := requestor.PrepareHttpRequest(http.MethodGet, base, obj.headers, nil)
	if err != nil {
		return nil, err
	}

	data, _, err := prepareHttpResponse[APIGetPlantDeviceListResponse](req)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	result := data.Result()
	return &result, nil
}

func (obj growattInverter) GetPlantDeviceList(plantID int) ([]DeviceItem, error) {
	devices := []DeviceItem{}
	page := 1
	for {
		data, err := obj.GetPlantDeviceListWithPagination(plantID, page, MAX_PAGE_SIZE)
		if err != nil {
			return nil, err
		}

		devices = append(devices, data.Data.Devices...)

		if len(devices) >= data.Data.Count {
			break
		}

		page += 1
	}

	return devices, nil
}

func (obj growattInverter) GetRealtimeDeviceData(deviceSN string) (*GetRealtimeDeviceDataResponse, error) {
	queryMap := map[string]interface{}{
		"device_sn": deviceSN,
	}

	query := requestor.BuildQueryParams(queryMap)
	base := obj.URL + "/device/inverter/last_new_data" + query

	req, err := requestor.PrepareHttpRequest(http.MethodGet, base, obj.headers, nil)
	if err != nil {
		return nil, err
	}

	data, _, err := prepareHttpResponse[APIGetRealtimeDeviceDataResponse](req)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	result := data.Result()
	return &result, nil
}

// TODO: GetRealtimeDeviceBatches
func (obj growattInverter) GetRealtimeDeviceBatchesWithPagination(deviceSNs []string, page int) (*GetRealtimeDeviceBatchesDataResponse, error) {
	queryMap := map[string]interface{}{
		"inverters": strings.Join(deviceSNs, ","),
		"pageNum":   page,
	}

	query := requestor.BuildQueryParams(queryMap)
	base := obj.URL + "/device/inverter/invs_data" + query

	req, err := requestor.PrepareHttpRequest(http.MethodPost, base, obj.headers, nil)
	if err != nil {
		return nil, err
	}

	data, _, err := prepareHttpResponse[APIGetRealtimeDeviceBatchesDataResponse](req)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	result := data.Result()
	return &result, nil
}

func (obj growattInverter) GetInverterAlertListWithPagination(deviceSN string, timestamp int64, page, size int) (*GetInverterAlertListData, error) {
	queryMap := map[string]interface{}{
		"device_sn": deviceSN,
		"date":      time.Unix(timestamp, 0).Format("2006-01-02"),
		"page":      page,
		"perpage":   size,
	}

	query := requestor.BuildQueryParams(queryMap)
	base := obj.URL + "/device/inverter/alarm" + query

	req, err := requestor.PrepareHttpRequest(http.MethodGet, base, obj.headers, nil)
	if err != nil {
		return nil, err
	}

	data, _, err := prepareHttpResponse[APIGetInverterAlertListResponse](req)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	result := data.Result().Data
	return &result, nil
}

func (obj growattInverter) GetInverterAlertList(deviceSN string, timestamp int64) ([]AlarmItem, error) {
	alarms := []AlarmItem{}
	page := 1
	for {
		data, err := obj.GetInverterAlertListWithPagination(deviceSN, timestamp, page, MAX_PAGE_SIZE)
		if err != nil {
			return nil, err
		}

		alarms = append(alarms, data.Alarms...)

		if len(alarms) >= data.Count {
			break
		}

		page += 1
	}

	return alarms, nil
}

func (obj growattInverter) GetEnergyStorageMachineAlertList(deviceSN string, timestamp int64) (*GetEnergyStorageMachineAlertListData, error) {
	queryMap := map[string]interface{}{
		"storage_sn": deviceSN,
		"date":       time.Unix(timestamp, 0).Format("2006-01-02"),
	}

	query := requestor.BuildQueryParams(queryMap)
	base := obj.URL + "/device/storage/alarm_data" + query

	req, err := requestor.PrepareHttpRequest(http.MethodPost, base, obj.headers, nil)
	if err != nil {
		return nil, err
	}

	data, _, err := prepareHttpResponse[APIGetEnergyStorageMachineAlertListResponse](req)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	result := data.Result().Data
	return &result, nil
}

func (obj growattInverter) GetMaxAlertListWithPagination(deviceSN string, timestamp int64, page, size int) (*GetMaxAlertListData, error) {
	queryMap := map[string]interface{}{
		"max_sn":  deviceSN,
		"date":    time.Unix(timestamp, 0).Format("2006-01-02"),
		"page":    page,
		"perpage": size,
	}

	query := requestor.BuildQueryParams(queryMap)
	base := obj.URL + "/device/max/alarm_data" + query

	req, err := requestor.PrepareHttpRequest(http.MethodPost, base, obj.headers, nil)
	if err != nil {
		return nil, err
	}

	data, _, err := prepareHttpResponse[APIGetMaxAlertListResponse](req)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	result := data.Result().Data
	return &result, nil
}

func (obj growattInverter) GetMaxAlertList(deviceSN string, timestamp int64) ([]AlarmItem, error) {
	alarms := []AlarmItem{}
	page := 1
	for {
		data, err := obj.GetMaxAlertListWithPagination(deviceSN, timestamp, page, MAX_PAGE_SIZE)
		if err != nil {
			return nil, err
		}

		alarms = append(alarms, data.Alarms...)
		if len(alarms) >= data.Count {
			break
		}
		page += 1
	}

	return alarms, nil
}

func (obj growattInverter) GetMixAlertListWithPagination(deviceSN string, timestamp int64, page, size int) (*GetMixAlertListData, error) {
	queryMap := map[string]interface{}{
		"mix_sn":  deviceSN,
		"date":    time.Unix(timestamp, 0).Format("2006-01-02"),
		"page":    page,
		"perpage": size,
	}

	query := requestor.BuildQueryParams(queryMap)
	base := obj.URL + "/device/mix/alarm_data" + query

	req, err := requestor.PrepareHttpRequest(http.MethodPost, base, obj.headers, nil)
	if err != nil {
		return nil, err
	}

	data, _, err := prepareHttpResponse[APIGetMixAlertListResponse](req)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	result := data.Result().Data
	return &result, nil
}

func (obj growattInverter) GetMixAlertList(deviceSN string, timestamp int64) ([]AlarmItem, error) {
	alarms := []AlarmItem{}
	page := 1
	for {
		data, err := obj.GetMixAlertListWithPagination(deviceSN, timestamp, page, MAX_PAGE_SIZE)
		if err != nil {
			return nil, err
		}

		alarms = append(alarms, data.Alarms...)
		if len(alarms) >= data.Count {
			break
		}
		page += 1
	}

	return alarms, nil
}

func (obj growattInverter) GetSpaAlertListWithPagination(deviceSN string, timestamp int64, page, size int) (*GetSpaAlertListData, error) {
	queryMap := map[string]interface{}{
		"spa_sn":  deviceSN,
		"date":    time.Unix(timestamp, 0).Format("2006-01-02"),
		"page":    page,
		"perpage": size,
	}

	query := requestor.BuildQueryParams(queryMap)
	base := obj.URL + "/device/spa/alarm_data" + query

	req, err := requestor.PrepareHttpRequest(http.MethodPost, base, obj.headers, nil)
	if err != nil {
		return nil, err
	}

	data, _, err := prepareHttpResponse[APIGetSpaAlertListResponse](req)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	result := data.Result().Data
	return &result, nil
}

func (obj growattInverter) GetSpaAlertList(deviceSN string, timestamp int64) ([]AlarmItem, error) {
	alarms := []AlarmItem{}
	page := 1
	for {
		data, err := obj.GetSpaAlertListWithPagination(deviceSN, timestamp, page, MAX_PAGE_SIZE)
		if err != nil {
			return nil, err
		}

		alarms = append(alarms, data.Alarms...)
		if len(alarms) >= data.Count {
			break
		}
		page += 1
	}

	return alarms, nil
}

func (obj growattInverter) GetMinAlertListWithPagination(deviceSN string, timestamp int64, page, size int) (*GetMinAlertListData, error) {
	queryMap := map[string]interface{}{
		"tlx_sn":  deviceSN,
		"date":    time.Unix(timestamp, 0).Format("2006-01-02"),
		"page":    page,
		"perpage": size,
	}

	query := requestor.BuildQueryParams(queryMap)
	base := obj.URL + "/device/tlx/alarm_data" + query

	req, err := requestor.PrepareHttpRequest(http.MethodPost, base, obj.headers, nil)
	if err != nil {
		return nil, err
	}

	data, _, err := prepareHttpResponse[APIGetMinAlertListResponse](req)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	result := data.Result().Data
	return &result, nil
}

func (obj growattInverter) GetMinAlertList(deviceSN string, timestamp int64) ([]AlarmItem, error) {
	alarms := []AlarmItem{}
	page := 1
	for {
		data, err := obj.GetMinAlertListWithPagination(deviceSN, timestamp, page, MAX_PAGE_SIZE)
		if err != nil {
			return nil, err
		}

		alarms = append(alarms, data.Alarms...)
		if len(alarms) >= data.Count {
			break
		}
		page += 1
	}

	return alarms, nil
}

func (obj growattInverter) GetPcsAlertListWithPagination(deviceSN string, timestamp int64, page, size int) (*GetPcsAlertListData, error) {
	queryMap := map[string]interface{}{
		"pcs_sn":  deviceSN,
		"date":    time.Unix(timestamp, 0).Format("2006-01-02"),
		"page":    page,
		"perpage": size,
	}

	query := requestor.BuildQueryParams(queryMap)
	base := obj.URL + "/device/pcs/alarm_data" + query

	req, err := requestor.PrepareHttpRequest(http.MethodPost, base, obj.headers, nil)
	if err != nil {
		return nil, err
	}

	data, _, err := prepareHttpResponse[APIGetPcsAlertListResponse](req)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	result := data.Result().Data
	return &result, nil
}

func (obj growattInverter) GetPcsAlertList(deviceSN string, timestamp int64) ([]AlarmItem, error) {
	alarms := []AlarmItem{}
	page := 1
	for {
		data, err := obj.GetPcsAlertListWithPagination(deviceSN, timestamp, page, MAX_PAGE_SIZE)
		if err != nil {
			return nil, err
		}

		alarms = append(alarms, data.Alarms...)
		if len(alarms) >= data.Count {
			break
		}
		page += 1
	}

	return alarms, nil
}

func (obj growattInverter) GetHpsAlertListWithPagination(deviceSN string, timestamp int64, page, size int) (*GetHpsAlertListData, error) {
	queryMap := map[string]interface{}{
		"hps_sn":  deviceSN,
		"date":    time.Unix(timestamp, 0).Format("2006-01-02"),
		"page":    page,
		"perpage": size,
	}

	query := requestor.BuildQueryParams(queryMap)
	base := obj.URL + "/device/hps/alarm_data" + query

	req, err := requestor.PrepareHttpRequest(http.MethodPost, base, obj.headers, nil)
	if err != nil {
		return nil, err
	}

	data, _, err := prepareHttpResponse[APIGetHpsAlertListResponse](req)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	result := data.Result().Data
	return &result, nil
}

func (obj growattInverter) GetHpsAlertList(deviceSN string, timestamp int64) ([]AlarmItem, error) {
	alarms := []AlarmItem{}
	page := 1
	for {
		data, err := obj.GetHpsAlertListWithPagination(deviceSN, timestamp, page, MAX_PAGE_SIZE)
		if err != nil {
			return nil, err
		}

		alarms = append(alarms, data.Alarms...)
		if len(alarms) >= data.Count {
			break
		}
		page += 1
	}

	return alarms, nil
}

func (obj growattInverter) GetPbdAlertListWithPagination(deviceSN string, timestamp int64, page, size int) (*GetPbdAlertListData, error) {
	queryMap := map[string]interface{}{
		"pbd_sn":  deviceSN,
		"date":    time.Unix(timestamp, 0).Format("2006-01-02"),
		"page":    page,
		"perpage": size,
	}

	query := requestor.BuildQueryParams(queryMap)
	base := obj.URL + "/device/pbd/alarm_data" + query

	req, err := requestor.PrepareHttpRequest(http.MethodPost, base, obj.headers, nil)
	if err != nil {
		return nil, err
	}

	data, _, err := prepareHttpResponse[APIGetPbdAlertListResponse](req)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	result := data.Result().Data
	return &result, nil
}

func (obj growattInverter) GetPbdAlertList(deviceSN string, timestamp int64) ([]AlarmItem, error) {
	alarms := []AlarmItem{}
	page := 1
	for {
		data, err := obj.GetPbdAlertListWithPagination(deviceSN, timestamp, page, MAX_PAGE_SIZE)
		if err != nil {
			return nil, err
		}

		alarms = append(alarms, data.Alarms...)
		if len(alarms) >= data.Count {
			break
		}
		page += 1
	}

	return alarms, nil
}
