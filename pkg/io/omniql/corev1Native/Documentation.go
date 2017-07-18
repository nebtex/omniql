package corev1Native

import "github.com/nebtex/omnibuff/pkg/io/omniql/corev1"

type Documentation struct {
	Short string `json:"short"`
	Long  string `json:"long"`
}


type DocumentationReader struct {
	documentation *Documentation
}


func (d *DocumentationReader) Short() (value string) {
	value = d.documentation.Short
	return
}

func (d *DocumentationReader) Long() (value string) {
	value = d.documentation.Long
	return
}

func NewDocumentationReader(d *Documentation) corev1.DocumentationReader {
	if d == nil {
		return nil
	}
	return &DocumentationReader{d}
}

func NewDeepDocumentationReader(d *Documentation) corev1.DocumentationReader {
	if d == nil {
		return nil
	}
	return &DocumentationReader{d}
}
