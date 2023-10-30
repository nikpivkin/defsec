package iam

func IsWildcardAllowed(actions ...string) ([]string, bool) {
	var result []string
	for _, action := range actions {
		if _, ok := allowedActionsForResourceWildcardsMap[action]; !ok {
			result = append(result, action)
		}
	}
	return result, len(result) == 0
}
