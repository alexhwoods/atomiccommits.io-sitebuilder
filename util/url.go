package util

import (
	"strings"
)

func reverse(arr []string) []string {
	newArr := make([]string, len(arr))
	for i, j := 0, len(arr)-1; i <= j; i, j = i+1, j-1 {
		newArr[i], newArr[j] = arr[j], arr[i]
	}
	return newArr
}

func InvertUrl(url string) string {
	newUrl := strings.ReplaceAll(url, "https://", "")
	newUrl = strings.ReplaceAll(newUrl, "http://", "")

	result := strings.SplitN(newUrl, "/", 2)
	isRootDomain := len(result) == 1
	path := ""
	if !isRootDomain {
		path = result[1]
	}

	reversedDomain := strings.Join(reverse(strings.SplitN(result[0], ".", -1)), ".")

	if isRootDomain {
		return reversedDomain
	}

	return reversedDomain + "/" + path
}
