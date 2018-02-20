//Package strtemplate designed for get string with data from string template
//v.0.2.20180221.2
package strtemplate

import (
	"errors"
	"regexp"
	"strings"
)

//GetString function for get string from template
func GetString(data map[string]interface{}, template string) (*string, error) {

	//getStr(dataelement)

	outmsg := template

	for key, val := range data {
		reg := regexp.MustCompile("{{" + key + "}}")
		if reg.FindStringIndex(outmsg) == nil {
			return nil, errors.New(errDataNotFound)
		}
		outmsg = reg.ReplaceAllString(outmsg, getStr(val))
	}

	if strings.Count(outmsg, "{{") > 0 {
		return nil, errors.New(errElementNotFound)
	}

	return &outmsg, nil
}
