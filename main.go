//Package strtemplate designed for get string with data from string template
//v.0.2.20180221
package strtemplate

import (
	"regexp"
)

//GetString function for get string from template
func GetString(data map[string]interface{}, template string) (*string, error) {

	//getStr(dataelement)

	outmsg := template

	for key, val := range data {
		reg := regexp.MustCompile("{{" + key + "}}")
		outmsg = reg.ReplaceAllString(outmsg, getStr(val))
	}

	return &outmsg, nil
}
