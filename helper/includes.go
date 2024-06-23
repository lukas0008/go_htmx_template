package helper

func includes[K comparable](arr []K, value K) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}