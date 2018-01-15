//Package strtemplate designed for get string with data from string template
//v.0.1.20180116
package strtemplate

import (
	"errors"
	"strings"
)

//GetString function for get string from template
func GetString(data map[string]interface{}, template string) (*string, error) {

	var outmsg string
	nreplacement := strings.Count(template, "{{")

	if nreplacement > 0 {
		for i := 0; i < nreplacement; i++ {
			beginkey := strings.Index(template, "{{")
			if beginkey == -1 {
				outmsg += template
				return &outmsg, nil
			}

			endkey := strings.Index(template, "}}")
			if beginkey < 0 || endkey < 3 {
				return &outmsg, errors.New(errTemplateFormat)
			}

			key := template[beginkey+2 : endkey]

			if dataelement, ok := data[key]; ok {
				outmsg = template[:beginkey] + getStr(dataelement)

				if endkey+2 < len(template) {
					template = template[endkey+2:]
				} else {
					return &outmsg, nil
				}
			} else {
				return &outmsg, errors.New(errElementNotFound)
			}
		}
		return &outmsg, nil
	}
	return &outmsg, errors.New(errTemplateDescription)
}
