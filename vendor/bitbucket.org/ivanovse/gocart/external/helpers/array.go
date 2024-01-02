package helpers

func MakeMapEmptyValue(values []string) map[string]struct{} {
	result := make(map[string]struct{}, len(values)) 
	for _, v := range values {
		result[v] = struct{}{}
	}

	return result
}
