package models

type ResponseWithData struct {
	Code    int         `json:"responseCode"`
	Message string      `json:"responseMessage"`
	Data    interface{} `json:"data"`
}
