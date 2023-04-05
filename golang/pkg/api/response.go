package api

type ApiResponse struct {
	StatusCode int
	Data       interface{}
	Error      ApiErrorResponse
}

type ApiErrorResponse struct {
	Code    string
	Message string
}

func NewApiResponse(statusCode int, data interface{}) ApiResponse {
	return ApiResponse{
		StatusCode: statusCode,
		Data:       data,
	}
}

func NewApiErrorResponse(statusCode int, message string, errorCode string) ApiResponse {
	return ApiResponse{
		StatusCode: statusCode,
		Error: ApiErrorResponse{
			Message: message,
			Code:    errorCode,
		},
	}
}

func NewApiErrorResponseFromError(err ApiError) ApiResponse {
	return ApiResponse{
		StatusCode: err.StatusCode(),
		Error: ApiErrorResponse{
			Message: err.Error(),
			Code:    err.ErrorCode(),
		},
	}
}
