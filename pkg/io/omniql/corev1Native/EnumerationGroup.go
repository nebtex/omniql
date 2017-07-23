package corev1Native

import (
	"github.com/nebtex/hybrids/golang/hybrids"
	"github.com/nebtex/omnibuff/pkg/io/omniql/corev1"
	"github.com/nebtex/hybrids/golang/hybrids/native"
)

//EnumerationGroup allow to group enumerations
type EnumerationGroup struct {
	Name          string `json:"name"`
	Documentation *Documentation `json:"documentation"`
	Items         []string `json:"items"`
}

//EnumerationGroupReader allow to group enumerations
type EnumerationGroupReader struct {
	_enumerationgroup *EnumerationGroup
	documentation     *DocumentationReader
	items             *native.VectorStringReader
}

//Name ...
func (g *EnumerationGroupReader) Name() (value string) {
	value = g._enumerationgroup.Name
	return
}

//Documentation ...
func (g *EnumerationGroupReader) Documentation() corev1.DocumentationReader {

	if g.documentation != nil {
		return g.documentation
	}

	return nil
}

//Items ...
func (g *EnumerationGroupReader) Items() hybrids.VectorStringReader {
	if g.items != nil {
		return g.items
	}
	return nil
}

func NewEnumerationGroupReader(e *EnumerationGroup) *EnumerationGroupReader {
	if e == nil {
		return nil
	}
	return &EnumerationGroupReader{
		_enumerationgroup: e,
		documentation:     NewDocumentationReader(e.Documentation),
		items:             native.NewVectorStringReader(e.Items),
	}
}

type VectorEnumerationGroupReader struct {
	_vector []*EnumerationGroupReader
}

func (vg *VectorEnumerationGroupReader) Len() (size int) {
	size = len(vg._vector)
	return
}

func (vg *VectorEnumerationGroupReader) Get(i int) (item corev1.EnumerationGroupReader, err error) {

	if i < 0 {
		err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vg._vector)}
		return
	}

	if i > len(vg._vector)-1 {
		err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vg._vector)}
		return
	}

	item = vg._vector[i]
	return

}

func NewVectorEnumerationGroupReader(v []*EnumerationGroup) (obj *VectorEnumerationGroupReader) {
	if v == nil {
		return nil
	}

	obj = &VectorEnumerationGroupReader{}
	obj._vector = make([]*EnumerationGroupReader, len(v))

	for i, value := range v {
		obj._vector[i] = NewEnumerationGroupReader(value)
	}
	return
}
