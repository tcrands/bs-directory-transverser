package main

import "regexp"

func getParams(regEx, val string) (paramsMap map[string][]string) {

	compRegEx := regexp.MustCompile(regEx)
	match := compRegEx.FindAllStringSubmatch(val, -1)

	paramsMap = make(map[string][]string)
	for i, name := range compRegEx.SubexpNames() {
		if i > 0 && i <= len(match) {
			for j := 0; j <= len(match)-1; j++ {
				paramsMap[name] = append(paramsMap[name], match[j][i])
			}
		}
	}
	return paramsMap
}

func getSplitPoint(regEx, val string) string {
	compRegEx := regexp.MustCompile(regEx)
	match := compRegEx.FindAllString(val, -1)

	return match[len(match)-1]
}
