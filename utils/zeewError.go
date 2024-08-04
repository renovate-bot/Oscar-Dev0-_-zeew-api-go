package ZeewUtils

import "fmt"

const (
	TokenEmpty = "Token is empty"
	EstiloEmpty = "Estilo is empty"
	ErrorMarshallingJson = "Error marshalling JSON"
	ErrorUnmarshallingJson = "Error unmarshalling JSON"
	ErrorRequestingImage = "Error requesting image"
    InvalidImage = "Invalid image"
    RequestError = "Request error"
)



type ZeewError struct {
    Name    string
    Message string
}

func (e *ZeewError) Error() string {
    return fmt.Sprintf("%s: %s", e.Name, e.Message)
}

func NewZeewError(message string) *ZeewError {
    return &ZeewError{
        Name:    "Zeew.API-Error",
        Message: message,
    }
}