package generators

import (
	"github.com/nebtex/omnibuff/pkg/io/omniql/corev1"
	"io"
	"text/template"
	"strings"
	"github.com/nebtex/omnibuff/pkg/io/omniql/corev1BasicTypes"
	"fmt"
)

type FieldTemplateFunctions interface {
	FieldName(corev1.FieldReader) (value string, err error)
	FieldType(corev1.FieldReader) (value string, err error)
	MutatePrefix() string
}

type FieldGenerator struct {
	readerMethodTemplate *template.Template
	writerMethodTemplate *template.Template
}

func NewFieldGenerator(fm FieldTemplateFunctions) (fg *FieldGenerator, err error) {
	var funcMap template.FuncMap
	if fm == nil {
		funcMap = template.FuncMap{
			"FieldName": func(field corev1.FieldReader) (value string, err error) {
				value = strings.Title(field.Name())
				return
			},
			"FieldType": func(field corev1.FieldReader) (value string, err error) {
				value = strings.Title(field.Name())
				return
			},
		}
	} else {
		funcMap = template.FuncMap{
			"FieldName": fm.FieldName,
			"FieldType": fm.FieldType,
		}
	}


	// create templates
	wmtTmpl, err := template.New("WriterMethodInterface").Funcs(funcMap).Parse(`{{.MutatePrefix}}{{.FieldName .}}({{.FieldType .}}) err`)
	if err != nil {
		return
	}


	fg = &FieldGenerator{}
	fg.readerMethodTemplate = rmtTmpl
	return fg, nil
}

func (fg *FieldGenerator) ReaderMethodInterface(fr corev1.FieldReader, wr io.Writer) (err error) {
	types := strings.Split(fr.Type(), "/")
	if len(types)==1{
		//Is a basic type
		fieldType := corev1BasicTypes.None.FromString(types[0])
		if fieldType == corev1BasicTypes.None{
			//Todo: better errors
			err = fmt.Errorf("Invalid BasicType %s", fieldType.String())
			return
		}
		if fieldType.IsScalar(){
			//TODO: template map
			// create templates
			rmtTmpl, err := template.New("ReaderMethodInterface").Funcs(fg.funcMap).Parse(`{{.FieldName .}}() {{.FieldType .}}`)
			if err != nil {
				return
			}
			err = fg.readerMethodTemplate.Execute(wr, fr)
			return
		}
		if fieldType == corev1BasicTypes.Vector{
			rmtTmpl, err := template.New("ReaderMethodInterface").Funcs(fg.funcMap).Parse(`{{.FieldName .}}() {{.FieldType .}}`)

		}
		if fieldType == corev1BasicTypes.String{

			rmtTmpl, err := template.New("ReaderMethodInterface").Funcs(fg.funcMap).Parse(`{{.FieldName .}}() {{.FieldType .}}`)


		}
	}



	return
}
func (fg *FieldGenerator) WriterMethodInterface(fr corev1.FieldReader, wr io.Writer) (err error) {
	err = fg.writerMethodTemplate.Execute(wr, fr)
	return
}
