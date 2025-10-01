package dto

type ResponseMeta struct {
	Success      bool        `json:"success"`
	MessageTitle string      `json:"messageTitle"`
	Message      interface{} `json:"message"`
	ResponseTime string      `json:"responseTime"`
	Data         interface{} `json:"data"`
	StatusCode   uint        `json:"statusCode"`
}
