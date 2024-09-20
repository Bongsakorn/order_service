package utils

import "github.com/sym01/htmlsanitizer"

func Sanitized(input map[string]interface{}) map[string]interface{} {
	output := make(map[string]interface{})
	for key, value := range input {

		switch v := value.(type) {
		case map[string]interface{}:
			value = Sanitized(v)
			output[key] = value

		case []interface{}:
			value = SanitizedList(v)
			output[key] = value

		case string:
			sanitized, _ := htmlsanitizer.SanitizeString(v)
			output[key] = sanitized

		default:
			output[key] = v
		}
	}
	return output
}

func SanitizedList(input []interface{}) []interface{} {
	nestedOutput := make([]interface{}, len(input))
	for index, value := range input {
		switch v := value.(type) {
		case []interface{}:
			SanitizedList(v)
		case map[string]interface{}:
			nestedOutput[index] = Sanitized(v)
		default:
			nestedOutput[index] = v
		}
	}
	return nestedOutput
}
