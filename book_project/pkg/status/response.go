package res

import (
	"net/http"
)

//struct returned to handler
type HTTPRES struct {
	Status  int         `json:"status"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//checking errors with staus passed by handler func
func CheckERRor(status int, data interface{}) HTTPRES {
	switch status {
	case http.StatusBadRequest,
		http.StatusInternalServerError,
		http.StatusRequestTimeout,
		http.StatusUnauthorized,
		http.StatusForbidden:
		if e, ok := data.(error); ok { //data contains err
			return HTTPRES{
				Status:  status,
				Success: false,
				Message: e.Error(),
			}
		}
		return HTTPRES{ //data contains string
			Status:  status,
			Success: false,
			Message: data.(string),
		}
	default:
		return HTTPRES{
			Status:  status,
			Success: true,
			Data:    data,
		}

	}
}
