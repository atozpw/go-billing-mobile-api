package models

type ResponseOnlyMessage struct {
	Code    int    `json:"responseCode"`
	Message string `json:"responseMessage"`
}
