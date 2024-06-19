package utils

func StringArrayDiff(src []string, rm []string) []string {

	removed := []string{}
	for _, t := range src {
		found := false
		for _, r := range rm {
			if t == r {
				found = true
			}
		}
		if found {
			continue
		}
		removed = append(removed, t)
	}
	return removed

}
