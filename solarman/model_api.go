package solarman

import "github.com/openlyinc/pointy"

// ===== REQUEST BODY =====
type PaginationBody struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

type StationBody struct {
	StationID int `json:"stationId"`
}

type DeviceBody struct {
	DeviceSerialNumber string `json:"deviceSn"`
}

type GetTokenBody struct {
	Username       string `json:"username,omitempty"`
	Password       string `json:"password"`
	AppSecret      string `json:"appSecret"`
	CountryCode    string `json:"countryCode,omitempty"`
	Email          string `json:"email,omitempty"`
	Mobile         string `json:"mobile,omitempty"`
	OrganizationID int    `json:"orgId,omitempty"`
}

// ===========================
// API Fetch: GetToken
// Result Model: GetTokenResponse
// ===========================

type APIGetTokenResponse struct {
	Code         *string `json:"code,omitempty"`
	Message      *string `json:"msg,omitempty"`
	Success      *bool   `json:"success,omitempty"`
	RequestID    *string `json:"requestId,omitempty"`
	Uid          *int    `json:"uid,omitempty"`
	TokenType    *string `json:"token_type,omitempty"`
	Scope        *string `json:"scope,omitempty"`
	AccessToken  *string `json:"access_token,omitempty"`
	RefreshToken *string `json:"refresh_token,omitempty"`
	ExpiresIn    *string `json:"expires_in,omitempty"`
}

func (obj APIGetTokenResponse) Result() GetTokenResponse {
	return GetTokenResponse{
		Code:         pointy.StringValue(obj.Code, ""),
		Message:      pointy.StringValue(obj.Message, ""),
		Success:      pointy.BoolValue(obj.Success, false),
		RequestID:    pointy.StringValue(obj.RequestID, ""),
		Uid:          pointy.IntValue(obj.Uid, 0),
		TokenType:    pointy.StringValue(obj.TokenType, ""),
		Scope:        pointy.StringValue(obj.Scope, ""),
		AccessToken:  pointy.StringValue(obj.AccessToken, ""),
		RefreshToken: pointy.StringValue(obj.RefreshToken, ""),
		ExpiresIn:    pointy.StringValue(obj.ExpiresIn, ""),
	}
}

// ===========================
// API Fetch: GetUserInfo
// Result Model: GetUserInfoResponse, OrganizationInfoItem
// ===========================
type APIGetUserInfoResponse struct {
	Code        *string                   `json:"code,omitempty"`
	Message     *string                   `json:"msg,omitempty"`
	Success     *bool                     `json:"success,omitempty"`
	RequestID   *string                   `json:"requestId,omitempty"`
	OrgInfoList []APIOrganizationInfoItem `json:"orgInfoList,omitempty"`
}

func (obj APIGetUserInfoResponse) Result() GetUserInfoResponse {
	orgInfoList := []OrganizationInfoItem{}
	for _, item := range obj.OrgInfoList {
		orgInfoList = append(orgInfoList, item.Result())
	}

	return GetUserInfoResponse{
		Code:        pointy.StringValue(obj.Code, ""),
		Message:     pointy.StringValue(obj.Message, ""),
		Success:     pointy.BoolValue(obj.Success, false),
		RequestID:   pointy.StringValue(obj.RequestID, ""),
		OrgInfoList: orgInfoList,
	}
}

type APIOrganizationInfoItem struct {
	CompanyID   *int    `json:"companyId,omitempty"`
	CompanyName *string `json:"companyName,omitempty"`
	RoleName    *string `json:"roleName,omitempty"`
}

func (obj APIOrganizationInfoItem) Result() OrganizationInfoItem {
	return OrganizationInfoItem{
		CompanyID:   pointy.IntValue(obj.CompanyID, 0),
		CompanyName: pointy.StringValue(obj.CompanyName, ""),
		RoleName:    pointy.StringValue(obj.RoleName, ""),
	}
}

// ===========================
// API Fetch: GetPlantList
// Result Model: GetPlantListResponse, PlantItem
// ===========================
type APIGetPlantListResponse struct {
	Code      *string        `json:"code,omitempty"`
	Message   *string        `json:"msg,omitempty"`
	Success   *bool          `json:"success,omitempty"`
	RequestID *string        `json:"requestId,omitempty"`
	Total     *int           `json:"total,omitempty"`
	PlantList []APIPlantItem `json:"stationList,omitempty"`
}

func (obj APIGetPlantListResponse) Result() GetPlantListResponse {
	stationList := []PlantItem{}
	for _, item := range obj.PlantList {
		stationList = append(stationList, item.Result())
	}

	return GetPlantListResponse{
		Code:      pointy.StringValue(obj.Code, ""),
		Message:   pointy.StringValue(obj.Message, ""),
		Success:   pointy.BoolValue(obj.Success, false),
		RequestID: pointy.StringValue(obj.RequestID, ""),
		Total:     pointy.IntValue(obj.Total, 0),
		PlantList: stationList,
	}
}

type APIPlantItem struct {
	ID                      *int     `json:"id,omitempty"`
	Name                    *string  `json:"name,omitempty"`
	LocationLat             *float64 `json:"locationLat,omitempty"`
	LocationLng             *float64 `json:"locationLng,omitempty"`
	LocationAddress         *string  `json:"locationAddress,omitempty"`
	RegionNationID          *int     `json:"regionNationId,omitempty"`
	RegionLevel1            *int     `json:"regionLevel1,omitempty"`
	RegionLevel2            *int     `json:"regionLevel2,omitempty"`
	RegionLevel3            *int     `json:"regionLevel3,omitempty"`
	RegionLevel4            *int     `json:"regionLevel4,omitempty"`
	RegionLevel5            *int     `json:"regionLevel5,omitempty"`
	RegionTimezone          *string  `json:"regionTimezone,omitempty"`
	Type                    *string  `json:"type,omitempty"`
	GridInterconnectionType *string  `json:"gridInterconnectionType,omitempty"`
	InstalledCapacity       *float64 `json:"installedCapacity,omitempty"`
	StartOperatingTime      *float64 `json:"startOperatingTime,omitempty"`
	StationImage            *string  `json:"stationImage,omitempty"`
	CreatedDate             *float64 `json:"createdDate,omitempty"`
	BatterySoc              *float64 `json:"batterySoc,omitempty"`
	NetworkStatus           *string  `json:"networkStatus,omitempty"`
	GenerationPower         *float64 `json:"generationPower,omitempty"`
	LastUpdateTime          *float64 `json:"lastUpdateTime,omitempty"`
}

func (obj APIPlantItem) Result() PlantItem {
	return PlantItem{
		ID:                      pointy.IntValue(obj.ID, 0),
		Name:                    pointy.StringValue(obj.Name, ""),
		LocationLat:             pointy.Float64Value(obj.LocationLat, 0),
		LocationLng:             pointy.Float64Value(obj.LocationLng, 0),
		LocationAddress:         pointy.StringValue(obj.LocationAddress, ""),
		RegionNationID:          pointy.IntValue(obj.RegionNationID, 0),
		RegionLevel1:            pointy.IntValue(obj.RegionLevel1, 0),
		RegionLevel2:            pointy.IntValue(obj.RegionLevel2, 0),
		RegionLevel3:            pointy.IntValue(obj.RegionLevel3, 0),
		RegionLevel4:            pointy.IntValue(obj.RegionLevel4, 0),
		RegionLevel5:            pointy.IntValue(obj.RegionLevel5, 0),
		RegionTimezone:          pointy.StringValue(obj.RegionTimezone, ""),
		Type:                    pointy.StringValue(obj.Type, ""),
		GridInterconnectionType: pointy.StringValue(obj.GridInterconnectionType, ""),
		InstalledCapacity:       pointy.Float64Value(obj.InstalledCapacity, 0),
		StartOperatingTime:      pointy.Float64Value(obj.StartOperatingTime, 0),
		StationImage:            pointy.StringValue(obj.StationImage, ""),
		CreatedDate:             pointy.Float64Value(obj.CreatedDate, 0),
		BatterySoc:              pointy.Float64Value(obj.BatterySoc, 0),
		NetworkStatus:           pointy.StringValue(obj.NetworkStatus, ""),
		GenerationPower:         pointy.Float64Value(obj.GenerationPower, 0),
		LastUpdateTime:          pointy.Float64Value(obj.LastUpdateTime, 0),
	}
}

// ===========================
// API Fetch: GetPlantDeviceList
// Result Model: GetPlantDeviceListResponse, PlantDeviceItem
// ===========================
type GetPlantDeviceListBody struct {
	StationID  int    `json:"stationId"`
	DeviceType string `json:"deviceType,omitempty"`
	Page       int    `json:"page,omitempty"`
	Size       int    `json:"size,omitempty"`
}

type APIGetPlantDeviceListResponse struct {
	Code            *string              `json:"code,omitempty"`
	Message         *string              `json:"msg,omitempty"`
	Success         *bool                `json:"success,omitempty"`
	RequestID       *string              `json:"requestId,omitempty"`
	Total           *int                 `json:"total,omitempty"`
	DeviceListItems []APIPlantDeviceItem `json:"deviceListItems,omitempty"`
}

func (obj APIGetPlantDeviceListResponse) Result() GetPlantDeviceListResponse {
	deviceListItem := []PlantDeviceItem{}
	for _, item := range obj.DeviceListItems {
		deviceListItem = append(deviceListItem, item.Result())
	}

	return GetPlantDeviceListResponse{
		Code:            pointy.StringValue(obj.Code, ""),
		Message:         pointy.StringValue(obj.Message, ""),
		Success:         pointy.BoolValue(obj.Success, false),
		RequestID:       pointy.StringValue(obj.RequestID, ""),
		Total:           pointy.IntValue(obj.Total, 0),
		DeviceListItems: deviceListItem,
	}
}

type APIPlantDeviceItem struct {
	DeviceSN       *string `json:"deviceSn,omitempty"`
	DeviceID       *int    `json:"deviceId,omitempty"`
	DeviceType     *string `json:"deviceType,omitempty"`
	ConnectStatus  *int    `json:"connectStatus,omitempty"`
	CollectionTime *int64  `json:"collectionTime,omitempty"`
}

func (obj APIPlantDeviceItem) Result() PlantDeviceItem {
	return PlantDeviceItem{
		DeviceSN:       pointy.StringValue(obj.DeviceSN, ""),
		DeviceID:       pointy.IntValue(obj.DeviceID, 0),
		DeviceType:     pointy.StringValue(obj.DeviceType, ""),
		ConnectStatus:  pointy.IntValue(obj.ConnectStatus, 0),
		CollectionTime: pointy.Int64Value(obj.CollectionTime, 0),
	}
}

// ===========================
// API Fetch: GetDeviceAlertList
// Result Model: GetDeviceAlertListResponse, DeviceAlertItem
// ===========================
type GetDeviceAlertListBody struct {
	DeviceSN       string `json:"deviceSn"`
	DeviceID       int    `json:"deviceId,omitempty"`
	StartTimestamp int64  `json:"startTimestamp"`
	EndTimestamp   int64  `json:"endTimestamp"`
	Page           int    `json:"page,omitempty"`
	Size           int    `json:"size,omitempty"`
}

type APIGetDeviceAlertListResponse struct {
	Code       *string              `json:"code,omitempty"`
	Message    *string              `json:"msg,omitempty"`
	Success    *bool                `json:"success,omitempty"`
	RequestID  *string              `json:"requestId,omitempty"`
	DeviceSN   *string              `json:"deviceSn,omitempty"`
	DeviceID   *int                 `json:"deviceId,omitempty"`
	DeviceType *string              `json:"deviceType,omitempty"`
	Total      *int                 `json:"total,omitempty"`
	AlertList  []APIDeviceAlertItem `json:"alertList,omitempty"`
}

func (obj APIGetDeviceAlertListResponse) Result() GetDeviceAlertListResponse {
	deviceAlertList := []DeviceAlertItem{}
	for _, item := range obj.AlertList {
		deviceAlertList = append(deviceAlertList, item.Result())
	}

	return GetDeviceAlertListResponse{
		Code:       pointy.StringValue(obj.Code, ""),
		Message:    pointy.StringValue(obj.Message, ""),
		Success:    pointy.BoolValue(obj.Success, false),
		RequestID:  pointy.StringValue(obj.RequestID, ""),
		DeviceSN:   pointy.StringValue(obj.DeviceSN, ""),
		DeviceID:   pointy.IntValue(obj.DeviceID, 0),
		DeviceType: pointy.StringValue(obj.DeviceType, ""),
		Total:      pointy.IntValue(obj.Total, 0),
		AlertList:  deviceAlertList,
	}
}

type APIDeviceAlertItem struct {
	AlertID         *int    `json:"alertId,omitempty"`
	AlertName       *string `json:"addr,omitempty"`
	AlertNameInPAAS *string `json:"alertName,omitempty"`
	Code            *string `json:"code,omitempty"`
	Level           *int    `json:"level,omitempty"`
	Influence       *int    `json:"influence,omitempty"`
	AlertTime       *int64  `json:"alertTime,omitempty"`
}

func (obj APIDeviceAlertItem) Result() DeviceAlertItem {
	return DeviceAlertItem{
		AlertID:         pointy.IntValue(obj.AlertID, 0),
		AlertName:       pointy.StringValue(obj.AlertName, ""),
		AlertNameInPAAS: pointy.StringValue(obj.AlertNameInPAAS, ""),
		Code:            pointy.StringValue(obj.Code, ""),
		Level:           pointy.IntValue(obj.Level, 0),
		Influence:       pointy.IntValue(obj.Influence, 0),
		AlertTime:       pointy.Int64Value(obj.AlertTime, 0),
	}
}

// ===========================
// API Fetch: GetPlantBaseInfo
// Result Model: GetPlantBaseInfoResponse, Region, StationImageItem
// ===========================
type APIGetPlantBaseInfoResponse struct {
	Code                     *string               `json:"code,omitempty"`
	Message                  *string               `json:"msg,omitempty"`
	Success                  *bool                 `json:"success,omitempty"`
	RequestID                *string               `json:"requestId,omitempty"`
	ID                       *int                  `json:"id,omitempty"`
	Name                     *string               `json:"name,omitempty"`
	LocationLat              *float64              `json:"locationLat,omitempty"`
	LocationLng              *float64              `json:"locationLng,omitempty"`
	LocationAddress          *string               `json:"locationAddress,omitempty"`
	Region                   APIRegion             `json:"region,omitempty"`
	Type                     *string               `json:"type,omitempty"`
	GridInterconnectionType  *string               `json:"gridInterconnectionType,omitempty"`
	InstalledCapacity        *float64              `json:"installedCapacity,omitempty"`
	InstallationAzimuthAngle *float64              `json:"installationAzimuthAngle,omitempty"`
	InstallationTiltAngle    *float64              `json:"installationTiltAngle,omitempty"`
	StartOperatingTime       *float64              `json:"startOperatingTime,omitempty"`
	Currency                 *string               `json:"currency,omitempty"`
	OwnerName                *string               `json:"ownerName,omitempty"`
	OwnerCompany             *string               `json:"ownerCompany,omitempty"`
	ContactPhone             *string               `json:"contactPhone,omitempty"`
	MergeElectricPrice       *float64              `json:"mergeElectricPrice,omitempty"`
	ConstructionCost         *float64              `json:"constructionCost,omitempty"`
	StationImage             *string               `json:"stationImage,omitempty"`
	StationImages            []APIStationImageItem `json:"stationImages,omitempty"`
	CreatedDate              *float64              `json:"createdDate,omitempty"`
}

func (obj APIGetPlantBaseInfoResponse) Result() GetPlantBaseInfoResponse {
	stationImages := []StationImageItem{}
	for _, item := range obj.StationImages {
		stationImages = append(stationImages, item.Result())
	}

	return GetPlantBaseInfoResponse{
		Code:                     pointy.StringValue(obj.Code, ""),
		Message:                  pointy.StringValue(obj.Message, ""),
		Success:                  pointy.BoolValue(obj.Success, false),
		RequestID:                pointy.StringValue(obj.RequestID, ""),
		ID:                       pointy.IntValue(obj.ID, 0),
		Name:                     pointy.StringValue(obj.Name, ""),
		LocationLat:              pointy.Float64Value(obj.LocationLat, 0),
		LocationLng:              pointy.Float64Value(obj.LocationLng, 0),
		LocationAddress:          pointy.StringValue(obj.LocationAddress, ""),
		Region:                   obj.Region.Result(),
		Type:                     pointy.StringValue(obj.Type, ""),
		GridInterconnectionType:  pointy.StringValue(obj.GridInterconnectionType, ""),
		InstalledCapacity:        pointy.Float64Value(obj.InstalledCapacity, 0),
		InstallationAzimuthAngle: pointy.Float64Value(obj.InstallationAzimuthAngle, 0),
		InstallationTiltAngle:    pointy.Float64Value(obj.InstallationTiltAngle, 0),
		StartOperatingTime:       pointy.Float64Value(obj.StartOperatingTime, 0),
		Currency:                 pointy.StringValue(obj.Currency, ""),
		OwnerName:                pointy.StringValue(obj.OwnerName, ""),
		OwnerCompany:             pointy.StringValue(obj.OwnerCompany, ""),
		ContactPhone:             pointy.StringValue(obj.ContactPhone, ""),
		MergeElectricPrice:       pointy.Float64Value(obj.MergeElectricPrice, 0),
		ConstructionCost:         pointy.Float64Value(obj.ConstructionCost, 0),
		StationImage:             pointy.StringValue(obj.StationImage, ""),
		StationImages:            stationImages,
		CreatedDate:              pointy.Float64Value(obj.CreatedDate, 0),
	}
}

type APIRegion struct {
	NationID *int    `json:"nationId,omitempty"`
	Level1   *int    `json:"level1,omitempty"`
	Level2   *int    `json:"level2,omitempty"`
	Level3   *int    `json:"level3,omitempty"`
	Level4   *int    `json:"level4,omitempty"`
	Level5   *int    `json:"level5,omitempty"`
	Timezone *string `json:"timezone,omitempty"`
}

func (obj APIRegion) Result() Region {
	return Region{
		NationID: pointy.IntValue(obj.NationID, 0),
		Level1:   pointy.IntValue(obj.Level1, 0),
		Level2:   pointy.IntValue(obj.Level2, 0),
		Level3:   pointy.IntValue(obj.Level3, 0),
		Level4:   pointy.IntValue(obj.Level4, 0),
		Level5:   pointy.IntValue(obj.Level5, 0),
		Timezone: pointy.StringValue(obj.Timezone, ""),
	}
}

type APIStationImageItem struct {
	ID          *int    `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Url         *string `json:"url,omitempty"`
}

func (obj APIStationImageItem) Result() StationImageItem {
	return StationImageItem{
		ID:          pointy.IntValue(obj.ID, 0),
		Name:        pointy.StringValue(obj.Url, ""),
		Description: pointy.StringValue(obj.Description, ""),
		Url:         pointy.StringValue(obj.Url, ""),
	}
}

// ===========================
// API Fetch: GetRealtimePlantData
// Result Model: GetRealtimePlantDataResponse
// ===========================
type APIGetRealtimePlantDataResponse struct {
	Code               *string  `json:"code,omitempty"`
	Message            *string  `json:"msg,omitempty"`
	Success            *bool    `json:"success,omitempty"`
	RequestID          *string  `json:"requestId,omitempty"`
	GenerationPower    *float64 `json:"generationPower,omitempty"`
	UsePower           *float64 `json:"usePower,omitempty"`
	GridPower          *float64 `json:"gridPower,omitempty"`
	PurchasePower      *float64 `json:"purchasePower,omitempty"`
	WirePower          *float64 `json:"wirePower,omitempty"`
	ChargePower        *float64 `json:"chargePower,omitempty"`
	DischargePower     *float64 `json:"dischargePower,omitempty"`
	BatterPower        *float64 `json:"batteryPower,omitempty"`
	BatterSoc          *float64 `json:"batterySoc,omitempty"`
	IrradiateIntensity *float64 `json:"irradiateIntensity,omitempty"`
	LastUpdateTime     *float64 `json:"lastUpdateTime,omitempty"`
}

func (obj APIGetRealtimePlantDataResponse) Result() GetRealtimePlantDataResponse {
	return GetRealtimePlantDataResponse{
		Code:               pointy.StringValue(obj.Code, ""),
		Message:            pointy.StringValue(obj.Message, ""),
		Success:            pointy.BoolValue(obj.Success, false),
		RequestID:          pointy.StringValue(obj.RequestID, ""),
		GenerationPower:    pointy.Float64Value(obj.GenerationPower, 0),
		UsePower:           pointy.Float64Value(obj.UsePower, 0),
		GridPower:          pointy.Float64Value(obj.GridPower, 0),
		PurchasePower:      pointy.Float64Value(obj.PurchasePower, 0),
		WirePower:          pointy.Float64Value(obj.WirePower, 0),
		ChargePower:        pointy.Float64Value(obj.ChargePower, 0),
		DischargePower:     pointy.Float64Value(obj.DischargePower, 0),
		BatterPower:        pointy.Float64Value(obj.BatterPower, 0),
		BatterSoc:          pointy.Float64Value(obj.BatterSoc, 0),
		IrradiateIntensity: pointy.Float64Value(obj.IrradiateIntensity, 0),
		LastUpdateTime:     pointy.Float64Value(obj.LastUpdateTime, 0),
	}
}

// ===========================
// API Fetch: GetHistoricalPlantData
// Result Model: GetHistoricalPlantDataResponse, HistoricalStationDataItem
// ===========================
type GetHistoricalPlantDataBody struct {
	StationID int    `json:"stationId"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	TimeType  int    `json:"timeType"`
}

type APIGetHistoricalPlantDataResponse struct {
	Code             *string                        `json:"code,omitempty"`
	Message          *string                        `json:"msg,omitempty"`
	Success          *bool                          `json:"success,omitempty"`
	RequestID        *string                        `json:"requestId,omitempty"`
	Total            *int                           `json:"total,omitempty"`
	StationDataItems []APIHistoricalStationDataItem `json:"stationDataItems,omitempty"`
}

func (obj APIGetHistoricalPlantDataResponse) Result() GetHistoricalPlantDataResponse {
	stationDataItems := []HistoricalStationDataItem{}
	for _, item := range obj.StationDataItems {
		stationDataItems = append(stationDataItems, item.Result())
	}

	return GetHistoricalPlantDataResponse{
		Code:             pointy.StringValue(obj.Code, ""),
		Message:          pointy.StringValue(obj.Message, ""),
		Success:          pointy.BoolValue(obj.Success, false),
		RequestID:        pointy.StringValue(obj.RequestID, ""),
		Total:            pointy.IntValue(obj.Total, 0),
		StationDataItems: stationDataItems,
	}
}

type APIHistoricalStationDataItem struct {
	GenerationPower       *float64 `json:"generationPower,omitempty"`
	UsePower              *float64 `json:"usePower,omitempty"`
	GridPower             *float64 `json:"gridPower,omitempty"`
	PurchasePower         *float64 `json:"purchasePower,omitempty"`
	WirePower             *float64 `json:"wirePower,omitempty"`
	ChargePower           *float64 `json:"chargePower,omitempty"`
	DischargePower        *float64 `json:"dischargePower,omitempty"`
	BatterPower           *float64 `json:"batteryPower,omitempty"`
	BatterSoc             *float64 `json:"batterySoc,omitempty"`
	IrradiateIntensity    *float64 `json:"irradiateIntensity,omitempty"`
	GenerationValue       *float64 `json:"generationValue,omitempty"`
	GenerationRation      *float64 `json:"generationRatio,omitempty"`
	GridRatio             *float64 `json:"gridRatio,omitempty"`
	ChargeRatio           *float64 `json:"chargeRatio,omitempty"`
	UseValue              *float64 `json:"useValue,omitempty"`
	UseRatio              *float64 `json:"useRatio,omitempty"`
	BuyRatio              *float64 `json:"buyRatio,omitempty"`
	UseDischargeRatio     *float64 `json:"useDischargeRatio,omitempty"`
	GridValue             *float64 `json:"gridValue,omitempty"`
	BuyValue              *float64 `json:"buyValue,omitempty"`
	ChargeValue           *float64 `json:"chargeValue,omitempty"`
	DischargeValue        *float64 `json:"dischargeValue,omitempty"`
	FullPowerHours        *float64 `json:"fullPowerHours,omitempty"`
	Irradiate             *float64 `json:"irradiate,omitempty"`
	TheoreticalGeneration *float64 `json:"theoreticalGeneration,omitempty"`
	PR                    *float64 `json:"pr,omitempty"`
	CPR                   *float64 `json:"cpr,omitempty"`
	DateTime              *float64 `json:"dateTime,omitempty"`
	Year                  *int     `json:"year,omitempty"`
	Month                 *int     `json:"month,omitempty"`
	Day                   *int     `json:"day,omitempty"`
}

func (obj APIHistoricalStationDataItem) Result() HistoricalStationDataItem {
	return HistoricalStationDataItem{
		GenerationPower:       pointy.Float64Value(obj.GenerationRation, 0),
		UsePower:              pointy.Float64Value(obj.UsePower, 0),
		GridPower:             pointy.Float64Value(obj.GridPower, 0),
		PurchasePower:         pointy.Float64Value(obj.PurchasePower, 0),
		WirePower:             pointy.Float64Value(obj.WirePower, 0),
		ChargePower:           pointy.Float64Value(obj.ChargePower, 0),
		DischargePower:        pointy.Float64Value(obj.DischargePower, 0),
		BatterPower:           pointy.Float64Value(obj.BatterPower, 0),
		BatterSoc:             pointy.Float64Value(obj.BatterSoc, 0),
		IrradiateIntensity:    pointy.Float64Value(obj.IrradiateIntensity, 0),
		GenerationValue:       pointy.Float64Value(obj.GenerationValue, 0),
		GenerationRation:      pointy.Float64Value(obj.GenerationRation, 0),
		GridRatio:             pointy.Float64Value(obj.GridRatio, 0),
		ChargeRatio:           pointy.Float64Value(obj.ChargeRatio, 0),
		UseValue:              pointy.Float64Value(obj.UseValue, 0),
		UseRatio:              pointy.Float64Value(obj.UseRatio, 0),
		BuyRatio:              pointy.Float64Value(obj.BuyRatio, 0),
		UseDischargeRatio:     pointy.Float64Value(obj.UseDischargeRatio, 0),
		GridValue:             pointy.Float64Value(obj.GridValue, 0),
		BuyValue:              pointy.Float64Value(obj.BuyValue, 0),
		ChargeValue:           pointy.Float64Value(obj.ChargeValue, 0),
		DischargeValue:        pointy.Float64Value(obj.DischargeValue, 0),
		FullPowerHours:        pointy.Float64Value(obj.FullPowerHours, 0),
		Irradiate:             pointy.Float64Value(obj.Irradiate, 0),
		TheoreticalGeneration: pointy.Float64Value(obj.TheoreticalGeneration, 0),
		PR:                    pointy.Float64Value(obj.PR, 0),
		CPR:                   pointy.Float64Value(obj.CPR, 0),
		DateTime:              pointy.Float64Value(obj.DateTime, 0),
		Year:                  pointy.IntValue(obj.Year, 0),
		Month:                 pointy.IntValue(obj.Month, 0),
		Day:                   pointy.IntValue(obj.Day, 0),
	}
}

// ===========================
// API Fetch: GetRealtimeDeviceData
// Result Model: GetRealtimeDeviceDataResponse, DataListItem
// ===========================
type APIGetRealtimeDeviceDataResponse struct {
	Code        *string           `json:"code,omitempty"`
	Message     *string           `json:"msg,omitempty"`
	Success     *bool             `json:"success,omitempty"`
	RequestID   *string           `json:"requestId,omitempty"`
	DeviceSN    *string           `json:"deviceSn,omitempty"`
	DeviceID    *int              `json:"deviceId,omitempty"`
	DeviceType  *string           `json:"deviceType,omitempty"`
	DeviceState *int              `json:"deviceState,omitempty"`
	DataList    []APIDataListItem `json:"dataList,omitempty"`
}

func (obj APIGetRealtimeDeviceDataResponse) Result() GetRealtimeDeviceDataResponse {
	dataList := []DataListItem{}
	for _, item := range obj.DataList {
		dataList = append(dataList, item.Result())
	}

	return GetRealtimeDeviceDataResponse{
		Code:        pointy.StringValue(obj.Code, ""),
		Message:     pointy.StringValue(obj.Message, ""),
		Success:     pointy.BoolValue(obj.Success, false),
		RequestID:   pointy.StringValue(obj.RequestID, ""),
		DeviceSN:    pointy.StringValue(obj.DeviceSN, ""),
		DeviceID:    pointy.IntValue(obj.DeviceID, 0),
		DeviceType:  pointy.StringValue(obj.DeviceType, ""),
		DeviceState: pointy.IntValue(obj.DeviceState, 0),
		DataList:    dataList,
	}
}

type APIDataListItem struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
	Unit  *string `json:"unit,omitempty"`
	Name  *string `json:"name,omitempty"`
}

func (obj APIDataListItem) Result() DataListItem {
	return DataListItem{
		Key:   pointy.StringValue(obj.Key, ""),
		Value: pointy.StringValue(obj.Value, ""),
		Unit:  pointy.StringValue(obj.Unit, ""),
		Name:  pointy.StringValue(obj.Name, ""),
	}
}

type GetHistoricalDeviceDataBody struct {
	DeviceSN  string `json:"deviceSn"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	TimeType  int    `json:"timeType"`
}

type APIGetHistoricalDeviceDataResponse struct {
	Code          *string            `json:"code,omitempty"`
	Message       *string            `json:"msg,omitempty"`
	Success       *bool              `json:"success,omitempty"`
	RequestID     *string            `json:"requestId,omitempty"`
	DeviceSN      *string            `json:"deviceSn,omitempty"`
	DeviceID      *int               `json:"deviceId,omitempty"`
	DeviceType    *string            `json:"deviceType,omitempty"`
	TimeType      *int               `json:"timeType,omitempty"`
	ParamDataList []APIParamDataItem `json:"paramDataList,omitempty"`
}

func (obj APIGetHistoricalDeviceDataResponse) Result() GetHistoricalDeviceDataResponse {
	paramList := []ParamDataItem{}
	for _, item := range obj.ParamDataList {
		paramList = append(paramList, item.Result())
	}

	return GetHistoricalDeviceDataResponse{
		Code:          pointy.StringValue(obj.Code, ""),
		Message:       pointy.StringValue(obj.Message, ""),
		Success:       pointy.BoolValue(obj.Success, false),
		RequestID:     pointy.StringValue(obj.RequestID, ""),
		DeviceSN:      pointy.StringValue(obj.DeviceSN, ""),
		DeviceID:      pointy.IntValue(obj.DeviceID, 0),
		DeviceType:    pointy.StringValue(obj.DeviceType, ""),
		TimeType:      pointy.IntValue(obj.TimeType, 0),
		ParamDataList: paramList,
	}
}

type APIParamDataItem struct {
	CollectTime *string           `json:"collectTime,omitempty"`
	DataList    []APIDataListItem `json:"dataList,omitempty"`
}

func (obj APIParamDataItem) Result() ParamDataItem {
	dataList := []DataListItem{}
	for _, item := range obj.DataList {
		dataList = append(dataList, item.Result())
	}
	return ParamDataItem{
		CollectTime: pointy.StringValue(obj.CollectTime, ""),
		DataList:    dataList,
	}
}
