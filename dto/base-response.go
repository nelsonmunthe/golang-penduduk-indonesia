package dto

type BaseResponseList struct {
	PreviousPage interface{} `json:"previousPage"`
	CurrentPage  int         `json:"currentPage"`
	NextPage     interface{} `json:"nextPage"`
	Total        int64       `json:"total"`
	PerPage      int         `json:"perPage"`
	Success      bool        `json:"success"`
	MessageTitle string      `json:"messageTitle"`
	Message      string      `json:"message"`
	Data         interface{} `json:"data"`
}

type BaseResponse struct {
	Data         interface{} `json:"data"`
	Success      bool        `json:"success"`
	MessageTitle string      `json:"messageTitle"`
	Message      string      `json:"message"`
}

func DefaultErrorBaseResponseList(err error) BaseResponseList {
	return BaseResponseList{
		Data:         nil,
		Success:      false,
		MessageTitle: "",
		Message:      err.Error(),
	}
}

func DefaultErrorBaseResponse(err error) BaseResponse {
	return BaseResponse{
		Data:         nil,
		Success:      false,
		MessageTitle: "",
		Message:      err.Error(),
	}
}
func DefaultErrorBaseResponseWithMessage(message string) BaseResponse {
	return BaseResponse{
		Data:         nil,
		Success:      false,
		MessageTitle: "",
		Message:      message,
	}
}
