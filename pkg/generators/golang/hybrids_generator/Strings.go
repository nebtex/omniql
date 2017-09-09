package hybrids_generator


import (
	"text/template"
	"io"
	"github.com/nebtex/omniql/pkg/io/omniql/corev1"
)

func StringFieldReaderAccessor(freader corev1.FieldReader, wr io.Writer) (fn string, err error) {

	tmpl, err := template.New("StringFieldReaderAccessor").Parse(`
func (t *TableSpec) {{.Field.Name}}() {{ToNativeType .Field.Type}} {
    value, ok := t.table.{{.Field.Type}}()
    if !ok{
		value = {{.Field.Default}}f
    }
}`)
	if err != nil {
		return
	}
	err = tmpl.Execute(wr, nil)
	return
}