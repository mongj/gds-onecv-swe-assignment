package api

type GenericResponse struct {
	Message string `json:"message"`
}

func BuildError(e error) GenericResponse {
	return GenericResponse{Message: e.Error()}
}
