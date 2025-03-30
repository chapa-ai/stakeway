//package response
//
//type CreateValidatorResponse struct {
//	RequestID string `json:"request_id"`
//	Message   string `json:"message"`
//}
//
//type GetValidatorStatusResponse struct {
//	Status string   `json:"status"`
//	Keys   []string `json:"keys,omitempty"`
//}
//
//type ErrorResponse struct {
//	Error   string `json:"error"`
//	Details string `json:"details,omitempty"`
//}
//
//func NewCreateValidatorResponse(requestID string) CreateValidatorResponse {
//	return CreateValidatorResponse{
//		RequestID: requestID,
//		Message:   "Validator creation in progress",
//	}
//}
//
//func NewGetValidatorStatusResponse(status string, keys []string) GetValidatorStatusResponse {
//	return GetValidatorStatusResponse{
//		Status: status,
//		Keys:   keys,
//	}
//}
//
//func NewErrorResponse(message string, details ...string) ErrorResponse {
//	var detail string
//	if len(details) > 0 {
//		detail = details[0]
//	}
//	return ErrorResponse{
//		Error:   message,
//		Details: detail,
//	}
//}

package response

type CreateValidatorResponse struct {
	RequestID string `json:"request_id"`
	Message   string `json:"message"`
}

type GetValidatorStatusResponse struct {
	Status string   `json:"status"`
	Keys   []string `json:"keys,omitempty"`
}

type FailedResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func NewCreateValidatorResponse(requestID string) CreateValidatorResponse {
	return CreateValidatorResponse{
		RequestID: requestID,
		Message:   "Validator creation in progress",
	}
}

func NewGetValidatorStatusResponse(status string, keys []string) GetValidatorStatusResponse {
	return GetValidatorStatusResponse{
		Status: status,
		Keys:   keys,
	}
}

func NewFailedResponse(message string) FailedResponse {
	return FailedResponse{
		Status:  "failed",
		Message: message,
	}
}
