package generators

import (
	"github.com/nebtex/omnibuff/pkg/io/omniql/corev1"
	"io"
)

type FieldGenerator interface {
	ReaderMethodInterface(corev1.FieldReader, io.Writer) error
	WriterMethodInterface(corev1.FieldReader, io.Writer) error
	NativeProperty(corev1.FieldReader, io.Writer) error
	ReaderMethodImplementation(corev1.FieldReader, io.Writer) error
	WriterMethodImplementation(corev1.FieldReader, io.Writer) error
}
