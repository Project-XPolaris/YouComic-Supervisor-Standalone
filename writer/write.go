package writer

import (
	"fmt"
	"regexp"
	"strings"
)

func WriteJSStringProperty(configContent []byte, propertyName string, newValue string) []byte {
	configString := string(configContent)
	apiURLRe, err := regexp.Compile(fmt.Sprintf("%s:\\s{0,}\".*?\"", propertyName))
	if err != nil { // Handle errors reading the config file
		panic(err)
	}
	apiURLConfigString := string(apiURLRe.Find(configContent))
	configString = strings.Replace(configString, apiURLConfigString, fmt.Sprintf("%s: \"%s\"", propertyName, newValue), 1)
	return []byte(configString)
}
