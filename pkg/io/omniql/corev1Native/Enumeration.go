package corev1Native

import (
	"github.com/nebtex/omnibuff/pkg/io/omniql/corev1"
)

//Enumeration ...
type Enumeration struct {
	RID      []byte `json:"rid"`
	Metadata *Metadata `json:"meta"`
	Kind     string `json:"kind"`
	Items    []*EnumerationItem `json:"items"`
	Groups   []*EnumerationGroup `json:"groups"`
}

//EnumerationReader ...
type EnumerationReader struct {
	_enumeration *Enumeration
	_rid         *IDReader
	metadata     *MetadataReader
	items        *VectorEnumerationItemReader
	groups       *VectorEnumerationGroupReader
}

//RID get resource id
func (e *EnumerationReader) RID() corev1.ResourceIDReader {
	return e._rid
}

//Metadata ...
func (e *EnumerationReader) Metadata() corev1.MetadataReader {

	if e.metadata != nil {
		return e.metadata
	}

	return nil
}
func (e *EnumerationReader) Kind() string {
	return e._enumeration.Kind
}

//Items ...
func (e *EnumerationReader) Items() corev1.VectorEnumerationItemReader {

	if e.items != nil {
		return e.items
	}

	return nil
}

//Groups ...
func (e *EnumerationReader) Groups() corev1.VectorEnumerationGroupReader {

	if e.groups != nil {
		return e.groups
	}

	return nil
}

func NewEnumerationReader(e *Enumeration) *EnumerationReader {
	if e == nil {
		return nil
	}
	return &EnumerationReader{
		_enumeration: e,
		_rid:         NewIDReader(e.RID, false),
		metadata:     NewMetadataReader(e.Metadata),
		items:        NewVectorEnumerationItemReader(e.Items),
		groups:       NewVectorEnumerationGroupReader(e.Groups),

	}
}
