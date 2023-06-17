package growatt

import "github.com/openlyinc/pointy"

// ===== DEFAULT RESPONSE =====
type APIDefaultResponse struct {
	ErrorCode *int    `json:"error_code,omitempty"`
	ErrorMsg  *string `json:"error_msg,omitempty"`
}

func (obj APIDefaultResponse) Result() DefaultResponse {
	return DefaultResponse{
		ErrorCode: pointy.IntValue(obj.ErrorCode, 0),
		ErrorMsg:  pointy.StringValue(obj.ErrorMsg, ""),
	}
}

// ===== ERROR RESPONSE =====
type APIErrorResponse struct {
	Data *string `json:"data,omitempty"`
	APIDefaultResponse
}

func (obj APIErrorResponse) Result() ErrorResponse {
	return ErrorResponse{
		Data:            pointy.StringValue(obj.Data, ""),
		DefaultResponse: obj.APIDefaultResponse.Result(),
	}
}

// ==============================
// API Fetch: GetPlantList
// Result Model: GetPlantListResopnse, GetPlantData, PlantItem
// ==============================
type APIGetPlantListResponse struct {
	Data *APIGetPlantListData `json:"data,omitempty"`
	APIDefaultResponse
}

func (obj APIGetPlantListResponse) Result() GetPlantListResponse {
	return GetPlantListResponse{
		Data:            obj.Data.Result(),
		DefaultResponse: obj.APIDefaultResponse.Result(),
	}
}

type APIGetPlantListData struct {
	Count  *int           `json:"count,omitempty"`
	Plants []APIPlantItem `json:"plants,omitempty"`
}

func (obj APIGetPlantListData) Result() GetPlantListData {
	plants := []PlantItem{}
	for _, item := range obj.Plants {
		plants = append(plants, item.Result())
	}

	return GetPlantListData{
		Count:  pointy.IntValue(obj.Count, 0),
		Plants: plants,
	}
}

type APIPlantItem struct {
	Status       *int         `json:"status,omitempty"`
	Locale       *string      `json:"locale,omitempty"`
	TotalEnergy  *string      `json:"total_energy,omitempty"`
	Operator     *string      `json:"operator,omitempty"`
	Country      *string      `json:"country,omitempty"`
	City         *string      `json:"city,omitempty"`
	CurrentPower *string      `json:"current_power,omitempty"`
	CreateDate   *string      `json:"create_date,omitempty"`
	ImageURL     *string      `json:"image_url,omitempty"`
	PlantID      *int         `json:"plant_id,omitempty"`
	Name         *string      `json:"name,omitempty"`
	Installer    *string      `json:"installer,omitempty"`
	UserID       *int         `json:"user_id,omitempty"`
	Longitude    *string      `json:"longitude,omitempty"`
	Latitude     *string      `json:"latitude,omitempty"`
	PeakPower    *float64     `json:"peak_power,omitempty"`
	LatitudeD    *interface{} `json:"latitude_d,omitempty"`
	LatitudeF    *interface{} `json:"latitude_f,omitempty"`
}

func (obj APIPlantItem) Result() PlantItem {
	return PlantItem{
		Status:       pointy.IntValue(obj.Status, 0),
		Locale:       pointy.StringValue(obj.Locale, ""),
		TotalEnergy:  pointy.StringValue(obj.TotalEnergy, ""),
		Operator:     pointy.StringValue(obj.Operator, ""),
		Country:      pointy.StringValue(obj.Country, ""),
		City:         pointy.StringValue(obj.City, ""),
		CurrentPower: pointy.StringValue(obj.CurrentPower, ""),
		CreateDate:   pointy.StringValue(obj.CreateDate, ""),
		ImageURL:     pointy.StringValue(obj.ImageURL, ""),
		PlantID:      pointy.IntValue(obj.PlantID, 0),
		Name:         pointy.StringValue(obj.Name, ""),
		Installer:    pointy.StringValue(obj.Installer, ""),
		UserID:       pointy.IntValue(obj.UserID, 0),
		Longitude:    pointy.StringValue(obj.Longitude, ""),
		Latitude:     pointy.StringValue(obj.Latitude, ""),
		PeakPower:    pointy.Float64Value(obj.PeakPower, 0),
		LatitudeD:    pointy.PointerValue(obj.LatitudeD, nil),
		LatitudeF:    pointy.PointerValue(obj.LatitudeF, nil),
	}
}

// ==============================
// API Fetch: GetPlantOverviewInfo
// Result Model: GetPlantOverviewResponse, PlantOverviewData
// ==============================
type APIGetPlantOverviewInfoResponse struct {
	Data *APIPlantOverviewInfoData `json:"data,omitempty"`
	APIDefaultResponse
}

func (obj APIGetPlantOverviewInfoResponse) Result() GetPlantOverviewInfoResponse {
	return GetPlantOverviewInfoResponse{
		Data:            obj.Data.Result(),
		DefaultResponse: obj.APIDefaultResponse.Result(),
	}
}

type APIPlantOverviewInfoData struct {
	PeakPowerActual *float64 `json:"peak_power_actual,omitempty"`
	MonthlyEnergy   *string  `json:"monthly_energy,omitempty"`
	LastUpdateTime  *string  `json:"last_update_time,omitempty"`
	CurrentPower    *float64 `json:"current_power,omitempty"`
	Timezone        *string  `json:"timezone,omitempty"`
	YearlyEnergy    *string  `json:"yearly_energy,omitempty"`
	TodayEnergy     *string  `json:"today_energy,omitempty"`
	CarbonOffset    *string  `json:"carbon_offset,omitempty"`
	Efficiency      *string  `json:"efficiency,omitempty"`
	TotalEnergy     *string  `json:"total_energy,omitempty"`
}

func (obj APIPlantOverviewInfoData) Result() PlantOverviewInfoData {
	return PlantOverviewInfoData{
		PeakPowerActual: pointy.Float64Value(obj.PeakPowerActual, 0),
		MonthlyEnergy:   pointy.StringValue(obj.MonthlyEnergy, ""),
		LastUpdateTime:  pointy.StringValue(obj.LastUpdateTime, ""),
		CurrentPower:    pointy.Float64Value(obj.CurrentPower, 0),
		Timezone:        pointy.StringValue(obj.Timezone, ""),
		YearlyEnergy:    pointy.StringValue(obj.YearlyEnergy, ""),
		TodayEnergy:     pointy.StringValue(obj.TodayEnergy, ""),
		CarbonOffset:    pointy.StringValue(obj.CarbonOffset, ""),
		Efficiency:      pointy.StringValue(obj.Efficiency, ""),
		TotalEnergy:     pointy.StringValue(obj.TotalEnergy, ""),
	}
}

// ==============================
// API Fetch: GetPlantDataLoggerInfo
// Result Model: GetPlantDataLoggerInfoResponse, GetPlantDataLoggerInfoData, PeakPowerActual, DataLoggerItem, DateTime
// ==============================
type APIGetPlantDataLoggerInfoResponse struct {
	Data *APIGetDataLoggerInfoData `json:"data,omitempty"`
	APIDefaultResponse
}

func (obj APIGetPlantDataLoggerInfoResponse) Result() GetPlantDataLoggerInfoResponse {
	return GetPlantDataLoggerInfoResponse{
		Data:            obj.Data.Result(),
		DefaultResponse: obj.APIDefaultResponse.Result(),
	}
}

type APIGetDataLoggerInfoData struct {
	PeakPowerActual *APIPeakPowerActual `json:"peak_power_actual,omitempty"`
	Count           *int                `json:"count,omitempty"`
	DataLoggers     []APIDataLoggerItem `json:"dataloggers,omitempty"`
}

func (obj APIGetDataLoggerInfoData) Result() PlantDataLoggerInfoData {
	dataLoggers := []DataLoggerItem{}
	for _, item := range obj.DataLoggers {
		dataLoggers = append(dataLoggers, item.Result())
	}

	return PlantDataLoggerInfoData{
		PeakPowerActual: obj.PeakPowerActual.Result(),
		Count:           pointy.IntValue(obj.Count, 0),
		DataLoggers:     dataLoggers,
	}
}

type APIPeakPowerActual struct {
	MapCityID                *int         `json:"map_cityId,omitempty"`
	MapCountryID             *int         `json:"map_countryId,omitempty"`
	StorageBatteryPercentage *float64     `json:"storage_BattoryPercentage,omitempty"`
	DefaultPlant             *bool        `json:"defaultPlant,omitempty"`
	PeakPeriodPrice          *float64     `json:"peakPeriodPrice,omitempty"`
	City                     *string      `json:"city,omitempty"`
	NominalPower             *float64     `json:"nominalPower,omitempty"`
	AlarmValue               *float64     `json:"alarmValue,omitempty"`
	CurrentPacTxt            *string      `json:"currentPacTxt,omitempty"`
	FixedPowerPrice          *float64     `json:"fixedPowerPrice,omitempty"`
	PlantFromBean            *interface{} `json:"plantFromBean,omitempty"`
	DeviceCount              *int         `json:"deviceCount,omitempty"`
	PlantImgName             *string      `json:"plantImgName,omitempty"`
	EtodaySo2Text            *string      `json:"etodaySo2Text,omitempty"`
	CompanyName              *string      `json:"companyName,omitempty"`
	EmonthMoneyText          *string      `json:"emonthMoneyText,omitempty"`
	FormulaMoney             *float64     `json:"formulaMoney,omitempty"`
	UserAccount              *string      `json:"userAccount,omitempty"`
	MapLat                   *string      `json:"mapLat,omitempty"`
	CreateDateTextA          *string      `json:"createDateTextA,omitempty"`
	MapLng                   *string      `json:"mapLng,omitempty"`
	OnLineEnvCount           *int         `json:"onLineEnvCount,omitempty"`
	EventMessBeanList        *interface{} `json:"eventMessBeanList,omitempty"`
	LatitudeText             *string      `json:"latitudeText,omitempty"`
	PlantAddress             *string      `json:"plantAddress,omitempty"`
	HasDeviceOnLine          *int         `json:"hasDeviceOnLine,omitempty"`
	FormulaMoneyStr          *string      `json:"formulaMoneyStr,omitempty"`
	EtodayMoney              *float64     `json:"etodayMoney,omitempty"`
	CreateDate               *APIDateTime `json:"createDate,omitempty"`
	MapCity                  *string      `json:"mapCity,omitempty"`
	PrMonth                  *string      `json:"prMonth,omitempty"`
	StorageTodayToGrid       *float64     `json:"storage_TodayToGrid,omitempty"`
	FormulaCo2               *float64     `json:"formulaCo2,omitempty"`
	ETotal                   *float64     `json:"eTotal,omitempty"`
	EmonthSo2Text            *string      `json:"emonthSo2Text,omitempty"`
	WindAngle                *float64     `json:"windAngle,omitempty"`
	EtotalCoalText           *string      `json:"etotalCoalText,omitempty"`
	WindSpeed                *float64     `json:"windSpeed,omitempty"`
	EmonthCoalText           *string      `json:"emonthCoalText,omitempty"`
	EtodayMoneyText          *string      `json:"etodayMoneyText,omitempty"`
	EYearMoneyText           *string      `json:"EYearMoneyText,omitempty"`
	PlantLng                 *string      `json:"plant_lng,omitempty"`
	LatitudeM                *string      `json:"latitude_m,omitempty"`
	PairViewUserAccount      *string      `json:"pairViewUserAccount,omitempty"`
	StorageTotalToUser       *float64     `json:"storage_TotalToUser,omitempty"`
	LatitudeD                *string      `json:"latitude_d,omitempty"`
	LatitudeF                *string      `json:"latitude_f,omitempty"`
	Remark                   *string      `json:"remark,omitempty"`
	TreeID                   *string      `json:"treeID,omitempty"`
	FlatPeriodPrice          *float64     `json:"flatPeriodPrice,omitempty"`
	LongitudeText            *string      `json:"longitudeText,omitempty"`
	StorageEChargeToday      *float64     `json:"storage_eChargeToday,omitempty"`
	DataLogList              *interface{} `json:"dataLogList,omitempty"`
	DesignCompany            *string      `json:"designCompany,omitempty"`
	TimezoneText             *string      `json:"timezoneText,omitempty"`
	FormulaCoal              *float64     `json:"formulaCoal,omitempty"`
	StorageEDisChargeToday   *float64     `json:"storage_eDisChargeToday,omitempty"`
	UnitMap                  *interface{} `json:"unitMap,omitempty"`
	Timezone                 *int         `json:"timezone,omitempty"`
	PhoneNum                 *string      `json:"phoneNum,omitempty"`
	Level                    *int         `json:"level,omitempty"`
	FormulaMoneyUnitID       *string      `json:"formulaMoneyUnitId,omitempty"`
	ImgPath                  *string      `json:"imgPath,omitempty"`
	PanelTemp                *float64     `json:"panelTemp,omitempty"`
	LocationImgName          *string      `json:"locationImgName,omitempty"`
	MoneyUnitText            *string      `json:"moneyUnitText,omitempty"`
	StorageTotalToGrid       *float64     `json:"storage_TotalToGrid,omitempty"`
	PrToday                  *string      `json:"prToday,omitempty"`
	EnergyMonth              *float64     `json:"energyMonth,omitempty"`
	PlantName                *string      `json:"plantName,omitempty"`
	EToday                   *float64     `json:"eToday,omitempty"`
	Status                   *int         `json:"status,omitempty"`
	PlantType                *int         `json:"plantType,omitempty"`
	Country                  *string      `json:"country,omitempty"`
	LongitudeD               *string      `json:"longitude_d,omitempty"`
	MapAreaID                *int         `json:"map_areaId,omitempty"`
	LongitudeF               *string      `json:"longitude_f,omitempty"`
	CreateDateText           *string      `json:"createDateText,omitempty"`
	LongitudeM               *string      `json:"longitude_m,omitempty"`
	FormulaSo2               *float64     `json:"formulaSo2,omitempty"`
	ValleyPeriodPrice        *float64     `json:"valleyPeriodPrice,omitempty"`
	EnergyYear               *float64     `json:"energyYear,omitempty"`
	TreeName                 *string      `json:"treeName,omitempty"`
	PlantLat                 *string      `json:"plant_lat,omitempty"`
	EtodayCo2Text            *string      `json:"etodayCo2Text,omitempty"`
	NominalPowerStr          *string      `json:"nominalPowerStr,omitempty"`
	FormulaTree              *float64     `json:"formulaTree,omitempty"`
	EtotalSo2Text            *string      `json:"etotalSo2Text,omitempty"`
	Children                 *interface{} `json:"children,omitempty"`
	ID                       *int         `json:"id,omitempty"`
	EtodayCoalText           *string      `json:"etodayCoalText,omitempty"`
	ParamBean                *interface{} `json:"paramBean,omitempty"`
	EtotalMoney              *int         `json:"etotalMoney,omitempty"`
	EnvTemp                  *int         `json:"envTemp,omitempty"`
	LogoImgName              *string      `json:"logoImgName,omitempty"`
	Alias                    *string      `json:"alias,omitempty"`
	EtotalCo2Text            *string      `json:"etotalCo2Text,omitempty"`
	CurrentPacStr            *string      `json:"currentPacStr,omitempty"`
	MapProvinceID            *int         `json:"map_provinceId,omitempty"`
	EtotalMoneyText          *string      `json:"etotalMoneyText,omitempty"`
	EmonthCo2Text            *string      `json:"emonthCo2Text,omitempty"`
	Irradiance               *int         `json:"irradiance,omitempty"`
	HasStorage               *int         `json:"hasStorage,omitempty"`
	ParentID                 *string      `json:"parentID,omitempty"`
	UserBean                 *interface{} `json:"userBean,omitempty"`
	StorageTodayToUser       *int         `json:"storage_TodayToUser,omitempty"`
	IsShare                  *bool        `json:"isShare,omitempty"`
	CurrentPac               *int         `json:"currentPac,omitempty"`
}

func (obj APIPeakPowerActual) Result() PeakPowerActual {
	return PeakPowerActual{
		MapCityID:                pointy.IntValue(obj.MapCityID, 0),
		MapCountryID:             pointy.IntValue(obj.MapCountryID, 0),
		StorageBatteryPercentage: pointy.Float64Value(obj.StorageBatteryPercentage, 0),
		DefaultPlant:             pointy.BoolValue(obj.DefaultPlant, false),
		PeakPeriodPrice:          pointy.Float64Value(obj.PeakPeriodPrice, 0),
		City:                     pointy.StringValue(obj.City, ""),
		NominalPower:             pointy.Float64Value(obj.NominalPower, 0),
		AlarmValue:               pointy.Float64Value(obj.AlarmValue, 0),
		CurrentPacTxt:            pointy.StringValue(obj.CurrentPacTxt, ""),
		FixedPowerPrice:          pointy.Float64Value(obj.FixedPowerPrice, 0),
		PlantFromBean:            pointy.PointerValue(obj.PlantFromBean, nil),
		DeviceCount:              pointy.IntValue(obj.DeviceCount, 0),
		PlantImgName:             pointy.StringValue(obj.PlantImgName, ""),
		EtodaySo2Text:            pointy.StringValue(obj.EtodaySo2Text, ""),
		CompanyName:              pointy.StringValue(obj.CompanyName, ""),
		EmonthMoneyText:          pointy.StringValue(obj.EmonthMoneyText, ""),
		FormulaMoney:             pointy.Float64Value(obj.FormulaMoney, 0),
		UserAccount:              pointy.StringValue(obj.UserAccount, ""),
		MapLat:                   pointy.StringValue(obj.MapLat, ""),
		CreateDateTextA:          pointy.StringValue(obj.CreateDateTextA, ""),
		MapLng:                   pointy.StringValue(obj.MapLng, ""),
		OnLineEnvCount:           pointy.IntValue(obj.OnLineEnvCount, 0),
		EventMessBeanList:        pointy.PointerValue(obj.EventMessBeanList, nil),
		LatitudeText:             pointy.StringValue(obj.LatitudeText, ""),
		PlantAddress:             pointy.StringValue(obj.PlantAddress, ""),
		HasDeviceOnLine:          pointy.IntValue(obj.HasDeviceOnLine, 0),
		FormulaMoneyStr:          pointy.StringValue(obj.FormulaMoneyStr, ""),
		EtodayMoney:              pointy.Float64Value(obj.EtodayMoney, 0),
		CreateDate:               obj.CreateDate.Result(),
		MapCity:                  pointy.StringValue(obj.MapCity, ""),
		PrMonth:                  pointy.StringValue(obj.PrMonth, ""),
		StorageTodayToGrid:       pointy.Float64Value(obj.StorageTodayToGrid, 0),
		FormulaCo2:               pointy.Float64Value(obj.FormulaCo2, 0),
		ETotal:                   pointy.Float64Value(obj.ETotal, 0),
		EmonthSo2Text:            pointy.StringValue(obj.EmonthSo2Text, ""),
		WindAngle:                pointy.Float64Value(obj.WindAngle, 0),
		EtotalCoalText:           pointy.StringValue(obj.EtotalCoalText, ""),
		WindSpeed:                pointy.Float64Value(obj.WindSpeed, 0),
		EmonthCoalText:           pointy.StringValue(obj.EmonthCoalText, ""),
		EtodayMoneyText:          pointy.StringValue(obj.EtodayMoneyText, ""),
		EYearMoneyText:           pointy.StringValue(obj.EYearMoneyText, ""),
		PlantLng:                 pointy.StringValue(obj.PlantLng, ""),
		LatitudeM:                pointy.StringValue(obj.LatitudeM, ""),
		PairViewUserAccount:      pointy.StringValue(obj.PairViewUserAccount, ""),
		StorageTotalToUser:       pointy.Float64Value(obj.StorageTotalToUser, 0),
		LatitudeD:                pointy.StringValue(obj.LatitudeD, ""),
		LatitudeF:                pointy.StringValue(obj.LatitudeF, ""),
		Remark:                   pointy.StringValue(obj.Remark, ""),
		TreeID:                   pointy.StringValue(obj.TreeID, ""),
		FlatPeriodPrice:          pointy.Float64Value(obj.FlatPeriodPrice, 0),
		LongitudeText:            pointy.StringValue(obj.LongitudeText, ""),
		StorageEChargeToday:      pointy.Float64Value(obj.StorageEChargeToday, 0),
		DataLogList:              pointy.PointerValue(obj.DataLogList, nil),
		DesignCompany:            pointy.StringValue(obj.DesignCompany, ""),
		TimezoneText:             pointy.StringValue(obj.TimezoneText, ""),
		FormulaCoal:              pointy.Float64Value(obj.FormulaCoal, 0),
		StorageEDisChargeToday:   pointy.Float64Value(obj.StorageEDisChargeToday, 0),
		UnitMap:                  pointy.PointerValue(obj.UnitMap, nil),
		Timezone:                 pointy.IntValue(obj.Timezone, 0),
		PhoneNum:                 pointy.StringValue(obj.PhoneNum, ""),
		Level:                    pointy.IntValue(obj.Level, 0),
		FormulaMoneyUnitID:       pointy.StringValue(obj.FormulaMoneyUnitID, ""),
		ImgPath:                  pointy.StringValue(obj.ImgPath, ""),
		PanelTemp:                pointy.Float64Value(obj.PanelTemp, 0),
		LocationImgName:          pointy.StringValue(obj.LocationImgName, ""),
		MoneyUnitText:            pointy.StringValue(obj.MoneyUnitText, ""),
		StorageTotalToGrid:       pointy.Float64Value(obj.StorageTodayToGrid, 0),
		PrToday:                  pointy.StringValue(obj.PrToday, ""),
		EnergyMonth:              pointy.Float64Value(obj.EnergyMonth, 0),
		PlantName:                pointy.StringValue(obj.PlantName, ""),
		EToday:                   pointy.Float64Value(obj.EToday, 0),
		Status:                   pointy.IntValue(obj.Status, 0),
		PlantType:                pointy.IntValue(obj.PlantType, 0),
		Country:                  pointy.StringValue(obj.Country, ""),
		LongitudeD:               pointy.StringValue(obj.LongitudeD, ""),
		MapAreaID:                pointy.IntValue(obj.MapAreaID, 0),
		LongitudeF:               pointy.StringValue(obj.LongitudeF, ""),
		CreateDateText:           pointy.StringValue(obj.CreateDateText, ""),
		LongitudeM:               pointy.StringValue(obj.LongitudeM, ""),
		FormulaSo2:               pointy.Float64Value(obj.FormulaSo2, 0),
		ValleyPeriodPrice:        pointy.Float64Value(obj.ValleyPeriodPrice, 0),
		EnergyYear:               pointy.Float64Value(obj.EnergyYear, 0),
		TreeName:                 pointy.StringValue(obj.TreeName, ""),
		PlantLat:                 pointy.StringValue(obj.PlantLat, ""),
		EtodayCo2Text:            pointy.StringValue(obj.EtodayCo2Text, ""),
		NominalPowerStr:          pointy.StringValue(obj.NominalPowerStr, ""),
		FormulaTree:              pointy.Float64Value(obj.FormulaTree, 0),
		EtotalSo2Text:            pointy.StringValue(obj.EtodaySo2Text, ""),
		Children:                 pointy.PointerValue(obj.Children, nil),
		ID:                       pointy.IntValue(obj.ID, 0),
		EtodayCoalText:           pointy.StringValue(obj.EtodayCoalText, ""),
		ParamBean:                pointy.PointerValue(obj.ParamBean, nil),
		EtotalMoney:              pointy.IntValue(obj.EtotalMoney, 0),
		EnvTemp:                  pointy.IntValue(obj.EnvTemp, 0),
		LogoImgName:              pointy.StringValue(obj.LogoImgName, ""),
		Alias:                    pointy.StringValue(obj.Alias, ""),
		EtotalCo2Text:            pointy.StringValue(obj.EtotalCo2Text, ""),
		CurrentPacStr:            pointy.StringValue(obj.CurrentPacStr, ""),
		MapProvinceID:            pointy.IntValue(obj.MapProvinceID, 0),
		EtotalMoneyText:          pointy.StringValue(obj.EtotalMoneyText, ""),
		EmonthCo2Text:            pointy.StringValue(obj.EmonthCo2Text, ""),
		Irradiance:               pointy.IntValue(obj.Irradiance, 0),
		HasStorage:               pointy.IntValue(obj.HasStorage, 0),
		ParentID:                 pointy.StringValue(obj.ParentID, ""),
		UserBean:                 pointy.PointerValue(obj.UserBean, nil),
		StorageTodayToUser:       pointy.IntValue(obj.StorageTodayToUser, 0),
		IsShare:                  pointy.BoolValue(obj.IsShare, false),
		CurrentPac:               pointy.IntValue(obj.CurrentPac, 0),
	}
}

type APIDataLoggerItem struct {
	LastUpdateTime *APIDateTime `json:"last_update_time,omitempty"`
	Model          *string      `json:"model,omitempty"`
	SN             *string      `json:"sn,omitempty"`
	Lost           *bool        `json:"lost,omitempty"`
	Manufacturer   *string      `json:"manufacturer,omitempty"`
	Type           *int         `json:"type,omitempty"`
}

func (obj APIDataLoggerItem) Result() DataLoggerItem {
	return DataLoggerItem{
		LastUpdateTime: obj.LastUpdateTime.Result(),
		Model:          pointy.StringValue(obj.Model, ""),
		SN:             pointy.StringValue(obj.SN, ""),
		Lost:           pointy.BoolValue(obj.Lost, false),
		Manufacturer:   pointy.StringValue(obj.Manufacturer, ""),
		Type:           pointy.IntValue(obj.Type, 0),
	}
}

type APIDateTime struct {
	Time           *int64 `json:"time,omitempty"`
	Minutes        *int   `json:"minutes,omitempty"`
	Seconds        *int   `json:"seconds,omitempty"`
	Hours          *int   `json:"hours,omitempty"`
	Month          *int   `json:"month,omitempty"`
	TimezoneOffset *int   `json:"timezoneOffset,omitempty"`
	Year           *int   `json:"year,omitempty"`
	Day            *int   `json:"day,omitempty"`
	Date           *int   `json:"date,omitempty"`
}

func (obj APIDateTime) Result() DateTime {
	return DateTime{
		Time:           pointy.Int64Value(obj.Time, 0),
		Minutes:        pointy.IntValue(obj.Minutes, 0),
		Seconds:        pointy.IntValue(obj.Seconds, 0),
		Hours:          pointy.IntValue(obj.Hours, 0),
		Month:          pointy.IntValue(obj.Month, 0),
		TimezoneOffset: pointy.IntValue(obj.TimezoneOffset, 0),
		Year:           pointy.IntValue(obj.Year, 0),
		Day:            pointy.IntValue(obj.Day, 0),
		Date:           pointy.IntValue(obj.Date, 0),
	}
}

// ==============================
// API Fetch: GetPlantDataLoggerInfo
// Result Model: GetPlantDeviceListResponse, GetPlantDeviceListData, DeviceItem
// ==============================
type APIGetPlantDeviceListResponse struct {
	Data *APIGetPlantDeviceListData `json:"data,omitempty"`
	APIDefaultResponse
}

func (obj APIGetPlantDeviceListResponse) Result() GetPlantDeviceListResponse {
	return GetPlantDeviceListResponse{
		Data:            obj.Data.Result(),
		DefaultResponse: obj.APIDefaultResponse.Result(),
	}
}

type APIGetPlantDeviceListData struct {
	Count   *int            `json:"count,omitempty"`
	Devices []APIDeviceItem `json:"devices,omitempty"`
}

func (obj APIGetPlantDeviceListData) Result() GetPlantDeviceListData {
	devices := []DeviceItem{}
	for _, item := range obj.Devices {
		devices = append(devices, item.Result())
	}

	return GetPlantDeviceListData{
		Count:   pointy.IntValue(obj.Count, 0),
		Devices: devices,
	}
}

type APIDeviceItem struct {
	DeviceSN       *string `json:"device_sn,omitempty"`
	LastUpdateTime *string `json:"last_update_time,omitempty"`
	Model          *string `json:"model,omitempty"`
	Lost           *bool   `json:"lost,omitempty"`
	Status         *int    `json:"status,omitempty"`
	Manufacturer   *string `json:"manufacturer,omitempty"`
	DeviceID       *int    `json:"device_id,omitempty"`
	DataLoggerSN   *string `json:"datalogger_sn,omitempty"`
	Type           *int    `json:"type,omitempty"`
}

func (obj APIDeviceItem) Result() DeviceItem {
	return DeviceItem{
		DeviceSN:       pointy.StringValue(obj.DeviceSN, ""),
		LastUpdateTime: pointy.StringValue(obj.LastUpdateTime, ""),
		Model:          pointy.StringValue(obj.Model, ""),
		Lost:           pointy.BoolValue(obj.Lost, false),
		Status:         pointy.IntValue(obj.Status, 0),
		Manufacturer:   pointy.StringValue(obj.Manufacturer, ""),
		DeviceID:       pointy.IntValue(obj.DeviceID, 0),
		DataLoggerSN:   pointy.StringValue(obj.DataLoggerSN, ""),
		Type:           pointy.IntValue(obj.Type, 0),
	}
}

// ==============================
// API Fetch: GetRealtimeDeviceData
// Result Model: GetRealtimeDeviceDataResponse, RealtimeDeviceData
// ==============================
type APIGetRealtimeDeviceDataResponse struct {
	DeviceSN     *string                `json:"device_sn,omitempty"`
	DataLoggerSN *string                `json:"dataloggerSn,omitempty"`
	Data         *APIRealtimeDeviceData `json:"data,omitempty"`
	APIDefaultResponse
}

func (obj APIGetRealtimeDeviceDataResponse) Result() GetRealtimeDeviceDataResponse {
	return GetRealtimeDeviceDataResponse{
		DeviceSN:        pointy.StringValue(obj.DeviceSN, ""),
		DataLoggerSN:    pointy.StringValue(obj.DataLoggerSN, ""),
		Data:            obj.Data.Result(),
		DefaultResponse: obj.APIDefaultResponse.Result(),
	}
}

type APIRealtimeDeviceData struct {
	IPidPvcpe             *float64         `json:"iPidPvcpe,omitempty"`
	Epv4Total             *float64         `json:"epv4Total,omitempty"`
	RealOPPercent         *float64         `json:"realOPPercent,omitempty"`
	PidBus                *float64         `json:"pidBus,omitempty"`
	Ppv7                  *float64         `json:"ppv7,omitempty"`
	Ctharis               *float64         `json:"ctharis,omitempty"`
	Ctir                  *float64         `json:"ctir,omitempty"`
	VacTr                 *float64         `json:"vacTr,omitempty"`
	Ppv6                  *float64         `json:"ppv6,omitempty"`
	ERacTotal             *float64         `json:"eRacTotal,omitempty"`
	Ctharit               *float64         `json:"ctharit,omitempty"`
	Ctis                  *float64         `json:"ctis,omitempty"`
	Ppv9                  *float64         `json:"ppv9,omitempty"`
	Epv1Total             *float64         `json:"epv1Total,omitempty"`
	WStringStatusValue    *float64         `json:"wStringStatusValue,omitempty"`
	Ppv8                  *float64         `json:"ppv8,omitempty"`
	Ctharir               *float64         `json:"ctharir,omitempty"`
	WarningValue3         *float64         `json:"warningValue3,omitempty"`
	VPidPvape             *float64         `json:"vPidPvape,omitempty"`
	WarningValue1         *float64         `json:"warningValue1,omitempty"`
	Ctit                  *float64         `json:"ctit,omitempty"`
	FaultCode1            *float64         `json:"faultCode1,omitempty"`
	WarningValue2         *float64         `json:"warningValue2,omitempty"`
	Temperature           *float64         `json:"temperature,omitempty"`
	FaultCode2            *float64         `json:"faultCode2,omitempty"`
	Time                  *string          `json:"time,omitempty"`
	IPidPvbpe             *float64         `json:"iPidPvbpe,omitempty"`
	IPidPvdpe             *float64         `json:"iPidPvdpe,omitempty"`
	Epv2Total             *float64         `json:"epv2Total,omitempty"`
	WarnBit               *float64         `json:"warnBit,omitempty"`
	IPidPvepe             *float64         `json:"iPidPvepe,omitempty"`
	VacSt                 *float64         `json:"vacSt,omitempty"`
	VPidPvcpe             *float64         `json:"vPidPvcpe,omitempty"`
	Epv8Total             *float64         `json:"epv8Total,omitempty"`
	Ipv10                 *float64         `json:"ipv10,omitempty"`
	Again                 *bool            `json:"again,omitempty"`
	Vpv10                 *float64         `json:"vpv10,omitempty"`
	StrBreak              *float64         `json:"strBreak,omitempty"`
	Compqt                *float64         `json:"compqt,omitempty"`
	IpmTemperature        *float64         `json:"ipmTemperature,omitempty"`
	Compqs                *float64         `json:"compqs,omitempty"`
	Ppv                   *float64         `json:"ppv,omitempty"`
	Compqr                *float64         `json:"compqr,omitempty"`
	Ctqt                  *float64         `json:"ctqt,omitempty"`
	Pf                    *float64         `json:"pf,omitempty"`
	Epv7Total             *float64         `json:"epv7Total,omitempty"`
	Vpv1                  *float64         `json:"vpv1,omitempty"`
	Epv10Today            *float64         `json:"epv10Today,omitempty"`
	IPidPvape             *float64         `json:"iPidPvape,omitempty"`
	Vpv3                  *float64         `json:"vpv3,omitempty"`
	Ctqr                  *float64         `json:"ctqr,omitempty"`
	Vpv2                  *float64         `json:"vpv2,omitempty"`
	Ctqs                  *float64         `json:"ctqs,omitempty"`
	Vpv5                  *float64         `json:"vpv5,omitempty"`
	Vpv4                  *float64         `json:"vpv4,omitempty"`
	Vpv7                  *float64         `json:"vpv7,omitempty"`
	Vpv6                  *float64         `json:"vpv6,omitempty"`
	PowerTotal            *float64         `json:"powerTotal,omitempty"`
	Vpv9                  *float64         `json:"vpv9,omitempty"`
	TDci                  *float64         `json:"tDci,omitempty"`
	Vpv8                  *float64         `json:"vpv8,omitempty"`
	Epv2Today             *float64         `json:"epv2Today,omitempty"`
	TimeTotal             *float64         `json:"timeTotal,omitempty"`
	Epv1Today             *float64         `json:"epv1Today,omitempty"`
	Epv6Today             *float64         `json:"epv6Today,omitempty"`
	Epv9Today             *float64         `json:"epv9Today,omitempty"`
	TimeTotalText         *string          `json:"timeTotalText,omitempty"`
	DwStringWarningValue1 *float64         `json:"dwStringWarningValue1,omitempty"`
	VPidPvepe             *float64         `json:"vPidPvepe,omitempty"`
	EpvTotal              *float64         `json:"epvTotal,omitempty"`
	VPidPvgpe             *float64         `json:"vPidPvgpe,omitempty"`
	FaultType             *float64         `json:"faultType,omitempty"`
	CurrentString12       *float64         `json:"currentString12,omitempty"`
	CurrentString11       *float64         `json:"currentString11,omitempty"`
	CurrentString10       *float64         `json:"currentString10,omitempty"`
	ERacToday             *float64         `json:"eRacToday,omitempty"`
	CurrentString16       *float64         `json:"currentString16,omitempty"`
	Epv5Today             *float64         `json:"epv5Today,omitempty"`
	CurrentString15       *float64         `json:"currentString15,omitempty"`
	CurrentString14       *float64         `json:"currentString14,omitempty"`
	CurrentString13       *float64         `json:"currentString13,omitempty"`
	WPIDFaultValue        *float64         `json:"wPIDFaultValue,omitempty"`
	VString11             *float64         `json:"vString11,omitempty"`
	VString10             *float64         `json:"vString10,omitempty"`
	PowerToday            *float64         `json:"powerToday,omitempty"`
	VString16             *float64         `json:"vString16,omitempty"`
	VString13             *float64         `json:"vString13,omitempty"`
	VString12             *float64         `json:"vString12,omitempty"`
	VString15             *float64         `json:"vString15,omitempty"`
	VString14             *float64         `json:"vString14,omitempty"`
	BigDevice             *bool            `json:"bigDevice,omitempty"`
	Epv9Total             *float64         `json:"epv9Total,omitempty"`
	WarnCode              *float64         `json:"warnCode,omitempty"`
	PvIso                 *float64         `json:"pvIso,omitempty"`
	Epv6Total             *float64         `json:"epv6Total,omitempty"`
	InverterID            *string          `json:"inverterId,omitempty"`
	Temperature3          *float64         `json:"temperature3,omitempty"`
	Temperature2          *float64         `json:"temperature2,omitempty"`
	TimeCalendar          *APITimeCalendar `json:"timeCalendar,omitempty"`
	PBusVoltage           *float64         `json:"pBusVoltage,omitempty"`
	CurrentString5        *float64         `json:"currentString5,omitempty"`
	StrFault              *float64         `json:"strFault,omitempty"`
	CurrentString4        *float64         `json:"currentString4,omitempty"`
	VPidPvdpe             *float64         `json:"vPidPvdpe,omitempty"`
	CurrentString3        *float64         `json:"currentString3,omitempty"`
	Epv3Today             *float64         `json:"epv3Today,omitempty"`
	CurrentString2        *float64         `json:"currentString2,omitempty"`
	CurrentString9        *float64         `json:"currentString9,omitempty"`
	Status                *float64         `json:"status,omitempty"`
	CurrentString8        *float64         `json:"currentString8,omitempty"`
	CurrentString7        *float64         `json:"currentString7,omitempty"`
	CurrentString6        *float64         `json:"currentString6,omitempty"`
	NBusVoltage           *float64         `json:"nBusVoltage,omitempty"`
	CurrentString1        *float64         `json:"currentString1,omitempty"`
	Pacs                  *float64         `json:"pacs,omitempty"`
	Pacr                  *float64         `json:"pacr,omitempty"`
	StrUnblance           *float64         `json:"strUnblance,omitempty"`
	StrUnmatch            *float64         `json:"strUnmatch,omitempty"`
	SDci                  *float64         `json:"sDci,omitempty"`
	Pact                  *float64         `json:"pact,omitempty"`
	Fac                   *float64         `json:"fac,omitempty"`
	VPidPvbpe             *float64         `json:"vPidPvbpe,omitempty"`
	FaultValue            *float64         `json:"faultValue,omitempty"`
	Epv5Total             *float64         `json:"epv5Total,omitempty"`
	Ipv6                  *float64         `json:"ipv6,omitempty"`
	Ipv5                  *float64         `json:"ipv5,omitempty"`
	Ipv4                  *float64         `json:"ipv4,omitempty"`
	Epv4Today             *float64         `json:"epv4Today,omitempty"`
	Ipv3                  *float64         `json:"ipv3,omitempty"`
	Ipv2                  *float64         `json:"ipv2,omitempty"`
	Ipv1                  *float64         `json:"ipv1,omitempty"`
	IPidPvfpe             *float64         `json:"iPidPvfpe,omitempty"`
	StatusText            *string          `json:"statusText,omitempty"`
	VacRs                 *float64         `json:"vacRs,omitempty"`
	IPidPvgpe             *float64         `json:"iPidPvgpe,omitempty"`
	Ipv9                  *float64         `json:"ipv9,omitempty"`
	Ipv8                  *float64         `json:"ipv8,omitempty"`
	Ipv7                  *float64         `json:"ipv7,omitempty"`
	ID                    *float64         `json:"id,omitempty"`
	Epv8Today             *float64         `json:"epv8Today,omitempty"`
	Gfci                  *float64         `json:"gfci,omitempty"`
	IPidPvhpe             *float64         `json:"iPidPvhpe,omitempty"`
	Ppv10                 *float64         `json:"ppv10,omitempty"`
	Epv3Total             *float64         `json:"epv3Total,omitempty"`
	ApfStatus             *float64         `json:"apfStatus,omitempty"`
	Temperature4          *float64         `json:"temperature4,omitempty"`
	RDci                  *float64         `json:"rDci,omitempty"`
	Pac                   *float64         `json:"pac,omitempty"`
	Temperature5          *float64         `json:"temperature5,omitempty"`
	Vact                  *float64         `json:"vact,omitempty"`
	Compharir             *float64         `json:"compharir,omitempty"`
	Vacr                  *float64         `json:"vacr,omitempty"`
	Compharis             *float64         `json:"compharis,omitempty"`
	Vacs                  *float64         `json:"vacs,omitempty"`
	PidFaultCode          *float64         `json:"pidFaultCode,omitempty"`
	Compharit             *float64         `json:"compharit,omitempty"`
	DeratingMode          *float64         `json:"deratingMode,omitempty"`
	VString1              *float64         `json:"vString1,omitempty"`
	Epv7Today             *float64         `json:"epv7Today,omitempty"`
	VString2              *float64         `json:"vString2,omitempty"`
	VString3              *float64         `json:"vString3,omitempty"`
	VPidPvhpe             *float64         `json:"vPidPvhpe,omitempty"`
	VString4              *float64         `json:"vString4,omitempty"`
	VString5              *float64         `json:"vString5,omitempty"`
	VString6              *float64         `json:"vString6,omitempty"`
	VString8              *float64         `json:"vString8,omitempty"`
	PidStatus             *float64         `json:"pidStatus,omitempty"`
	Iacs                  *float64         `json:"iacs,omitempty"`
	OpFullwatt            *float64         `json:"opFullwatt,omitempty"`
	VString7              *float64         `json:"vString7,omitempty"`
	Iact                  *float64         `json:"iact,omitempty"`
	VString9              *float64         `json:"vString9,omitempty"`
	Epv10Total            *float64         `json:"epv10Total,omitempty"`
	VPidPvfpe             *float64         `json:"vPidPvfpe,omitempty"`
	Ppv5                  *float64         `json:"ppv5,omitempty"`
	Debug1                *string          `json:"debug1,omitempty"`
	Ppv4                  *float64         `json:"ppv4,omitempty"`
	Debug2                *string          `json:"debug2,omitempty"`
	Ppv3                  *float64         `json:"ppv3,omitempty"`
	Ppv2                  *float64         `json:"ppv2,omitempty"`
	Ppv1                  *float64         `json:"ppv1,omitempty"`
	Rac                   *float64         `json:"rac,omitempty"`
	Iacr                  *float64         `json:"iacr,omitempty"`
}

func (obj APIRealtimeDeviceData) Result() RealtimeDeviceData {
	return RealtimeDeviceData{
		IPidPvcpe:             pointy.Float64Value(obj.IPidPvcpe, 0),
		Epv4Total:             pointy.Float64Value(obj.Epv4Total, 0),
		RealOPPercent:         pointy.Float64Value(obj.RealOPPercent, 0),
		PidBus:                pointy.Float64Value(obj.PidBus, 0),
		Ppv7:                  pointy.Float64Value(obj.Ppv7, 0),
		Ctharis:               pointy.Float64Value(obj.Ctharis, 0),
		Ctir:                  pointy.Float64Value(obj.Ctir, 0),
		VacTr:                 pointy.Float64Value(obj.VacTr, 0),
		Ppv6:                  pointy.Float64Value(obj.Ppv6, 0),
		ERacTotal:             pointy.Float64Value(obj.ERacTotal, 0),
		Ctharit:               pointy.Float64Value(obj.Ctharit, 0),
		Ctis:                  pointy.Float64Value(obj.Ctis, 0),
		Ppv9:                  pointy.Float64Value(obj.Ppv9, 0),
		Epv1Total:             pointy.Float64Value(obj.Epv1Total, 0),
		WStringStatusValue:    pointy.Float64Value(obj.WStringStatusValue, 0),
		Ppv8:                  pointy.Float64Value(obj.Ppv8, 0),
		Ctharir:               pointy.Float64Value(obj.Ctharir, 0),
		WarningValue3:         pointy.Float64Value(obj.WarningValue3, 0),
		VPidPvape:             pointy.Float64Value(obj.VPidPvape, 0),
		WarningValue1:         pointy.Float64Value(obj.WarningValue1, 0),
		Ctit:                  pointy.Float64Value(obj.Ctit, 0),
		FaultCode1:            pointy.Float64Value(obj.FaultCode1, 0),
		WarningValue2:         pointy.Float64Value(obj.WarningValue2, 0),
		Temperature:           pointy.Float64Value(obj.Temperature, 0),
		FaultCode2:            pointy.Float64Value(obj.FaultCode1, 0),
		Time:                  pointy.StringValue(obj.Time, ""),
		IPidPvbpe:             pointy.Float64Value(obj.IPidPvbpe, 0),
		IPidPvdpe:             pointy.Float64Value(obj.IPidPvdpe, 0),
		Epv2Total:             pointy.Float64Value(obj.Epv2Total, 0),
		WarnBit:               pointy.Float64Value(obj.WarnBit, 0),
		IPidPvepe:             pointy.Float64Value(obj.IPidPvepe, 0),
		VacSt:                 pointy.Float64Value(obj.VacSt, 0),
		VPidPvcpe:             pointy.Float64Value(obj.VPidPvcpe, 0),
		Epv8Total:             pointy.Float64Value(obj.Epv8Total, 0),
		Ipv10:                 pointy.Float64Value(obj.Ipv10, 0),
		Again:                 pointy.BoolValue(obj.Again, false),
		Vpv10:                 pointy.Float64Value(obj.Vpv10, 0),
		StrBreak:              pointy.Float64Value(obj.StrBreak, 0),
		Compqt:                pointy.Float64Value(obj.Compqt, 0),
		IpmTemperature:        pointy.Float64Value(obj.IpmTemperature, 0),
		Compqs:                pointy.Float64Value(obj.Compqs, 0),
		Ppv:                   pointy.Float64Value(obj.Ppv, 0),
		Compqr:                pointy.Float64Value(obj.Compqr, 0),
		Ctqt:                  pointy.Float64Value(obj.Ctqt, 0),
		Pf:                    pointy.Float64Value(obj.Pf, 0),
		Epv7Total:             pointy.Float64Value(obj.Epv7Total, 0),
		Vpv1:                  pointy.Float64Value(obj.Vpv1, 0),
		Epv10Today:            pointy.Float64Value(obj.Epv10Today, 0),
		IPidPvape:             pointy.Float64Value(obj.IPidPvape, 0),
		Vpv3:                  pointy.Float64Value(obj.Vpv3, 0),
		Ctqr:                  pointy.Float64Value(obj.Ctqr, 0),
		Vpv2:                  pointy.Float64Value(obj.Vpv2, 0),
		Ctqs:                  pointy.Float64Value(obj.Ctqs, 0),
		Vpv5:                  pointy.Float64Value(obj.Vpv5, 0),
		Vpv4:                  pointy.Float64Value(obj.Vpv4, 0),
		Vpv7:                  pointy.Float64Value(obj.Vpv7, 0),
		Vpv6:                  pointy.Float64Value(obj.Vpv6, 0),
		PowerTotal:            pointy.Float64Value(obj.PowerTotal, 0),
		Vpv9:                  pointy.Float64Value(obj.Vpv9, 0),
		TDci:                  pointy.Float64Value(obj.TDci, 0),
		Vpv8:                  pointy.Float64Value(obj.Vpv8, 0),
		Epv2Today:             pointy.Float64Value(obj.Epv2Today, 0),
		TimeTotal:             pointy.Float64Value(obj.TimeTotal, 0),
		Epv1Today:             pointy.Float64Value(obj.Epv1Today, 0),
		Epv6Today:             pointy.Float64Value(obj.Epv6Today, 0),
		Epv9Today:             pointy.Float64Value(obj.Epv9Today, 0),
		TimeTotalText:         pointy.StringValue(obj.TimeTotalText, ""),
		DwStringWarningValue1: pointy.Float64Value(obj.DwStringWarningValue1, 0),
		VPidPvepe:             pointy.Float64Value(obj.VPidPvepe, 0),
		EpvTotal:              pointy.Float64Value(obj.EpvTotal, 0),
		VPidPvgpe:             pointy.Float64Value(obj.VPidPvgpe, 0),
		FaultType:             pointy.Float64Value(obj.FaultType, 0),
		CurrentString12:       pointy.Float64Value(obj.CurrentString12, 0),
		CurrentString11:       pointy.Float64Value(obj.CurrentString11, 0),
		CurrentString10:       pointy.Float64Value(obj.CurrentString10, 0),
		ERacToday:             pointy.Float64Value(obj.ERacToday, 0),
		CurrentString16:       pointy.Float64Value(obj.CurrentString16, 0),
		Epv5Today:             pointy.Float64Value(obj.Epv5Today, 0),
		CurrentString15:       pointy.Float64Value(obj.CurrentString15, 0),
		CurrentString14:       pointy.Float64Value(obj.CurrentString14, 0),
		CurrentString13:       pointy.Float64Value(obj.CurrentString13, 0),
		WPIDFaultValue:        pointy.Float64Value(obj.WPIDFaultValue, 0),
		VString11:             pointy.Float64Value(obj.VString11, 0),
		VString10:             pointy.Float64Value(obj.VString10, 0),
		PowerToday:            pointy.Float64Value(obj.PowerToday, 0),
		VString16:             pointy.Float64Value(obj.VString16, 0),
		VString13:             pointy.Float64Value(obj.VString13, 0),
		VString12:             pointy.Float64Value(obj.VString12, 0),
		VString15:             pointy.Float64Value(obj.VString15, 0),
		VString14:             pointy.Float64Value(obj.VString14, 0),
		BigDevice:             pointy.BoolValue(obj.BigDevice, false),
		Epv9Total:             pointy.Float64Value(obj.Epv9Total, 0),
		WarnCode:              pointy.Float64Value(obj.WarnCode, 0),
		PvIso:                 pointy.Float64Value(obj.PvIso, 0),
		Epv6Total:             pointy.Float64Value(obj.Epv6Total, 0),
		InverterID:            pointy.StringValue(obj.InverterID, ""),
		Temperature3:          pointy.Float64Value(obj.Temperature3, 0),
		Temperature2:          pointy.Float64Value(obj.Temperature2, 0),
		TimeCalendar:          obj.TimeCalendar.Result(),
		PBusVoltage:           pointy.Float64Value(obj.PBusVoltage, 0),
		CurrentString5:        pointy.Float64Value(obj.CurrentString5, 0),
		StrFault:              pointy.Float64Value(obj.StrFault, 0),
		CurrentString4:        pointy.Float64Value(obj.CurrentString4, 0),
		VPidPvdpe:             pointy.Float64Value(obj.VPidPvdpe, 0),
		CurrentString3:        pointy.Float64Value(obj.CurrentString3, 0),
		Epv3Today:             pointy.Float64Value(obj.Epv3Today, 0),
		CurrentString2:        pointy.Float64Value(obj.CurrentString2, 0),
		CurrentString9:        pointy.Float64Value(obj.CurrentString9, 0),
		Status:                pointy.Float64Value(obj.Status, 0),
		CurrentString8:        pointy.Float64Value(obj.CurrentString8, 0),
		CurrentString7:        pointy.Float64Value(obj.CurrentString7, 0),
		CurrentString6:        pointy.Float64Value(obj.CurrentString6, 0),
		NBusVoltage:           pointy.Float64Value(obj.NBusVoltage, 0),
		CurrentString1:        pointy.Float64Value(obj.CurrentString1, 0),
		Pacs:                  pointy.Float64Value(obj.Pacs, 0),
		Pacr:                  pointy.Float64Value(obj.Pacr, 0),
		StrUnblance:           pointy.Float64Value(obj.StrUnblance, 0),
		StrUnmatch:            pointy.Float64Value(obj.StrUnmatch, 0),
		SDci:                  pointy.Float64Value(obj.SDci, 0),
		Pact:                  pointy.Float64Value(obj.Pact, 0),
		Fac:                   pointy.Float64Value(obj.Fac, 0),
		VPidPvbpe:             pointy.Float64Value(obj.VPidPvbpe, 0),
		FaultValue:            pointy.Float64Value(obj.FaultValue, 0),
		Epv5Total:             pointy.Float64Value(obj.Epv5Total, 0),
		Ipv6:                  pointy.Float64Value(obj.Ipv6, 0),
		Ipv5:                  pointy.Float64Value(obj.Ipv5, 0),
		Ipv4:                  pointy.Float64Value(obj.Ipv4, 0),
		Epv4Today:             pointy.Float64Value(obj.Epv4Today, 0),
		Ipv3:                  pointy.Float64Value(obj.Ipv3, 0),
		Ipv2:                  pointy.Float64Value(obj.Ipv2, 0),
		Ipv1:                  pointy.Float64Value(obj.Ipv1, 0),
		IPidPvfpe:             pointy.Float64Value(obj.IPidPvfpe, 0),
		StatusText:            pointy.StringValue(obj.StatusText, ""),
		VacRs:                 pointy.Float64Value(obj.VacRs, 0),
		IPidPvgpe:             pointy.Float64Value(obj.IPidPvgpe, 0),
		Ipv9:                  pointy.Float64Value(obj.Ipv9, 0),
		Ipv8:                  pointy.Float64Value(obj.Ipv8, 0),
		Ipv7:                  pointy.Float64Value(obj.Ipv7, 0),
		ID:                    pointy.Float64Value(obj.ID, 0),
		Epv8Today:             pointy.Float64Value(obj.Epv8Today, 0),
		Gfci:                  pointy.Float64Value(obj.Gfci, 0),
		IPidPvhpe:             pointy.Float64Value(obj.IPidPvhpe, 0),
		Ppv10:                 pointy.Float64Value(obj.Ppv10, 0),
		Epv3Total:             pointy.Float64Value(obj.Epv3Total, 0),
		ApfStatus:             pointy.Float64Value(obj.ApfStatus, 0),
		Temperature4:          pointy.Float64Value(obj.Temperature4, 0),
		RDci:                  pointy.Float64Value(obj.RDci, 0),
		Pac:                   pointy.Float64Value(obj.Pac, 0),
		Temperature5:          pointy.Float64Value(obj.Temperature5, 0),
		Vact:                  pointy.Float64Value(obj.Vact, 0),
		Compharir:             pointy.Float64Value(obj.Compharir, 0),
		Vacr:                  pointy.Float64Value(obj.Vacr, 0),
		Compharis:             pointy.Float64Value(obj.Compharis, 0),
		Vacs:                  pointy.Float64Value(obj.Vacs, 0),
		PidFaultCode:          pointy.Float64Value(obj.PidFaultCode, 0),
		Compharit:             pointy.Float64Value(obj.Compharit, 0),
		DeratingMode:          pointy.Float64Value(obj.DeratingMode, 0),
		VString1:              pointy.Float64Value(obj.VString1, 0),
		Epv7Today:             pointy.Float64Value(obj.Epv7Today, 0),
		VString2:              pointy.Float64Value(obj.VString2, 0),
		VString3:              pointy.Float64Value(obj.VString3, 0),
		VPidPvhpe:             pointy.Float64Value(obj.VPidPvhpe, 0),
		VString4:              pointy.Float64Value(obj.VString4, 0),
		VString5:              pointy.Float64Value(obj.VString5, 0),
		VString6:              pointy.Float64Value(obj.VString6, 0),
		VString8:              pointy.Float64Value(obj.VString8, 0),
		PidStatus:             pointy.Float64Value(obj.PidStatus, 0),
		Iacs:                  pointy.Float64Value(obj.Iacs, 0),
		OpFullwatt:            pointy.Float64Value(obj.OpFullwatt, 0),
		VString7:              pointy.Float64Value(obj.VString7, 0),
		Iact:                  pointy.Float64Value(obj.Iact, 0),
		VString9:              pointy.Float64Value(obj.VString9, 0),
		Epv10Total:            pointy.Float64Value(obj.Epv10Total, 0),
		VPidPvfpe:             pointy.Float64Value(obj.VPidPvfpe, 0),
		Ppv5:                  pointy.Float64Value(obj.Ppv4, 0),
		Debug1:                pointy.StringValue(obj.Debug1, ""),
		Ppv4:                  pointy.Float64Value(obj.Ppv4, 0),
		Debug2:                pointy.StringValue(obj.Debug2, ""),
		Ppv3:                  pointy.Float64Value(obj.Ppv3, 0),
		Ppv2:                  pointy.Float64Value(obj.Ppv2, 0),
		Ppv1:                  pointy.Float64Value(obj.Ppv1, 0),
		Rac:                   pointy.Float64Value(obj.Rac, 0),
		Iacr:                  pointy.Float64Value(obj.Iacr, 0),
	}
}

type APITimeCalendar struct {
	MinimalDaysInFirstWeek *int         `json:"minimalDaysInFirstWeek,omitempty"`
	WeekYear               *int         `json:"weekYear,omitempty"`
	Time                   *APIDateTime `json:"time,omitempty"`
	WeeksInWeekYear        *int         `json:"weeksInWeekYear,omitempty"`
	GregorianChange        *APIDateTime `json:"gregorianChange,omitempty"`
	TimeZone               *APITimeZone `json:"timeZone,omitempty"`
	TimeInMillis           *int64       `json:"timeInMillis,omitempty"`
	Lenient                *bool        `json:"lenient,omitempty"`
	FirstDayOfWeek         *int         `json:"firstDayOfWeek,omitempty"`
	WeekDateSupported      *bool        `json:"weekDateSupported,omitempty"`
}

func (obj APITimeCalendar) Result() TimeCalendar {
	return TimeCalendar{
		MinimalDaysInFirstWeek: pointy.IntValue(obj.MinimalDaysInFirstWeek, 0),
		WeekYear:               pointy.IntValue(obj.WeekYear, 0),
		Time:                   obj.Time.Result(),
		WeeksInWeekYear:        pointy.IntValue(obj.WeeksInWeekYear, 0),
		GregorianChange:        obj.GregorianChange.Result(),
		TimeZone:               obj.TimeZone.Result(),
		TimeInMillis:           pointy.Int64Value(obj.TimeInMillis, 0),
		Lenient:                pointy.BoolValue(obj.Lenient, false),
		FirstDayOfWeek:         pointy.IntValue(obj.FirstDayOfWeek, 0),
		WeekDateSupported:      pointy.BoolValue(obj.WeekDateSupported, false),
	}
}

type APITimeZone struct {
	LastRuleInstance *interface{} `json:"lastRuleInstance,omitempty"`
	RawOffset        *int         `json:"rawOffset,omitempty"`
	DSTSavings       *int         `json:"DSTSavings,omitempty"`
	Dirty            *bool        `json:"dirty,omitempty"`
	ID               *string      `json:"ID,omitempty"`
	DisplayName      *string      `json:"displayName,omitempty"`
}

func (obj APITimeZone) Result() TimeZone {
	return TimeZone{
		LastRuleInstance: pointy.PointerValue(obj.LastRuleInstance, nil),
		RawOffset:        pointy.IntValue(obj.RawOffset, 0),
		DSTSavings:       pointy.IntValue(obj.DSTSavings, 0),
		Dirty:            pointy.BoolValue(obj.Dirty, false),
		ID:               pointy.StringValue(obj.ID, ""),
		DisplayName:      pointy.StringValue(obj.DisplayName, ""),
	}
}

// ==============================
// API Fetch: GetRealtimeDeviceBatches
// Result Model: GetRealtimeDeviceBatchesDataResponse
// ==============================
type APIGetRealtimeDeviceBatchesDataResponse struct {
	Inverters []string                          `json:"inverters,omitempty"`
	Data      map[string]map[string]interface{} `json:"data,omitempty"`
	PageNum   *int                              `json:"pageNum,omitempty"`
	APIDefaultResponse
}

func (obj APIGetRealtimeDeviceBatchesDataResponse) Result() GetRealtimeDeviceBatchesDataResponse {
	return GetRealtimeDeviceBatchesDataResponse{
		Inverters:       obj.Inverters,
		Data:            obj.Data,
		PageNum:         pointy.IntValue(obj.PageNum, 0),
		DefaultResponse: obj.APIDefaultResponse.Result(),
	}
}

// ==============================
// API Fetch: GetRealtimeDeviceData
// Result Model: GetRealtimeDeviceDataResponse, RealtimeDeviceData
// ==============================
type APIGetInverterAlertListResponse struct {
	Data *APIGetInverterAlertListData `json:"data,omitempty"`
	APIDefaultResponse
}

func (obj APIGetInverterAlertListResponse) Result() GetInverterAlertListResponse {
	return GetInverterAlertListResponse{
		Data:            obj.Data.Result(),
		DefaultResponse: obj.APIDefaultResponse.Result(),
	}
}

type APIGetInverterAlertListData struct {
	SN     *string        `json:"sn,omitempty"`
	Count  *int           `json:"count,omitempty"`
	Alarms []APIAlarmItem `json:"alarms,omitempty"`
}

func (obj APIGetInverterAlertListData) Result() GetInverterAlertListData {
	alarms := []AlarmItem{}
	for _, item := range obj.Alarms {
		alarms = append(alarms, item.Result())
	}

	return GetInverterAlertListData{
		SN:     pointy.StringValue(obj.SN, ""),
		Count:  pointy.IntValue(obj.Count, 0),
		Alarms: nil,
	}
}

// ==============================
// API Fetch: GetRealtimeDeviceData
// Result Model: GetRealtimeDeviceDataResponse, RealtimeDeviceData
// ==============================
type APIGetEnergyStorageMachineAlertListResponse struct {
	Data *APIGetEnergyStorageMachineAlertListData `json:"data,omitempty"`
	APIDefaultResponse
}

func (obj APIGetEnergyStorageMachineAlertListResponse) Result() GetEnergyStorageMachineAlertListResponse {
	return GetEnergyStorageMachineAlertListResponse{
		Data:            obj.Data.Result(),
		DefaultResponse: obj.APIDefaultResponse.Result(),
	}
}

type APIGetEnergyStorageMachineAlertListData struct {
	StorageSN *string        `json:"storage_sn,omitempty"`
	Count     *int           `json:"count,omitempty"`
	Alarms    []APIAlarmItem `json:"alarms,omitempty"`
}

func (obj APIGetEnergyStorageMachineAlertListData) Result() GetEnergyStorageMachineAlertListData {
	alarms := []AlarmItem{}
	for _, item := range obj.Alarms {
		alarms = append(alarms, item.Result())
	}

	return GetEnergyStorageMachineAlertListData{
		StorageSN: pointy.StringValue(obj.StorageSN, ""),
		Count:     pointy.IntValue(obj.Count, 0),
		Alarms:    nil,
	}
}

// ==============================
// API Fetch: GetRealtimeDeviceData
// Result Model: GetRealtimeDeviceDataResponse, RealtimeDeviceData
// ==============================
type APIGetMaxAlertListResponse struct {
	Data *APIGetMaxAlertListData `json:"data,omitempty"`
	APIDefaultResponse
}

func (obj APIGetMaxAlertListResponse) Result() GetMaxAlertListResponse {
	return GetMaxAlertListResponse{
		Data:            obj.Data.Result(),
		DefaultResponse: obj.APIDefaultResponse.Result(),
	}
}

type APIGetMaxAlertListData struct {
	MaxSN  *string        `json:"max_sn,omitempty"`
	Count  *int           `json:"count,omitempty"`
	Alarms []APIAlarmItem `json:"alarms,omitempty"`
}

func (obj APIGetMaxAlertListData) Result() GetMaxAlertListData {
	alarms := []AlarmItem{}
	for _, item := range obj.Alarms {
		alarms = append(alarms, item.Result())
	}

	return GetMaxAlertListData{
		MaxSN:  pointy.StringValue(obj.MaxSN, ""),
		Count:  pointy.IntValue(obj.Count, 0),
		Alarms: nil,
	}
}

// ==============================
// API Fetch: GetRealtimeDeviceData
// Result Model: GetRealtimeDeviceDataResponse, RealtimeDeviceData
// ==============================
type APIGetMixAlertListResponse struct {
	Data *APIGetMixAlertListData `json:"data,omitempty"`
	APIDefaultResponse
}

func (obj APIGetMixAlertListResponse) Result() GetMixAlertListResponse {
	return GetMixAlertListResponse{
		Data:            obj.Data.Result(),
		DefaultResponse: obj.APIDefaultResponse.Result(),
	}
}

type APIGetMixAlertListData struct {
	MixSN  *string        `json:"mix_sn,omitempty"`
	Count  *int           `json:"count,omitempty"`
	Alarms []APIAlarmItem `json:"alarms,omitempty"`
}

func (obj APIGetMixAlertListData) Result() GetMixAlertListData {
	alarms := []AlarmItem{}
	for _, item := range obj.Alarms {
		alarms = append(alarms, item.Result())
	}

	return GetMixAlertListData{
		MixSN:  pointy.StringValue(obj.MixSN, ""),
		Count:  pointy.IntValue(obj.Count, 0),
		Alarms: nil,
	}
}

// ==============================
// API Fetch: GetRealtimeDeviceData
// Result Model: GetRealtimeDeviceDataResponse, RealtimeDeviceData
// ==============================
type APIGetSpaAlertListResponse struct {
	Data *APIGetSpaAlertListData `json:"data,omitempty"`
	APIDefaultResponse
}

func (obj APIGetSpaAlertListResponse) Result() GetSpaAlertListResponse {
	return GetSpaAlertListResponse{
		Data:            obj.Data.Result(),
		DefaultResponse: obj.APIDefaultResponse.Result(),
	}
}

type APIGetSpaAlertListData struct {
	SpaSN  *string        `json:"spa_sn,omitempty"`
	Count  *int           `json:"count,omitempty"`
	Alarms []APIAlarmItem `json:"alarms,omitempty"`
}

func (obj APIGetSpaAlertListData) Result() GetSpaAlertListData {
	alarms := []AlarmItem{}
	for _, item := range obj.Alarms {
		alarms = append(alarms, item.Result())
	}

	return GetSpaAlertListData{
		SpaSN:  pointy.StringValue(obj.SpaSN, ""),
		Count:  pointy.IntValue(obj.Count, 0),
		Alarms: nil,
	}
}

// ==============================
// API Fetch: GetRealtimeDeviceData
// Result Model: GetRealtimeDeviceDataResponse, RealtimeDeviceData
// ==============================
type APIGetMinAlertListResponse struct {
	Data *APIGetMinAlertListData `json:"data,omitempty"`
	APIDefaultResponse
}

func (obj APIGetMinAlertListResponse) Result() GetMinAlertListResponse {
	return GetMinAlertListResponse{
		Data:            obj.Data.Result(),
		DefaultResponse: obj.APIDefaultResponse.Result(),
	}
}

type APIGetMinAlertListData struct {
	TlxSn  *string        `json:"tlx_sn,omitempty"`
	Count  *int           `json:"count,omitempty"`
	Alarms []APIAlarmItem `json:"alarms,omitempty"`
}

func (obj APIGetMinAlertListData) Result() GetMinAlertListData {
	alarms := []AlarmItem{}
	for _, item := range obj.Alarms {
		alarms = append(alarms, item.Result())
	}

	return GetMinAlertListData{
		TlxSn:  pointy.StringValue(obj.TlxSn, ""),
		Count:  pointy.IntValue(obj.Count, 0),
		Alarms: nil,
	}
}

// ==============================
// API Fetch: GetRealtimeDeviceData
// Result Model: GetRealtimeDeviceDataResponse, RealtimeDeviceData
// ==============================
type APIGetPcsAlertListResponse struct {
	Data *APIGetPcsAlertListData `json:"data,omitempty"`
	APIDefaultResponse
}

func (obj APIGetPcsAlertListResponse) Result() GetPcsAlertListResponse {
	return GetPcsAlertListResponse{
		Data:            obj.Data.Result(),
		DefaultResponse: obj.APIDefaultResponse.Result(),
	}
}

type APIGetPcsAlertListData struct {
	PcsSN  *string        `json:"pcs_sn,omitempty"`
	Count  *int           `json:"count,omitempty"`
	Alarms []APIAlarmItem `json:"alarms,omitempty"`
}

func (obj APIGetPcsAlertListData) Result() GetPcsAlertListData {
	alarms := []AlarmItem{}
	for _, item := range obj.Alarms {
		alarms = append(alarms, item.Result())
	}

	return GetPcsAlertListData{
		PcsSN:  pointy.StringValue(obj.PcsSN, ""),
		Count:  pointy.IntValue(obj.Count, 0),
		Alarms: nil,
	}
}

// ==============================
// API Fetch: GetRealtimeDeviceData
// Result Model: GetRealtimeDeviceDataResponse, RealtimeDeviceData
// ==============================
type APIGetHpsAlertListResponse struct {
	Data *APIGetHpsAlertListData `json:"data,omitempty"`
	APIDefaultResponse
}

func (obj APIGetHpsAlertListResponse) Result() GetHpsAlertListResponse {
	return GetHpsAlertListResponse{
		Data:            obj.Data.Result(),
		DefaultResponse: obj.APIDefaultResponse.Result(),
	}
}

type APIGetHpsAlertListData struct {
	HpsSN  *string        `json:"hps_sn,omitempty"`
	Count  *int           `json:"count,omitempty"`
	Alarms []APIAlarmItem `json:"alarms,omitempty"`
}

func (obj APIGetHpsAlertListData) Result() GetHpsAlertListData {
	alarms := []AlarmItem{}
	for _, item := range obj.Alarms {
		alarms = append(alarms, item.Result())
	}

	return GetHpsAlertListData{
		HpsSN:  pointy.StringValue(obj.HpsSN, ""),
		Count:  pointy.IntValue(obj.Count, 0),
		Alarms: nil,
	}
}

// ==============================
// API Fetch: GetRealtimeDeviceData
// Result Model: GetRealtimeDeviceDataResponse, RealtimeDeviceData
// ==============================
type APIGetPbdAlertListResponse struct {
	Data *APIGetPbdAlertListData `json:"data,omitempty"`
	APIDefaultResponse
}

func (obj APIGetPbdAlertListResponse) Result() GetPbdAlertListResponse {
	return GetPbdAlertListResponse{
		Data:            obj.Data.Result(),
		DefaultResponse: obj.APIDefaultResponse.Result(),
	}
}

type APIGetPbdAlertListData struct {
	PbdSN  *string        `json:"pbd_sn,omitempty"`
	Count  *int           `json:"count,omitempty"`
	Alarms []APIAlarmItem `json:"alarms,omitempty"`
}

func (obj APIGetPbdAlertListData) Result() GetPbdAlertListData {
	alarms := []AlarmItem{}
	for _, item := range obj.Alarms {
		alarms = append(alarms, item.Result())
	}

	return GetPbdAlertListData{
		PbdSN:  pointy.StringValue(obj.PbdSN, ""),
		Count:  pointy.IntValue(obj.Count, 0),
		Alarms: nil,
	}
}

type APIAlarmItem struct {
	AlarmCode    *int    `json:"alarm_code,omitempty"`
	Status       *int    `json:"status,omitempty"`
	EndTime      *string `json:"end_time,omitempty"`
	StartTime    *string `json:"start_time,omitempty"`
	AlarmMessage *string `json:"alarm_message,omitempty"`
}

func (obj APIAlarmItem) Result() AlarmItem {
	return AlarmItem{
		AlarmCode:    pointy.IntValue(obj.AlarmCode, 0),
		Status:       pointy.IntValue(obj.Status, 0),
		EndTime:      pointy.StringValue(obj.EndTime, ""),
		StartTime:    pointy.StringValue(obj.StartTime, ""),
		AlarmMessage: pointy.StringValue(obj.AlarmMessage, ""),
	}
}
