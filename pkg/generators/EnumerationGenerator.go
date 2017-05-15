package generators

import (
	"github.com/nebtex/omnibuff/pkg/io/omniql/corev1"
	"io"
	"strings"
	"fmt"
)

type libraryWriterKey string

func StringToKey(str string) (lwk libraryWriterKey, err error) {
	if strings.Contains(str, "/") {
		err = fmt.Errorf("%s", "Invalid Key: / slash are not allowed")
	}
	lwk = libraryWriterKey(str)
	return
}

//LibraryWriter
type LibraryWriter interface {
	NewKey(key libraryWriterKey) (io.WriteCloser, error)
	NewFolder(key libraryWriterKey) (LibraryWriter, error)
}

//EnumerationGenerator
//Implementations must not retain any of the parameters
type EnumerationGenerator interface {
	GenerateEnumeration(er corev1.EnumerationReader, lw LibraryWriter, ew corev1.ErrorWriter) error
}
