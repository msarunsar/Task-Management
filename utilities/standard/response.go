package standard

type StandardReponse struct {
	Code         int
	Message      string
	ErrorMessage *string `json:"ErrorMessage,omitempty"`
}

func BadRequest(err string) StandardReponse {
	return StandardReponse{
		Code:         400,
		Message:      "Bad Request",
		ErrorMessage: &err,
	}
}

func NotFound(err string) StandardReponse {
	return StandardReponse{
		Code:         404,
		Message:      "Not Found",
		ErrorMessage: &err,
	}
}

func CreatedOrUpdatedStatus(msg string) StandardReponse {
	return StandardReponse{
		Code:    201,
		Message: msg,
	}
}

func OkStatus() StandardReponse {
	return StandardReponse{
		Code:    200,
		Message: "OK",
	}
}
