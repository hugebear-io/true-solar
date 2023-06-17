package kstar

type KStarInverter interface {
	EncryptParameter(data map[string]string) string
	GetPlantList() ([]PlantItem, error)
	GetDeviceListWithPagination(page, size int) (*DeviceData, error)
	GetDeviceList() ([]DeviceItem, error)
	GetRealtimeDeviceData(deviceID string) (*DeviceInfo, error)
	GetRealtimeAlarmListOfDevice(deviceID string) ([]DeviceAlarmItem, error)
	GetRealtimeAlarmListOfPlantWithPagination(from, to int64, page, size int) (*DeviceAlarmData, error)
	GetRealtimeAlarmListOfPlant(from, to int64) ([]DeviceAlarmItem, error)
}
