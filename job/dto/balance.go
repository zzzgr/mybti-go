package dto

// MyBtiBalanceResponse myBti抢票响应
type MyBtiBalanceResponse struct {
	Balance         int    `json:"balance"`
	AppointmentId   string `json:"appointmentId"`
	StationEntrance string `json:"stationEntrance"`
}
