package aci

import (
	"strconv"
	"strings"

	container "github.com/Jeffail/gabs"
)

func GetMOName(dn string) string {
	arr := strings.Split(dn, "/")
	// Get the last element
	last_ele := arr[len(arr)-1]
	// split on -
	dash_split := strings.Split(last_ele, "-")
	// join except first element as that will be rn
	return strings.Join(dash_split[1:], "-")
}

func StrtoInt(s string, startIndex int, bitSize int) (int64, error) {
	return strconv.ParseInt(s, startIndex, bitSize)
}

func stripQuotes(word string) string {
	if strings.HasPrefix(word, "\"") && strings.HasSuffix(word, "\"") {
		return strings.TrimSuffix(strings.TrimPrefix(word, "\""), "\"")
	}
	return word
}

func G(cont *container.Container, key string) string {
	return stripQuotes(cont.S(key).String())
}
