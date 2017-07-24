package golang

import (
	"strings"
	"io"
	"fmt"
)

type Imports struct {
	list []string
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func (i *Imports) AddImport(str string) {
	if i.DoImportExist(str) {
		return
	}
	i.list = append(i.list, strings.TrimSpace(str))

}

func (i *Imports) DoImportExist(str string) bool {
	return stringInSlice(strings.TrimSpace(str), i.list)

}

func (i *Imports) Write(writer io.Writer) (err error) {
	if len(i.list) == 0 {
		return
	}

	value := "import("
	for _, item := range i.list {
		value += fmt.Sprintf("\n    \"%s\"", item)
	}
	value += "\n)"
	_, err = writer.Write([]byte(value))
	return
}

func NewImports() *Imports {
	return &Imports{}
}
