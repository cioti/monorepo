package api

type ApiResponse interface {
	StatusCode() int
	Payload() interface{}
}

type apiResponse struct {
	statusCode int
	payload    interface{}
}

func (r *apiResponse) StatusCode() int {
	return r.statusCode
}

func (r *apiResponse) Payload() interface{} {
	return r.payload
}

func NewApiResponse(statusCode int, payload interface{}) ApiResponse {
	return &apiResponse{
		statusCode: statusCode,
		payload:    payload,
	}
}
