package app

import (
	"fmt"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"
)

// Deprecated read from env variables instead of files
type (
	// Params is used to replace placeholders in an error template with the corresponding values.
	Params map[string]interface{}

	errorTemplate struct {
		Message          string `yaml:"message"`
		DeveloperMessage string `yaml:"developer_message"`
	}

	APIError struct {
		// Status represents the HTTP status code
		Status int `json:"-"`
		// ErrorCode is the code uniquely identifying an error
		ErrorCode string `json:"error_code"`
		// Message is the error message that may be displayed to end users
		Message string `json:"message"`
		// DeveloperMessage is the error message that is mainly meant for developers
		DeveloperMessage string `json:"developer_message,omitempty"`
		// Details specifies the additional error information
		Details interface{} `json:"details,omitempty"`
	}
)

var templates map[string]errorTemplate

// Deprecated read from env variables instead of files
func LoadMessages(file string) error {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	templates = map[string]errorTemplate{}
	return yaml.Unmarshal(bytes, &templates)
}

// NewAPIError creates a new APIError with the given HTTP status code, error code, and parameters for replacing placeholders in the error template.
// The param can be nil, indicating there is no need for placeholder replacement.
func NewAPIError(status int, code string, params Params) *APIError {
	err := &APIError{
		Status:    status,
		ErrorCode: code,
		Message:   code,
	}

	if template, ok := templates[code]; ok {
		err.Message = template.getMessage(params)
		err.DeveloperMessage = template.getDeveloperMessage(params)
	}

	return err
}

// getMessage returns the error message by replacing placeholders in the error template with the actual parameters.
func (e errorTemplate) getMessage(params Params) string {
	return replacePlaceholders(e.Message, params)
}

// getDeveloperMessage returns the developer message by replacing placeholders in the error template with the actual parameters.
func (e errorTemplate) getDeveloperMessage(params Params) string {
	return replacePlaceholders(e.DeveloperMessage, params)
}

func replacePlaceholders(message string, params Params) string {
	if len(message) == 0 {
		return ""
	}
	for key, value := range params {
		message = strings.Replace(message, "{"+key+"}", fmt.Sprint(value), -1)
	}
	return message
}
