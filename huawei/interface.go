package huawei

type HuaweiInverter interface {
	GetToken(username, password string) (string, error)
	GetPlantList() ([]PlantItem, error)
	GetPlantRealtimeData(plantCode string) ([]PlantRealtimeData, error)
	GetDailyPlantData(plantCode string, timestamp int64) ([]HistoricalPlantData, error)
	GetMonthlyPlantData(plantCode string, timestamp int64) ([]HistoricalPlantData, error)
	GetYearlyPlantData(plantCode string, timestamp int64) ([]HistoricalPlantData, error)
	GetDeviceList(plantCode string) ([]DeviceItem, error)
	GetDeviceRealtimeData(deviceID string, deviceTypeID string) ([]RealtimeDeviceData, error)
	GetDailyDeviceData(deviceID string, deviceTypeID string, timestamp int64) ([]HistoricalDeviceData, error)
	GetMonthlyDeviceData(deviceID string, deviceTypeID string, timestamp int64) ([]HistoricalDeviceData, error)
	GetYearlyDeviceData(deviceID string, deviceTypeID string, timestamp int64) ([]HistoricalDeviceData, error)
	GetDeviceAlarmList(plantCode string, from int64, to int64) ([]DeviceAlarmItem, error)
}
