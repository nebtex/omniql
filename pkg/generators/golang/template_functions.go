package golang

import (
	"github.com/nebtex/omniql/pkg/utils"
	"strings"
)

var DefaultTemplateFunctions map[string]interface{}

func ShortName(str string) (value string) {
	snake := utils.ToSnakeCase(str)
	items := strings.Split(snake, "_")
	for _, item := range items {
		if item == "" {
			continue
		}
		value += string(item[0])
	}
	return
}

func ScalarToGolangType(str string) (value string) {
	if str == "Integer8"{
		value =  "int8"
	}
	if str == "UnsignedShort"{
		value =  "uint16"
	}
	return
}

func init() {
	DefaultTemplateFunctions = map[string]interface{}{
		"TableName":      utils.TableName,
		"GetPackageName": utils.GolangGetPackageName,
		"ShortName":      ShortName,
		"ToLower":        strings.ToLower,
		"Capitalize":     strings.Title,
		"GoDoc":          GenerateDocs,
		"ToSnakeCase":    utils.ToSnakeCase,
		"ScalarToGolangType": ScalarToGolangType,
	}
}
