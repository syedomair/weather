package common

type Config struct {
	DatabaseURL   string `yaml:"database_url"`
	HttpAddress   string `yaml:"http_addr"`
	MapsGoogleURL string `yaml:"maps_google_url"`
	DarkskyURL    string `yaml:"darksky_url"`
	DarkskyKey    string `yaml:"darksky_key"`
}

type ErrorResponseType struct {
	Error_code    string `json:"error_code"`
	Error_message string `json:"error_message"`
}

func ErrorResponse(class interface{}) map[string]interface{} {

	var errType ErrorResponseType
	errType.Error_code = class.(string)[:5]
	errType.Error_message = class.(string)[6:]
	return commonResponse(errType, "error", "500")
}

func SuccessResponse(class interface{}) map[string]interface{} {
	return commonResponse(class, "success", "200")
}

func commonResponse(class interface{}, result string, code string) map[string]interface{} {
	response := make(map[string]interface{})
	response["data"] = class
	response["result"] = result
	response["code"] = code
	return response
}
