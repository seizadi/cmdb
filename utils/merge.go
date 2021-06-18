package utils

func MergeInterfaceToStringMaps(b map[interface{}]interface{}) map[string]interface{} {
	out := make(map[string]interface{}, len(b))

	for k, v := range b {
		if ks, ok := k.(string); ok {
			if v, ok := v.(map[interface{}]interface{}); ok {
						out[ks] = MergeInterfaceToStringMaps(v)
						continue
			}
			out[ks] = v
		}
	}
	return out
}

func MergeMaps(a, b map[string]interface{}) map[string]interface{} {
	out := make(map[string]interface{}, len(a))
	for k, v := range a {
		out[k] = v
	}
	for k, v := range b {
		if v, ok := v.(map[string]interface{}); ok {
			if bv, ok := out[k]; ok {
				if bv, ok := bv.(map[string]interface{}); ok {
					out[k] = MergeMaps(bv, v)
					continue
				}
			}
		}
		out[k] = v
	}
	return out
}
