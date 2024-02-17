package api

type ErrorResponse struct {
	Message string `json:"message"`
}

func BuildError(e error) ErrorResponse {
	return ErrorResponse{Message: e.Error()}
}