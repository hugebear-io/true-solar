package huawei

// ===== DEFAULT RESPONSE =====
type DefaultResponse struct {
	Success    bool        `json:"success,omitempty"`
	FailCode   int         `json:"failCode,omitempty"`
	Parameters interface{} `json:"params,omitempty"`
	Message    string      `json:"message,omitempty"`
}

// ===== PlantItem =====
type PlantItem struct {
	Code           string  `json:"stationCode,omitempty"`
	Name           string  `json:"stationName,omitempty"`
	Address        string  `json:"stationAddr,omitempty"`
	Capacity       float64 `json:"capacity,omitempty"`
	BuildState     string  `json:"buildState,omitempty"`
	CombineType    string  `json:"combineType,omitempty"`
	AIDType        int     `json:"aidType,omitempty"`
	StationLinkMan string  `json:"stationLinkman,omitempty"`
	LinkManPhone   string  `json:"linkmanPho,omitempty"`
}

// ===== PlantRealTime =====
type PlantRealtimeData struct {
	Code        string            `json:"stationCode,omitempty"`
	DataItemMap PlantRealtimeItem `json:"dataItemMap,omitempty"`
}

type PlantRealtimeItem struct {
	TotalIncome     float64 `json:"total_income,omitempty"`
	TotalPower      float64 `json:"total_power,omitempty"`
	DayPower        float64 `json:"day_power,omitempty"`
	DayIncome       float64 `json:"day_income,omitempty"`
	RealHealthState int     `json:"real_health_state,omitempty"`
	MonthPower      float64 `json:"month_power,omitempty"`
}

// ===== HistoricalPlant =====
type HistoricalPlantData struct {
	Code        string              `json:"stationCode,omitempty"`
	CollectTime int64               `json:"collectTime,omitempty"`
	DataItemMap HistoricalPlantItem `json:"dataItemMap,omitempty"`
}

type HistoricalPlantItem struct {
	RadiationIntensity float64 `json:"radiation_intensity,omitempty"`
	InstalledCapacity  float64 `json:"installed_capacity,omitempty"`
	UsePower           float64 `json:"use_power,omitempty"`
	InverterPower      float64 `json:"inverter_power,omitempty"`
	PowerProfit        float64 `json:"power_profit,omitempty"`
	TheoryPower        float64 `json:"theory_power,omitempty"`
	PerPowerRatio      float64 `json:"perpower_ratio,omitempty"`
	OnGridPower        float64 `json:"ongrid_power,omitempty"`
	PerformanceRatio   float64 `json:"performance_ratio,omitempty"`
	ReductionTotalCO2  float64 `json:"reduction_total_co2,omitempty"`
	ReductionTotalCoal float64 `json:"reduction_total_coal,omitempty"`
	ReductionTotalTree float64 `json:"reduction_total_tree,omitempty"`
}

// ===== DeviceItem =====
type DeviceItem struct {
	ID              int     `json:"id,omitempty"`
	SerialNumber    string  `json:"esnCode,omitempty"`
	Name            string  `json:"devName,omitempty"`
	TypeID          int     `json:"devTypeId,omitempty"`
	InverterModel   string  `json:"invType,omitempty"`
	Latitude        float64 `json:"latitude,omitempty"`
	Longitude       float64 `json:"longitude,omitempty"`
	SoftwareVersion string  `json:"softwareVersion,omitempty"`
	PlantCode       string  `json:"stationCode,omitempty"`
}

// ===== RealtimeDeviceItem =====
type RealtimeDeviceData struct {
	ID          int
	DataItemMap RealtimeDeviceItem
}

type RealtimeDeviceItem struct {
	InverterState      float64     `json:"inverter_state,omitempty"`
	GridABVoltage      float64     `json:"ab_u,omitempty"`
	GridBCVoltage      float64     `json:"bc_u,omitempty"`
	GridCAVoltage      float64     `json:"ca_u,omitempty"`
	PhaseAVoltage      float64     `json:"a_u,omitempty"`
	PhaseBVoltage      float64     `json:"b_u,omitempty"`
	PhaseCVoltage      float64     `json:"c_u,omitempty"`
	GridPhaseACurrent  float64     `json:"a_i,omitempty"`
	GridPhaseBCurrent  float64     `json:"b_i,omitempty"`
	GridPhaseCCurrent  float64     `json:"c_i,omitempty"`
	Efficiency         float64     `json:"efficiency,omitempty"`
	Temperature        float64     `json:"temperature,omitempty"`
	PowerFactor        float64     `json:"power_factor,omitempty"`
	GridFrequency      float64     `json:"elec_freq,omitempty"`
	ActivePower        float64     `json:"active_power,omitempty"`
	ReactivePower      float64     `json:"reactive_power,omitempty"`
	DailyEnergy        float64     `json:"day_cap,omitempty"`
	MPPTTotalPower     float64     `json:"mppt_power,omitempty"`
	PV1InputVoltage    float64     `json:"pv1_u,omitempty"`
	PV2InputVoltage    float64     `json:"pv2_u,omitempty"`
	PV3InputVoltage    float64     `json:"pv3_u,omitempty"`
	PV4InputVoltage    float64     `json:"pv4_u,omitempty"`
	PV5InputVoltage    float64     `json:"pv5_u,omitempty"`
	PV6InputVoltage    float64     `json:"pv6_u,omitempty"`
	PV7InputVoltage    float64     `json:"pv7_u,omitempty"`
	PV8InputVoltage    float64     `json:"pv8_u,omitempty"`
	PV9InputVoltage    float64     `json:"pv9_u,omitempty"`
	PV10InputVoltage   float64     `json:"pv10_u,omitempty"`
	PV11InputVoltage   float64     `json:"pv11_u,omitempty"`
	PV12InputVoltage   float64     `json:"pv12_u,omitempty"`
	PV13InputVoltage   float64     `json:"pv13_u,omitempty"`
	PV14InputVoltage   float64     `json:"pv14_u,omitempty"`
	PV15InputVoltage   float64     `json:"pv15_u,omitempty"`
	PV16InputVoltage   float64     `json:"pv16_u,omitempty"`
	PV17InputVoltage   float64     `json:"pv17_u,omitempty"`
	PV18InputVoltage   float64     `json:"pv18_u,omitempty"`
	PV19InputVoltage   float64     `json:"pv19_u,omitempty"`
	PV20InputVoltage   float64     `json:"pv20_u,omitempty"`
	PV21InputVoltage   float64     `json:"pv21_u,omitempty"`
	PV22InputVoltage   float64     `json:"pv22_u,omitempty"`
	PV23InputVoltage   float64     `json:"pv23_u,omitempty"`
	PV24InputVoltage   float64     `json:"pv24_u,omitempty"`
	PV1InputCurrent    float64     `json:"pv1_i,omitempty"`
	PV2InputCurrent    float64     `json:"pv2_i,omitempty"`
	PV3InputCurrent    float64     `json:"pv3_i,omitempty"`
	PV4InputCurrent    float64     `json:"pv4_i,omitempty"`
	PV5InputCurrent    float64     `json:"pv5_i,omitempty"`
	PV6InputCurrent    float64     `json:"pv6_i,omitempty"`
	PV7InputCurrent    float64     `json:"pv7_i,omitempty"`
	PV8InputCurrent    float64     `json:"pv8_i,omitempty"`
	PV9InputCurrent    float64     `json:"pv9_i,omitempty"`
	PV10InputCurrent   float64     `json:"pv10_i,omitempty"`
	PV11InputCurrent   float64     `json:"pv11_i,omitempty"`
	PV12InputCurrent   float64     `json:"pv12_i,omitempty"`
	PV13InputCurrent   float64     `json:"pv13_i,omitempty"`
	PV14InputCurrent   float64     `json:"pv14_i,omitempty"`
	PV15InputCurrent   float64     `json:"pv15_i,omitempty"`
	PV16InputCurrent   float64     `json:"pv16_i,omitempty"`
	PV17InputCurrent   float64     `json:"pv17_i,omitempty"`
	PV18InputCurrent   float64     `json:"pv18_i,omitempty"`
	PV19InputCurrent   float64     `json:"pv19_i,omitempty"`
	PV20InputCurrent   float64     `json:"pv20_i,omitempty"`
	PV21InputCurrent   float64     `json:"pv21_i,omitempty"`
	PV22InputCurrent   float64     `json:"pv22_i,omitempty"`
	PV23InputCurrent   float64     `json:"pv23_i,omitempty"`
	PV24InputCurrent   float64     `json:"pv24_i,omitempty"`
	TotalEnergy        float64     `json:"total_cap,omitempty"`
	InverterStartup    interface{} `json:"open_time,omitempty"`
	InverterShutdown   interface{} `json:"close_time,omitempty"`
	TotalDCInputEnergy float64     `json:"mppt_total_cap,omitempty"`
	MPPT1DCEnergy      float64     `json:"mppt_1_cap,omitempty"`
	MPPT2DCEnergy      float64     `json:"mppt_2_cap,omitempty"`
	MPPT3DCEnergy      float64     `json:"mppt_3_cap,omitempty"`
	MPPT4DCEnergy      float64     `json:"mppt_4_cap,omitempty"`
	MPPT5DCEnergy      float64     `json:"mppt_5_cap,omitempty"`
	MPPT6DCEnergy      float64     `json:"mppt_6_cap,omitempty"`
	MPPT7DCEnergy      float64     `json:"mppt_7_cap,omitempty"`
	MPPT8DCEnergy      float64     `json:"mppt_8_cap,omitempty"`
	MPPT9DCEnergy      float64     `json:"mppt_9_cap,omitempty"`
	MPPT10DCEnergy     float64     `json:"mppt_10_cap,omitempty"`
	Status             int         `json:"run_state,omitempty"`
}

// ===== HistoricalDeviceItem =====
type HistoricalDeviceData struct {
	ID          interface{}          `json:"devId,omitempty"`
	CollectTime int64                `json:"collectTime,omitempty"`
	DataItemMap HistoricalDeviceItem `json:"dataItemMap,omitempty"`
}

type HistoricalDeviceItem struct {
	InstalledCapacity float64 `json:"installed_capacity,omitempty"`
	ProductPower      float64 `json:"product_power,omitempty"`
	PerPowerRatio     float64 `json:"perpower_ratio,omitempty"`
}

// ===== DeviceAlarmItem =====
type DeviceAlarmItem struct {
	PlantCode        string `json:"stationCode,omitempty"`
	PlantName        string `json:"stationName,omitempty"`
	SerialNumber     string `json:"esnCode,omitempty"`
	DeviceName       string `json:"devName,omitempty"`
	DeviceTypeID     int    `json:"devTypeId,omitempty"`
	AlarmID          int    `json:"alarmId,omitempty"`
	AlarmName        string `json:"alarmName,omitempty"`
	AlarmCause       string `json:"alarmCause,omitempty"`
	AlarmType        int    `json:"alarmType,omitempty"`
	RepairSuggestion string `json:"repairSuggestion,omitempty"`
	CauseID          int    `json:"causeId,omitempty"`
	RaiseTime        int64  `json:"raiseTime,omitempty"`
	Level            int    `json:"lev,omitempty"`
	Status           int    `json:"status,omitempty"`
}
