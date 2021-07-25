package utils

import "strings"

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func GetResourceName(res string) string {
	resTableArr := strings.Split(res, "_")
	var resNameArr []string
	for _, r := range resTableArr {
		resNameArr = append(resNameArr, strings.Title(r))
	}

	return strings.Join(resNameArr, " ")
}
