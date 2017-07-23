package corev1Native

import (
	"bytes"
	"github.com/nebtex/omnibuff/pkg/io/omniql/corev1"
)

type IDReader struct {
	id          string
	application string
	_type       string
	parent      *IDReader
	isLocal     bool
}

func (i *IDReader) ID() string {
	return i.id
}
func (i *IDReader) Application() string {
	return i.application
}

func (i *IDReader) IsLocal() bool {
	return i.isLocal
}

func (i *IDReader) Parent() corev1.ResourceIDReader {
	if i.parent!=nil{
		return i.parent
	}
	return nil
}

func (i *IDReader) Kind() string {
	return i._type
}

func NewIDReader(ID []byte, isLocal bool) *IDReader {
	result := bytes.Split(ID, []byte("/"))
	if len(result) < 3 || len(result)%2 == 0 {
		return nil
	}
	var idObj *IDReader = nil
	app := result[0]
	for i := 1; i < len(result); i=i+2 {
		idObj = &IDReader{id: string(result[i+1]),
			_type:            string(result[i]),
			parent:           idObj,
			application:      string(app),
			isLocal:          isLocal}
	}

	return idObj
}
