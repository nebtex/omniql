package generators

import (
	"github.com/nebtex/omnibuff/pkg/io/omniql/corev1"
	"io"
)

type Field interface {
	ReaderMethodInterface(corev1.FieldReader, io.Writer, corev1.ErrorWriter) error
	WriterMethodInterface(corev1.FieldReader, io.Writer, corev1.ErrorWriter) error
	NativeProperty(corev1.FieldReader, io.Writer, corev1.ErrorWriter) error
	ReaderMethodImplementation(corev1.FieldReader, io.Writer, corev1.ErrorWriter) error
	WriterMethodImplementation(corev1.FieldReader, io.Writer, corev1.ErrorWriter) error
}


