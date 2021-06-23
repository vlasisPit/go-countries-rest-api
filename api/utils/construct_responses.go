package utils

import "net/http"

func ConstructErrorResponse(writer http.ResponseWriter, errorMessage string, serverError int) {
	writer.WriteHeader(serverError)
	writer.Write([]byte(errorMessage))
}
