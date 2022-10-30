package utils

func GetSessionValue(s interface{}) map[string]interface{} {
	values, _ := s.(map[string]interface{})
	result := make(map[string]interface{})

	for k, v := range values {
		result[k] = v
	}
	return result
}
