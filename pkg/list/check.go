package list

func CheckStringInList(v string, list []string) bool {
	for _, l := range list {
		if l == v {
			return true
		}
	}

	return false
}
