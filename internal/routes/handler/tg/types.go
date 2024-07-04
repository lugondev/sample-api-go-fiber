package tele_handler

type RequestTgAuth struct {
	Data     string `json:"data"`
	DeviceId string `json:"device_id"`
	Platform string `json:"platform"`
}
