package param

// HttpError example
type BadRequestHttpError struct {
	Error   string `json:"error" example:"code=400, message=bad request"`
	Message string `json:"message" example:"status bad request"`
}

type UnprocessableEntity struct {
	Error   string `json:"error" example:"code=422, message=Unprocessable Entity"`
	Message string `json:"message" example:"request parameters are not valid"`
}

type NotFound struct {
	Error   string `json:"error" example:"code=404, message=not found"`
	Message string `json:"message" example:"requested object not found"`
}

type InternalError struct {
	Error   string `json:"error" example:"code=500, message=internal server error"`
	Message string `json:"message" example:"nil pointer"`
}

type UnAuthorizedHttpError struct {
	Error   string `json:"error" example:"code=401, message=invalid or expired jwt"`
	Message string `json:"message" example:"unauthorized"`
}
