package aci

import (
	"fmt"
	"log"
	"regexp"
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

func filterChildrenDn(dn, parentDns string) string {
	parentDnList := strings.Split(parentDns, ":")

	for _, parentDn := range parentDnList {
		if strings.HasPrefix(dn, parentDn) {
			return dn
		}
	}

	return ""
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

func GetParentDn(dn string, rn string) string {
	arr := strings.Split(dn, rn)
	return arr[0]
}

func replaceSpecialCharsDn(dn string) string {
	hyphen := regexp.MustCompile(`[/.:]`)
	removeChars := regexp.MustCompile(`[\[\]]`)
	res := hyphen.ReplaceAllString(dn, "-")
	res = removeChars.ReplaceAllString(res, "")
	return res
}

func resourceNamefromDn(className, dn string, i int) string {
	return fmt.Sprintf("%s_%s_%d", className, replaceSpecialCharsDn(GetMOName(dn)), i)
}

var resourceRequiredFieldMap = map[string][]string{
	"aci_epg_to_contract": {"tDn"},
}

func checkTarget(rName string, container *container.Container) bool {
	return checkRequiredFields(rName, container)
}

func checkRequiredFields(rName string, container *container.Container) bool {
	attrs := resourceRequiredFieldMap[rName]
	for _, attr := range attrs {
		res := G(container, attr)
		if res == "" || res == "{}" {
			return false
		}
	}
	return true

}

func printMissTarget(missTargetResources []string) {
	resources := missTargetResources
	if len(resources) > 0 {
		log.Println("[WARNING]\n The following resources have not being imported due to 'miss-target'")
		for _, resource := range resources {
			fmt.Printf("   - %s\n", resource)
		}
	}
}
