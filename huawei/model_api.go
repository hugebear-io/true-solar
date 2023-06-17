package huawei

import "github.com/openlyinc/pointy"

// ===== DEFAULT RESPONSE =====
type APIDefaultResponse struct {
	Success    *bool        `json:"success,omitempty"`
	FailCode   *int         `json:"failCode,omitempty"`
	Parameters *interface{} `json:"params,omitempty"`
	Message    *string      `json:"message,omitempty"`
}

func (obj APIDefaultResponse) Result() DefaultResponse {
	return DefaultResponse{
		Success:    pointy.BoolValue(obj.Success, false),
		FailCode:   pointy.IntValue(obj.FailCode, 0),
		Parameters: pointy.PointerValue(obj.Parameters, nil),
		Message:    pointy.StringValue(obj.Message, ""),
	}
}

// ===========================
// API Fetch: GetToken
// Result Model: -
// ===========================
type APIGetTokenResponse struct {
	APIDefaultResponse
	Data *string `jsoN:"data,omitempty"`
}

// ===========================
// API Fetch: GetPlantList
// Result Model: PlantItem
// ===========================
type APIGetPlantListResponse struct {
	APIDefaultResponse
	Data []APIPlantItem `json:"data,omitempty"`
}

type APIPlantItem struct {
	Code           *string  `json:"stationCode,omitempty"`
	Name           *string  `json:"stationName,omitempty"`
	Address        *string  `json:"stationAddr,omitempty"`
	Capacity       *float64 `json:"capacity,omitempty"`
	BuildState     *string  `json:"buildState,omitempty"`
	CombineType    *string  `json:"combineType,omitempty"`
	AIDType        *int     `json:"aidType,omitempty"`
	StationLinkMan *string  `json:"stationLinkman,omitempty"`
	LinkManPhone   *string  `json:"linkmanPho,omitempty"`
}

func (obj APIPlantItem) Result() PlantItem {
	return PlantItem{
		Code:           pointy.StringValue(obj.Code, ""),
		Name:           pointy.StringValue(obj.Name, ""),
		Address:        pointy.StringValue(obj.Address, ""),
		Capacity:       pointy.Float64Value(obj.Capacity, 0),
		BuildState:     pointy.StringValue(obj.BuildState, ""),
		CombineType:    pointy.StringValue(obj.CombineType, ""),
		AIDType:        pointy.IntValue(obj.AIDType, 0),
		StationLinkMan: pointy.StringValue(obj.StationLinkMan, ""),
		LinkManPhone:   pointy.StringValue(obj.LinkManPhone, ""),
	}
}

// ===========================
// API Fetch: GetPlantRealtimeData
// Result Model: PlantRealtimeData, PlantRealtimeItem
// ===========================
type APIGetPlantRealtimeDataResponse struct {
	APIDefaultResponse
	Data []APIPlantRealtimeData `json:"data,omitempty"`
}

type APIPlantRealtimeData struct {
	Code        *string               `json:"stationCode,omitempty"`
	DataItemMap *APIPlantRealtimeItem `json:"dataItemMap,omitempty"`
}

func (obj APIPlantRealtimeData) Result() PlantRealtimeData {
	return PlantRealtimeData{
		Code:        pointy.StringValue(obj.Code, ""),
		DataItemMap: obj.DataItemMap.Result(),
	}
}

type APIPlantRealtimeItem struct {
	TotalIncome     *float64 `json:"total_income,omitempty"`
	TotalPower      *float64 `json:"total_power,omitempty"`
	DayPower        *float64 `json:"day_power,omitempty"`
	DayIncome       *float64 `json:"day_income,omitempty"`
	RealHealthState *int     `json:"real_health_state,omitempty"`
	MonthPower      *float64 `json:"month_power,omitempty"`
}

func (obj APIPlantRealtimeItem) Result() PlantRealtimeItem {
	return PlantRealtimeItem{
		TotalIncome:     pointy.Float64Value(obj.TotalIncome, 0),
		TotalPower:      pointy.Float64Value(obj.TotalPower, 0),
		DayPower:        pointy.Float64Value(obj.DayPower, 0),
		DayIncome:       pointy.Float64Value(obj.DayIncome, 0),
		RealHealthState: pointy.IntValue(obj.RealHealthState, 0),
		MonthPower:      pointy.Float64Value(obj.MonthPower, 0),
	}
}

// ===========================
// API Fetch: GetDailyPlantData, GetMonthlyPlantData, GetYearlyPlantData
// Result Model: HistoricalPlantData, HistoricalPlantItem
// ===========================
type APIGetHistoricalPlantDataResponse struct {
	APIDefaultResponse
	Data []APIHistoricalPlantData `json:"data,omitempty"`
}

type APIHistoricalPlantData struct {
	Code        *string                 `json:"stationCode,omitempty"`
	CollectTime *int64                  `json:"collectTime,omitempty"`
	DataItemMap *APIHistoricalPlantItem `json:"dataItemMap,omitempty"`
}

func (obj APIHistoricalPlantData) Result() HistoricalPlantData {
	return HistoricalPlantData{
		Code:        pointy.StringValue(obj.Code, ""),
		CollectTime: pointy.Int64Value(obj.CollectTime, 0),
		DataItemMap: obj.DataItemMap.Result(),
	}
}

type APIHistoricalPlantItem struct {
	RadiationIntensity *float64 `json:"radiation_intensity,omitempty"`
	InstalledCapacity  *float64 `json:"installed_capacity,omitempty"`
	UsePower           *float64 `json:"use_power,omitempty"`
	InverterPower      *float64 `json:"inverter_power,omitempty"`
	PowerProfit        *float64 `json:"power_profit,omitempty"`
	TheoryPower        *float64 `json:"theory_power,omitempty"`
	PerPowerRatio      *float64 `json:"perpower_ratio,omitempty"`
	OnGridPower        *float64 `json:"ongrid_power,omitempty"`
	PerformanceRatio   *float64 `json:"performance_ratio,omitempty"`
	ReductionTotalCO2  *float64 `json:"reduction_total_co2,omitempty"`
	ReductionTotalCoal *float64 `json:"reduction_total_coal,omitempty"`
	ReductionTotalTree *float64 `json:"reduction_total_tree,omitempty"`
}

func (obj APIHistoricalPlantItem) Result() HistoricalPlantItem {
	return HistoricalPlantItem{
		RadiationIntensity: pointy.Float64Value(obj.RadiationIntensity, 0),
		InstalledCapacity:  pointy.Float64Value(obj.InstalledCapacity, 0),
		UsePower:           pointy.Float64Value(obj.UsePower, 0),
		InverterPower:      pointy.Float64Value(obj.InverterPower, 0),
		PowerProfit:        pointy.Float64Value(obj.PowerProfit, 0),
		TheoryPower:        pointy.Float64Value(obj.TheoryPower, 0),
		PerPowerRatio:      pointy.Float64Value(obj.PerPowerRatio, 0),
		OnGridPower:        pointy.Float64Value(obj.OnGridPower, 0),
		PerformanceRatio:   pointy.Float64Value(obj.PerformanceRatio, 0),
		ReductionTotalCO2:  pointy.Float64Value(obj.ReductionTotalCO2, 0),
		ReductionTotalCoal: pointy.Float64Value(obj.ReductionTotalCoal, 0),
		ReductionTotalTree: pointy.Float64Value(obj.ReductionTotalTree, 0),
	}
}

// ===========================
// API Fetch: GetDeviceList
// Result Model: DeviceItem
// ===========================
type APIGetDeviceListResponse struct {
	APIDefaultResponse
	Data []APIDeviceItem `json:"data,omitempty"`
}

type APIDeviceItem struct {
	ID              *int     `json:"id,omitempty"`
	SerialNumber    *string  `json:"esnCode,omitempty"`
	Name            *string  `json:"devName,omitempty"`
	TypeID          *int     `json:"devTypeId,omitempty"`
	InverterModel   *string  `json:"invType,omitempty"`
	Latitude        *float64 `json:"latitude,omitempty"`
	Longitude       *float64 `json:"longitude,omitempty"`
	SoftwareVersion *string  `json:"softwareVersion,omitempty"`
	PlantCode       *string  `json:"stationCode,omitempty"`
}

func (obj APIDeviceItem) Result() DeviceItem {
	return DeviceItem{
		ID:              pointy.IntValue(obj.ID, 0),
		SerialNumber:    pointy.StringValue(obj.SerialNumber, ""),
		Name:            pointy.StringValue(obj.Name, ""),
		TypeID:          pointy.IntValue(obj.TypeID, 0),
		InverterModel:   pointy.StringValue(obj.InverterModel, ""),
		Latitude:        pointy.Float64Value(obj.Latitude, 0),
		Longitude:       pointy.Float64Value(obj.Longitude, 0),
		SoftwareVersion: pointy.StringValue(obj.SoftwareVersion, ""),
		PlantCode:       pointy.StringValue(obj.PlantCode, ""),
	}
}

// ===========================
// API Fetch: GetRealtimeDeviceData
// Result Model: RealtimeDeviceData, RealtimeDeviceItem
// ===========================
type APIGetRealtimeDeviceDataResponse struct {
	APIDefaultResponse
	Data []APIRealtimeDeviceData `json:"data,omitempty"`
}

type APIRealtimeDeviceData struct {
	ID          *int                  `json:"devId,omitempty"`
	DataItemMap APIRealtimeDeviceItem `json:"dataItemMap,omitempty"`
}

func (obj APIRealtimeDeviceData) Result() RealtimeDeviceData {
	return RealtimeDeviceData{
		ID:          pointy.IntValue(obj.ID, 0),
		DataItemMap: obj.DataItemMap.Result(),
	}
}

type APIRealtimeDeviceItem struct {
	InverterState      *float64     `json:"inverter_state,omitempty"`
	GridABVoltage      *float64     `json:"ab_u,omitempty"`
	GridBCVoltage      *float64     `json:"bc_u,omitempty"`
	GridCAVoltage      *float64     `json:"ca_u,omitempty"`
	PhaseAVoltage      *float64     `json:"a_u,omitempty"`
	PhaseBVoltage      *float64     `json:"b_u,omitempty"`
	PhaseCVoltage      *float64     `json:"c_u,omitempty"`
	GridPhaseACurrent  *float64     `json:"a_i,omitempty"`
	GridPhaseBCurrent  *float64     `json:"b_i,omitempty"`
	GridPhaseCCurrent  *float64     `json:"c_i,omitempty"`
	Efficiency         *float64     `json:"efficiency,omitempty"`
	Temperature        *float64     `json:"temperature,omitempty"`
	PowerFactor        *float64     `json:"power_factor,omitempty"`
	GridFrequency      *float64     `json:"elec_freq,omitempty"`
	ActivePower        *float64     `json:"active_power,omitempty"`
	ReactivePower      *float64     `json:"reactive_power,omitempty"`
	DailyEnergy        *float64     `json:"day_cap,omitempty"`
	MPPTTotalPower     *float64     `json:"mppt_power,omitempty"`
	PV1InputVoltage    *float64     `json:"pv1_u,omitempty"`
	PV2InputVoltage    *float64     `json:"pv2_u,omitempty"`
	PV3InputVoltage    *float64     `json:"pv3_u,omitempty"`
	PV4InputVoltage    *float64     `json:"pv4_u,omitempty"`
	PV5InputVoltage    *float64     `json:"pv5_u,omitempty"`
	PV6InputVoltage    *float64     `json:"pv6_u,omitempty"`
	PV7InputVoltage    *float64     `json:"pv7_u,omitempty"`
	PV8InputVoltage    *float64     `json:"pv8_u,omitempty"`
	PV9InputVoltage    *float64     `json:"pv9_u,omitempty"`
	PV10InputVoltage   *float64     `json:"pv10_u,omitempty"`
	PV11InputVoltage   *float64     `json:"pv11_u,omitempty"`
	PV12InputVoltage   *float64     `json:"pv12_u,omitempty"`
	PV13InputVoltage   *float64     `json:"pv13_u,omitempty"`
	PV14InputVoltage   *float64     `json:"pv14_u,omitempty"`
	PV15InputVoltage   *float64     `json:"pv15_u,omitempty"`
	PV16InputVoltage   *float64     `json:"pv16_u,omitempty"`
	PV17InputVoltage   *float64     `json:"pv17_u,omitempty"`
	PV18InputVoltage   *float64     `json:"pv18_u,omitempty"`
	PV19InputVoltage   *float64     `json:"pv19_u,omitempty"`
	PV20InputVoltage   *float64     `json:"pv20_u,omitempty"`
	PV21InputVoltage   *float64     `json:"pv21_u,omitempty"`
	PV22InputVoltage   *float64     `json:"pv22_u,omitempty"`
	PV23InputVoltage   *float64     `json:"pv23_u,omitempty"`
	PV24InputVoltage   *float64     `json:"pv24_u,omitempty"`
	PV1InputCurrent    *float64     `json:"pv1_i,omitempty"`
	PV2InputCurrent    *float64     `json:"pv2_i,omitempty"`
	PV3InputCurrent    *float64     `json:"pv3_i,omitempty"`
	PV4InputCurrent    *float64     `json:"pv4_i,omitempty"`
	PV5InputCurrent    *float64     `json:"pv5_i,omitempty"`
	PV6InputCurrent    *float64     `json:"pv6_i,omitempty"`
	PV7InputCurrent    *float64     `json:"pv7_i,omitempty"`
	PV8InputCurrent    *float64     `json:"pv8_i,omitempty"`
	PV9InputCurrent    *float64     `json:"pv9_i,omitempty"`
	PV10InputCurrent   *float64     `json:"pv10_i,omitempty"`
	PV11InputCurrent   *float64     `json:"pv11_i,omitempty"`
	PV12InputCurrent   *float64     `json:"pv12_i,omitempty"`
	PV13InputCurrent   *float64     `json:"pv13_i,omitempty"`
	PV14InputCurrent   *float64     `json:"pv14_i,omitempty"`
	PV15InputCurrent   *float64     `json:"pv15_i,omitempty"`
	PV16InputCurrent   *float64     `json:"pv16_i,omitempty"`
	PV17InputCurrent   *float64     `json:"pv17_i,omitempty"`
	PV18InputCurrent   *float64     `json:"pv18_i,omitempty"`
	PV19InputCurrent   *float64     `json:"pv19_i,omitempty"`
	PV20InputCurrent   *float64     `json:"pv20_i,omitempty"`
	PV21InputCurrent   *float64     `json:"pv21_i,omitempty"`
	PV22InputCurrent   *float64     `json:"pv22_i,omitempty"`
	PV23InputCurrent   *float64     `json:"pv23_i,omitempty"`
	PV24InputCurrent   *float64     `json:"pv24_i,omitempty"`
	TotalEnergy        *float64     `json:"total_cap,omitempty"`
	InverterStartup    *interface{} `json:"open_time,omitempty"`
	InverterShutdown   *interface{} `json:"close_time,omitempty"`
	TotalDCInputEnergy *float64     `json:"mppt_total_cap,omitempty"`
	MPPT1DCEnergy      *float64     `json:"mppt_1_cap,omitempty"`
	MPPT2DCEnergy      *float64     `json:"mppt_2_cap,omitempty"`
	MPPT3DCEnergy      *float64     `json:"mppt_3_cap,omitempty"`
	MPPT4DCEnergy      *float64     `json:"mppt_4_cap,omitempty"`
	MPPT5DCEnergy      *float64     `json:"mppt_5_cap,omitempty"`
	MPPT6DCEnergy      *float64     `json:"mppt_6_cap,omitempty"`
	MPPT7DCEnergy      *float64     `json:"mppt_7_cap,omitempty"`
	MPPT8DCEnergy      *float64     `json:"mppt_8_cap,omitempty"`
	MPPT9DCEnergy      *float64     `json:"mppt_9_cap,omitempty"`
	MPPT10DCEnergy     *float64     `json:"mppt_10_cap,omitempty"`
	Status             *int         `json:"run_state,omitempty"`
}

func (obj APIRealtimeDeviceItem) Result() RealtimeDeviceItem {
	return RealtimeDeviceItem{
		InverterState:      pointy.Float64Value(obj.InverterState, 0),
		GridABVoltage:      pointy.Float64Value(obj.GridABVoltage, 0),
		GridBCVoltage:      pointy.Float64Value(obj.GridBCVoltage, 0),
		GridCAVoltage:      pointy.Float64Value(obj.GridCAVoltage, 0),
		PhaseAVoltage:      pointy.Float64Value(obj.PhaseAVoltage, 0),
		PhaseBVoltage:      pointy.Float64Value(obj.PhaseBVoltage, 0),
		PhaseCVoltage:      pointy.Float64Value(obj.PhaseCVoltage, 0),
		GridPhaseACurrent:  pointy.Float64Value(obj.GridPhaseACurrent, 0),
		GridPhaseBCurrent:  pointy.Float64Value(obj.GridPhaseBCurrent, 0),
		GridPhaseCCurrent:  pointy.Float64Value(obj.GridPhaseCCurrent, 0),
		Efficiency:         pointy.Float64Value(obj.Efficiency, 0),
		Temperature:        pointy.Float64Value(obj.Temperature, 0),
		PowerFactor:        pointy.Float64Value(obj.PowerFactor, 0),
		GridFrequency:      pointy.Float64Value(obj.GridFrequency, 0),
		ActivePower:        pointy.Float64Value(obj.ActivePower, 0),
		ReactivePower:      pointy.Float64Value(obj.ReactivePower, 0),
		DailyEnergy:        pointy.Float64Value(obj.DailyEnergy, 0),
		MPPTTotalPower:     pointy.Float64Value(obj.MPPTTotalPower, 0),
		PV1InputVoltage:    pointy.Float64Value(obj.PV1InputVoltage, 0),
		PV2InputVoltage:    pointy.Float64Value(obj.PV2InputVoltage, 0),
		PV3InputVoltage:    pointy.Float64Value(obj.PV3InputVoltage, 0),
		PV4InputVoltage:    pointy.Float64Value(obj.PV4InputVoltage, 0),
		PV5InputVoltage:    pointy.Float64Value(obj.PV5InputVoltage, 0),
		PV6InputVoltage:    pointy.Float64Value(obj.PV6InputVoltage, 0),
		PV7InputVoltage:    pointy.Float64Value(obj.PV7InputVoltage, 0),
		PV8InputVoltage:    pointy.Float64Value(obj.PV8InputVoltage, 0),
		PV9InputVoltage:    pointy.Float64Value(obj.PV9InputVoltage, 0),
		PV10InputVoltage:   pointy.Float64Value(obj.PV10InputVoltage, 0),
		PV11InputVoltage:   pointy.Float64Value(obj.PV11InputVoltage, 0),
		PV12InputVoltage:   pointy.Float64Value(obj.PV12InputVoltage, 0),
		PV13InputVoltage:   pointy.Float64Value(obj.PV13InputVoltage, 0),
		PV14InputVoltage:   pointy.Float64Value(obj.PV14InputVoltage, 0),
		PV15InputVoltage:   pointy.Float64Value(obj.PV15InputVoltage, 0),
		PV16InputVoltage:   pointy.Float64Value(obj.PV16InputVoltage, 0),
		PV17InputVoltage:   pointy.Float64Value(obj.PV17InputVoltage, 0),
		PV18InputVoltage:   pointy.Float64Value(obj.PV18InputVoltage, 0),
		PV19InputVoltage:   pointy.Float64Value(obj.PV19InputVoltage, 0),
		PV20InputVoltage:   pointy.Float64Value(obj.PV20InputVoltage, 0),
		PV21InputVoltage:   pointy.Float64Value(obj.PV21InputVoltage, 0),
		PV22InputVoltage:   pointy.Float64Value(obj.PV22InputVoltage, 0),
		PV23InputVoltage:   pointy.Float64Value(obj.PV23InputVoltage, 0),
		PV24InputVoltage:   pointy.Float64Value(obj.PV24InputVoltage, 0),
		PV1InputCurrent:    pointy.Float64Value(obj.PV1InputCurrent, 0),
		PV2InputCurrent:    pointy.Float64Value(obj.PV2InputCurrent, 0),
		PV3InputCurrent:    pointy.Float64Value(obj.PV3InputCurrent, 0),
		PV4InputCurrent:    pointy.Float64Value(obj.PV4InputCurrent, 0),
		PV5InputCurrent:    pointy.Float64Value(obj.PV5InputCurrent, 0),
		PV6InputCurrent:    pointy.Float64Value(obj.PV6InputCurrent, 0),
		PV7InputCurrent:    pointy.Float64Value(obj.PV7InputCurrent, 0),
		PV8InputCurrent:    pointy.Float64Value(obj.PV8InputCurrent, 0),
		PV9InputCurrent:    pointy.Float64Value(obj.PV9InputCurrent, 0),
		PV10InputCurrent:   pointy.Float64Value(obj.PV10InputCurrent, 0),
		PV11InputCurrent:   pointy.Float64Value(obj.PV11InputCurrent, 0),
		PV12InputCurrent:   pointy.Float64Value(obj.PV12InputCurrent, 0),
		PV13InputCurrent:   pointy.Float64Value(obj.PV13InputCurrent, 0),
		PV14InputCurrent:   pointy.Float64Value(obj.PV14InputCurrent, 0),
		PV15InputCurrent:   pointy.Float64Value(obj.PV15InputCurrent, 0),
		PV16InputCurrent:   pointy.Float64Value(obj.PV16InputCurrent, 0),
		PV17InputCurrent:   pointy.Float64Value(obj.PV17InputCurrent, 0),
		PV18InputCurrent:   pointy.Float64Value(obj.PV18InputCurrent, 0),
		PV19InputCurrent:   pointy.Float64Value(obj.PV19InputCurrent, 0),
		PV20InputCurrent:   pointy.Float64Value(obj.PV20InputCurrent, 0),
		PV21InputCurrent:   pointy.Float64Value(obj.PV21InputCurrent, 0),
		PV22InputCurrent:   pointy.Float64Value(obj.PV22InputCurrent, 0),
		PV23InputCurrent:   pointy.Float64Value(obj.PV23InputCurrent, 0),
		PV24InputCurrent:   pointy.Float64Value(obj.PV24InputCurrent, 0),
		TotalEnergy:        pointy.Float64Value(obj.TotalEnergy, 0),
		InverterStartup:    pointy.PointerValue(obj.InverterStartup, nil),
		InverterShutdown:   pointy.PointerValue(obj.InverterShutdown, nil),
		TotalDCInputEnergy: pointy.Float64Value(obj.TotalDCInputEnergy, 0),
		MPPT1DCEnergy:      pointy.Float64Value(obj.MPPT1DCEnergy, 0),
		MPPT2DCEnergy:      pointy.Float64Value(obj.MPPT2DCEnergy, 0),
		MPPT3DCEnergy:      pointy.Float64Value(obj.MPPT3DCEnergy, 0),
		MPPT4DCEnergy:      pointy.Float64Value(obj.MPPT4DCEnergy, 0),
		MPPT5DCEnergy:      pointy.Float64Value(obj.MPPT5DCEnergy, 0),
		MPPT6DCEnergy:      pointy.Float64Value(obj.MPPT6DCEnergy, 0),
		MPPT7DCEnergy:      pointy.Float64Value(obj.MPPT7DCEnergy, 0),
		MPPT8DCEnergy:      pointy.Float64Value(obj.MPPT8DCEnergy, 0),
		MPPT9DCEnergy:      pointy.Float64Value(obj.MPPT9DCEnergy, 0),
		MPPT10DCEnergy:     pointy.Float64Value(obj.MPPT10DCEnergy, 0),
		Status:             pointy.IntValue(obj.Status, 0),
	}
}

// ===========================
// API Fetch: GetDailyDeviceData, GetMonthlyDeviceData, GetYearlyDeviceData
// Result Model: HistoricalDeviceData, HistoricalDeviceItem
// ===========================
type APIGetHistoricalDeviceDataResponse struct {
	APIDefaultResponse
	Data []APIHistoricalDeviceData `json:"data,omitempty"`
}

type APIHistoricalDeviceData struct {
	ID          *interface{}            `json:"devId,omitempty"`
	CollectTime *int64                  `json:"collectTime,omitempty"`
	DataItemMap APIHistoricalDeviceItem `json:"dataItemMap,omitempty"`
}

func (obj APIHistoricalDeviceData) Result() HistoricalDeviceData {
	return HistoricalDeviceData{
		ID:          pointy.PointerValue(obj.ID, nil),
		CollectTime: pointy.Int64Value(obj.CollectTime, 0),
		DataItemMap: obj.DataItemMap.Result(),
	}
}

type APIHistoricalDeviceItem struct {
	InstalledCapacity *float64 `json:"installed_capacity,omitempty"`
	ProductPower      *float64 `json:"product_power,omitempty"`
	PerPowerRatio     *float64 `json:"perpower_ratio,omitempty"`
}

func (obj APIHistoricalDeviceItem) Result() HistoricalDeviceItem {
	return HistoricalDeviceItem{
		InstalledCapacity: pointy.Float64Value(obj.InstalledCapacity, 0),
		ProductPower:      pointy.Float64Value(obj.ProductPower, 0),
		PerPowerRatio:     pointy.Float64Value(obj.PerPowerRatio, 0),
	}
}

// ===========================
// API Fetch: GetDeviceAlarmList
// Result Model: DeviceAlarmItem
// ===========================
type APIGetDeviceAlarmListResponse struct {
	APIDefaultResponse
	Data []APIDeviceAlarmItem `json:"data,omitempty"`
}

type APIDeviceAlarmItem struct {
	PlantCode        *string `json:"stationCode,omitempty"`
	PlantName        *string `json:"stationName,omitempty"`
	SerialNumber     *string `json:"esnCode,omitempty"`
	DeviceName       *string `json:"devName,omitempty"`
	DeviceTypeID     *int    `json:"devTypeId,omitempty"`
	AlarmID          *int    `json:"alarmId,omitempty"`
	AlarmName        *string `json:"alarmName,omitempty"`
	AlarmCause       *string `json:"alarmCause,omitempty"`
	AlarmType        *int    `json:"alarmType,omitempty"`
	RepairSuggestion *string `json:"repairSuggestion,omitempty"`
	CauseID          *int    `json:"causeId,omitempty"`
	RaiseTime        *int64  `json:"raiseTime,omitempty"`
	Level            *int    `json:"lev,omitempty"`
	Status           *int    `json:"status,omitempty"`
}

func (obj APIDeviceAlarmItem) Result() DeviceAlarmItem {
	return DeviceAlarmItem{
		PlantCode:        pointy.StringValue(obj.PlantCode, ""),
		PlantName:        pointy.StringValue(obj.PlantName, ""),
		SerialNumber:     pointy.StringValue(obj.SerialNumber, ""),
		DeviceName:       pointy.StringValue(obj.DeviceName, ""),
		DeviceTypeID:     pointy.IntValue(obj.DeviceTypeID, 0),
		AlarmID:          pointy.IntValue(obj.AlarmID, 0),
		AlarmName:        pointy.StringValue(obj.AlarmName, ""),
		AlarmCause:       pointy.StringValue(obj.AlarmCause, ""),
		AlarmType:        pointy.IntValue(obj.AlarmType, 0),
		RepairSuggestion: pointy.StringValue(obj.RepairSuggestion, ""),
		CauseID:          pointy.IntValue(obj.CauseID, 0),
		RaiseTime:        pointy.Int64Value(obj.RaiseTime, 0),
		Level:            pointy.IntValue(obj.Level, 0),
		Status:           pointy.IntValue(obj.Status, 0),
	}
}
