package kstar

import (
	"github.com/openlyinc/pointy"
)

// ===== DEFAULT RESPONSE =====
type APIMeta struct {
	Success *bool   `json:"success,omitempty"`
	Code    *string `json:"code,omitempty"`
}

func (obj APIMeta) Result() Meta {
	return Meta{
		Success: pointy.BoolValue(obj.Success, false),
		Code:    pointy.StringValue(obj.Code, ""),
	}
}

type APIDefaultResponse struct {
	Meta *APIMeta `json:"meta,omitempty"`
}

func (obj APIDefaultResponse) Result() DefaultResponse {
	var meta *Meta
	if obj.Meta != nil {
		meta = pointy.Pointer(obj.Meta.Result())
	}

	return DefaultResponse{
		Meta: meta,
	}
}

// ===========================
// API Fetch: GetPlantList
// Result Model: PlantItem
// ===========================
type APIGetPlantListResponse struct {
	Meta *APIMeta       `json:"meta,omitempty"`
	Data []APIPlantItem `json:"data,omitempty"`
}

type APIPlantItem struct {
	ID                *string  `json:"powerId,omitempty"`
	Name              *string  `json:"powerName,omitempty"`
	InstalledCapacity *float64 `json:"powerCap,omitempty"`
	Longitude         *float64 `json:"longitude,omitempty"`
	Latitude          *float64 `json:"latitude,omitempty"`
	CityCode          *string  `json:"cityCode,omitempty"`
	CityName          *string  `json:"cityName,omitempty"`
	DistrictCode      *string  `json:"districtCode,omitempty"`
	Address           *string  `json:"powerArea,omitempty"`
	CreatedTime       *string  `json:"createTime,omitempty"`
	DealerCode        *string  `json:"dealerCode,omitempty"`
	ElectricPrice     *float64 `json:"elecPrice,omitempty"`
	ElectricUnit      *string  `json:"elecUnit,omitempty"`
	TimeZone          *string  `json:"timeZone,omitempty"`
}

func (obj APIPlantItem) Result() PlantItem {
	return PlantItem{
		ID:                pointy.StringValue(obj.ID, ""),
		Name:              pointy.StringValue(obj.Name, ""),
		InstalledCapacity: pointy.Float64Value(obj.InstalledCapacity, 0),
		Longitude:         pointy.Float64Value(obj.Longitude, 0),
		Latitude:          pointy.Float64Value(obj.Latitude, 0),
		CityCode:          pointy.StringValue(obj.CityCode, ""),
		CityName:          pointy.StringValue(obj.CityName, ""),
		DistrictCode:      pointy.StringValue(obj.DistrictCode, ""),
		Address:           pointy.StringValue(obj.Address, ""),
		CreatedTime:       pointy.StringValue(obj.CreatedTime, ""),
		DealerCode:        pointy.StringValue(obj.DealerCode, ""),
		ElectricPrice:     pointy.Float64Value(obj.ElectricPrice, 0),
		ElectricUnit:      pointy.StringValue(obj.ElectricUnit, ""),
		TimeZone:          pointy.StringValue(obj.TimeZone, ""),
	}
}

// ===========================
// API Fetch: GetRealtimeDeviceData
// Result Model: DeviceInfo
// ===========================
type APIGetRealtimeDeviceDataResponse struct {
	Meta *APIMeta       `json:"meta,omitempty"`
	Data *APIDeviceInfo `json:"data,omitempty"`
}

type APIDeviceInfo struct {
	ID                 *string  `json:"device_id,omitempty"`
	InverterID         *string  `json:"inverter_id,omitempty"`
	Status             *int     `json:"status,omitempty"`
	Version            *int     `json:"version,omitempty"`
	SaveTime           *string  `json:"save_time,omitempty"`
	VoltagePV1         *float64 `json:"voltage_pv1,omitempty"`
	VoltagePV2         *float64 `json:"voltage_pv2,omitempty"`
	VoltagePV3         *float64 `json:"voltage_pv3,omitempty"`
	CurrentPV1         *float64 `json:"current_pv1,omitempty"`
	CurrentPV2         *float64 `json:"current_pv2,omitempty"`
	CurrentPV3         *float64 `json:"current_pv3,omitempty"`
	PowerPV1           *float64 `json:"power_pv1,omitempty"`
	PowerPV2           *float64 `json:"power_pv2,omitempty"`
	PowerPV3           *float64 `json:"power_pv3,omitempty"`
	VoltagePBUS        *float64 `json:"voltage_pbus,omitempty"`
	VoltageNBUS        *float64 `json:"voltage_nbus,omitempty"`
	VoltageRS          *float64 `json:"voltage_rs,omitempty"`
	VoltageST          *float64 `json:"voltage_st,omitempty"`
	VoltageTR          *float64 `json:"voltage_tr,omitempty"`
	FrequencyRS        *float64 `json:"frequency_rs,omitempty"`
	FrequencyST        *float64 `json:"frequency_st,omitempty"`
	FrequencyTR        *float64 `json:"frequency_tr,omitempty"`
	CurrentR           *float64 `json:"current_r,omitempty"`
	CurrentS           *float64 `json:"current_s,omitempty"`
	CurrentT           *float64 `json:"current_t,omitempty"`
	PowerInter         *float64 `json:"power_inter,omitempty"`
	RadiatorTemp       *float64 `json:"radiator_temp,omitempty"`
	ModuleTemp         *float64 `json:"module_temp,omitempty"`
	DSPAlarmCode       *string  `json:"dsp_alarm_code,omitempty"`
	DSPErrorCode       *string  `json:"dsp_error_code,omitempty"`
	InverterType       *int     `json:"inverter_type,omitempty"`
	DeviceModel        *int     `json:"device_model,omitempty"`
	FanA               *float64 `json:"fan_A,omitempty"`
	FanB               *float64 `json:"fan_B,omitempty"`
	FanC               *float64 `json:"fan_C,omitempty"`
	ARMAlarmCode       *string  `json:"arm_alarm_code,omitempty"`
	ARMErrorCode       *string  `json:"arm_error_code,omitempty"`
	InputType          *int     `json:"input_type,omitempty"`
	GenerationStandard *int     `json:"generation_standard,omitempty"`
	TotalGeneration    *float64 `json:"total_generation,omitempty"`
	YearGeneration     *float64 `json:"year_generation,omitempty"`
	DayGeneration      *float64 `json:"day_generation,omitempty"`
	MonthGeneration    *float64 `json:"month_generation,omitempty"`
	VoltageOpen        *float64 `json:"voltage_open,omitempty"`
	DelayTime          *float64 `json:"delay_time,omitempty"`
	VoltageLower       *float64 `json:"voltage_lower,omitempty"`
	VoltageCeiling     *float64 `json:"voltage_ceiling,omitempty"`
	PowerLower         *float64 `json:"power_lower,omitempty"`
	PowerCeiling       *float64 `json:"power_ceiling,omitempty"`
	ReactSet           *float64 `json:"react_set,omitempty"`
	ActiveSet          *float64 `json:"active_set,omitempty"`
	ReactiveSet        *float64 `json:"reactive_set,omitempty"`
	ReactiveType       *int     `json:"reactive_type,omitempty"`
	PowerApparent      *float64 `json:"power_apparent,omitempty"`
	PowerReactive      *float64 `json:"power_reactive,omitempty"`
	PowerFactors       *float64 `json:"power_factors,omitempty"`
	CurrentImpedance   *float64 `json:"current_impedance,omitempty"`
	OverPower          *int     `json:"over_power,omitempty"`
	OverPowerSet       *float64 `json:"over_power_set,omitempty"`
	QVHighSet          *float64 `json:"qv_high_set,omitempty"`
	QVHighPercentage   *float64 `json:"qv_high_percentage,omitempty"`
	QVLowerSet         *float64 `json:"qv_lower_set,omitempty"`
	QVLowerPercentage  *float64 `json:"qv_lower_percentage,omitempty"`
	WeakGeneration     *float64 `json:"weak_generation,omitempty"`
	InverterPower      *float64 `json:"inverter_power,omitempty"`
	InverterMake       *int     `json:"inverter_make,omitempty"`
	DREDMake           *int     `json:"dred_make,omitempty"`
}

func (obj APIDeviceInfo) Result() DeviceInfo {
	return DeviceInfo{
		ID:                 pointy.StringValue(obj.ID, ""),
		InverterID:         pointy.StringValue(obj.InverterID, ""),
		Status:             pointy.IntValue(obj.Status, 0),
		Version:            pointy.IntValue(obj.Version, 0),
		SaveTime:           pointy.StringValue(obj.SaveTime, ""),
		VoltagePV1:         pointy.Float64Value(obj.VoltagePV1, 0),
		VoltagePV2:         pointy.Float64Value(obj.VoltagePV2, 0),
		VoltagePV3:         pointy.Float64Value(obj.VoltagePV3, 0),
		CurrentPV1:         pointy.Float64Value(obj.CurrentPV1, 0),
		CurrentPV2:         pointy.Float64Value(obj.CurrentPV2, 0),
		CurrentPV3:         pointy.Float64Value(obj.CurrentPV3, 0),
		PowerPV1:           pointy.Float64Value(obj.PowerPV1, 0),
		PowerPV2:           pointy.Float64Value(obj.PowerPV2, 0),
		PowerPV3:           pointy.Float64Value(obj.PowerPV3, 0),
		VoltagePBUS:        pointy.Float64Value(obj.VoltagePBUS, 0),
		VoltageNBUS:        pointy.Float64Value(obj.VoltageNBUS, 0),
		VoltageRS:          pointy.Float64Value(obj.VoltageRS, 0),
		VoltageST:          pointy.Float64Value(obj.VoltageST, 0),
		VoltageTR:          pointy.Float64Value(obj.VoltageTR, 0),
		FrequencyRS:        pointy.Float64Value(obj.FrequencyRS, 0),
		FrequencyST:        pointy.Float64Value(obj.FrequencyST, 0),
		FrequencyTR:        pointy.Float64Value(obj.FrequencyTR, 0),
		CurrentR:           pointy.Float64Value(obj.CurrentR, 0),
		CurrentS:           pointy.Float64Value(obj.CurrentS, 0),
		CurrentT:           pointy.Float64Value(obj.CurrentT, 0),
		PowerInter:         pointy.Float64Value(obj.PowerInter, 0),
		RadiatorTemp:       pointy.Float64Value(obj.RadiatorTemp, 0),
		ModuleTemp:         pointy.Float64Value(obj.ModuleTemp, 0),
		DSPAlarmCode:       pointy.StringValue(obj.DSPAlarmCode, ""),
		DSPErrorCode:       pointy.StringValue(obj.DSPErrorCode, ""),
		InverterType:       pointy.IntValue(obj.InverterType, 0),
		DeviceModel:        pointy.IntValue(obj.DeviceModel, 0),
		FanA:               pointy.Float64Value(obj.FanA, 0),
		FanB:               pointy.Float64Value(obj.FanB, 0),
		FanC:               pointy.Float64Value(obj.FanC, 0),
		ARMAlarmCode:       pointy.StringValue(obj.ARMAlarmCode, ""),
		ARMErrorCode:       pointy.StringValue(obj.ARMErrorCode, ""),
		InputType:          pointy.IntValue(obj.InputType, 0),
		GenerationStandard: pointy.IntValue(obj.GenerationStandard, 0),
		TotalGeneration:    pointy.Float64Value(obj.TotalGeneration, 0),
		YearGeneration:     pointy.Float64Value(obj.YearGeneration, 0),
		DayGeneration:      pointy.Float64Value(obj.DayGeneration, 0),
		MonthGeneration:    pointy.Float64Value(obj.MonthGeneration, 0),
		VoltageOpen:        pointy.Float64Value(obj.VoltageOpen, 0),
		DelayTime:          pointy.Float64Value(obj.DelayTime, 0),
		VoltageLower:       pointy.Float64Value(obj.VoltageLower, 0),
		VoltageCeiling:     pointy.Float64Value(obj.VoltageCeiling, 0),
		PowerLower:         pointy.Float64Value(obj.PowerLower, 0),
		PowerCeiling:       pointy.Float64Value(obj.PowerCeiling, 0),
		ReactSet:           pointy.Float64Value(obj.ReactSet, 0),
		ActiveSet:          pointy.Float64Value(obj.ActiveSet, 0),
		ReactiveSet:        pointy.Float64Value(obj.ReactiveSet, 0),
		ReactiveType:       pointy.IntValue(obj.ReactiveType, 0),
		PowerApparent:      pointy.Float64Value(obj.PowerApparent, 0),
		PowerReactive:      pointy.Float64Value(obj.PowerReactive, 0),
		PowerFactors:       pointy.Float64Value(obj.PowerFactors, 0),
		CurrentImpedance:   pointy.Float64Value(obj.CurrentImpedance, 0),
		OverPower:          pointy.IntValue(obj.OverPower, 0),
		OverPowerSet:       pointy.Float64Value(obj.OverPowerSet, 0),
		QVHighSet:          pointy.Float64Value(obj.QVHighSet, 0),
		QVHighPercentage:   pointy.Float64Value(obj.QVHighPercentage, 0),
		QVLowerSet:         pointy.Float64Value(obj.QVLowerSet, 0),
		QVLowerPercentage:  pointy.Float64Value(obj.QVLowerPercentage, 0),
		WeakGeneration:     pointy.Float64Value(obj.WeakGeneration, 0),
		InverterPower:      pointy.Float64Value(obj.InverterPower, 0),
		InverterMake:       pointy.IntValue(obj.InverterMake, 0),
		DREDMake:           pointy.IntValue(obj.DREDMake, 0),
	}
}

// ===========================
// API Fetch: GetDeviceList
// Result Model: DeviceItem
// ===========================
type APIGetDeviceListResponse struct {
	Meta *APIMeta       `json:"meta,omitempty"`
	Data *APIDeviceData `json:"data,omitempty"`
}

type APIDeviceData struct {
	Code  *int            `json:"code,omitempty"`
	Count *int            `json:"count,omitempty"`
	List  []APIDeviceItem `json:"list,omitempty"`
}

func (obj APIDeviceData) Result() DeviceData {
	deviceItems := []DeviceItem{}
	for _, item := range obj.List {
		deviceItems = append(deviceItems, item.Result())
	}

	return DeviceData{
		Code:  pointy.IntValue(obj.Code, 0),
		Count: pointy.IntValue(obj.Count, 0),
		List:  deviceItems,
	}
}

type APIDeviceItem struct {
	ID         *string `json:"deviceId,omitempty"`
	InverterID *string `json:"inverterId,omitempty"`
	Name       *string `json:"deviceName,omitempty"`
	Status     *int    `json:"deviceStatus,omitempty"`
	PlantID    *string `json:"powerId,omitempty"`
	PlantName  *string `json:"powerName,omitempty"`
	SaveTime   *string `json:"saveTime,omitempty"`
}

func (obj APIDeviceItem) Result() DeviceItem {
	return DeviceItem{
		ID:         pointy.StringValue(obj.ID, ""),
		InverterID: pointy.StringValue(obj.InverterID, ""),
		Name:       pointy.StringValue(obj.Name, ""),
		Status:     pointy.IntValue(obj.Status, 0),
		PlantID:    pointy.StringValue(obj.PlantID, ""),
		PlantName:  pointy.StringValue(obj.PlantName, ""),
		SaveTime:   pointy.StringValue(obj.SaveTime, ""),
	}
}

// ===========================
// API Fetch: GetDeviceAlarmList, GetRealtimeAlarmListOfDevice
// Result Model: DeviceAlarm
// ===========================
type APIGetRealtimeAlarmListOfDeviceResponse struct {
	Meta *APIMeta             `json:"meta,omitempty"`
	Data []APIDeviceAlarmItem `json:"data,omitempty"`
}

type APIGetDeviceAlarmListResponse struct {
	Meta *APIMeta            `json:"meta,omitempty"`
	Data *APIDeviceAlarmData `json:"data,omitempty"`
}

type APIDeviceAlarmData struct {
	Code  *int                 `json:"code,omitempty"`
	Count *int                 `json:"count,omitempty"`
	Data  []APIDeviceAlarmItem `json:"data,omitempty"`
}

func (obj APIDeviceAlarmData) Result() DeviceAlarmData {
	deviceAlarmItems := []DeviceAlarmItem{}
	for _, item := range obj.Data {
		deviceAlarmItems = append(deviceAlarmItems, item.Result())
	}

	return DeviceAlarmData{
		Code:  pointy.IntValue(obj.Code, 0),
		Count: pointy.IntValue(obj.Count, 0),
		Data:  deviceAlarmItems,
	}
}

type APIDeviceAlarmItem struct {
	DeviceID   *string `json:"deviceId,omitempty"`
	DeviceName *string `json:"deviceName,omitempty"`
	SaveTime   *string `json:"saveTime,omitempty"`
	Message    *string `json:"message,omitempty"`
	ErrorLevel *int    `json:"errorLevel,omitempty"`
	RemoveTime *string `json:"removeTime,omitempty"`
	PlantID    *string `json:"powerId,omitempty"`
	PlantName  *string `json:"powerName,omitempty"`
}

func (obj APIDeviceAlarmItem) Result() DeviceAlarmItem {
	return DeviceAlarmItem{
		DeviceID:   pointy.StringValue(obj.DeviceID, ""),
		DeviceName: pointy.StringValue(obj.DeviceName, ""),
		SaveTime:   pointy.StringValue(obj.SaveTime, ""),
		Message:    pointy.StringValue(obj.Message, ""),
		ErrorLevel: pointy.IntValue(obj.ErrorLevel, 0),
		RemoveTime: pointy.StringValue(obj.RemoveTime, ""),
		PlantID:    pointy.StringValue(obj.PlantID, ""),
		PlantName:  pointy.StringValue(obj.PlantName, ""),
	}
}
