package main

// merge multiple results into single map
func mergeResults(maps []map[string]interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}
