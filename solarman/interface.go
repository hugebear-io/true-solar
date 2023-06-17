package solarman

type SolarmanInverter interface {
	GetBasicToken() (*GetTokenResponse, error)
	GetBusinessToken(orgID int) (*GetTokenResponse, error)
	GetUserInfo() (*GetUserInfoResponse, error)
	GetPlantListWithPagination(businessToken string, page, size int) (*GetPlantListResponse, error)
	GetPlantList(businessToken string) ([]PlantItem, error)
	GetPlantBaseInfo(businessToken string, plantID int) (*GetPlantBaseInfoResponse, error)
	GetPlantRealtimeData(businessToken string, plantID int) (*GetRealtimePlantDataResponse, error)
	GetHistoricalPlantData(businessToken string, plantID int, timeType int, from, to int64) (*GetHistoricalPlantDataResponse, error)
	GetPlantDeviceListWithPagination(businessToken string, plantID int, page, size int) (*GetPlantDeviceListResponse, error)
	GetPlantDeviceList(businessToken string, plantID int) ([]PlantDeviceItem, error)
	GetDeviceRealtimeData(businessToken string, deviceSN string) (*GetRealtimeDeviceDataResponse, error)
	GetHistoricalDeviceData(businessToken string, deviceSN string, timeType int, from, to int64) (*GetHistoricalDeviceDataResponse, error)
	GetDeviceAlertListWithPagination(businessToken string, deviceSN string, from, to int64, page, size int) (*GetDeviceAlertListResponse, error)
	GetDeviceAlertList(businessToken string, deviceSN string, from, to int64) ([]DeviceAlertItem, error)
}
