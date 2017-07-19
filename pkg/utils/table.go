package utils

import (
	"github.com/nebtex/omnibuff/pkg/io/omniql/corev1"
	"github.com/nebtex/omnibuff/pkg/io/omniql/corev1Native"
)

//TableName return the full table name taken into account the parent
func TableName(t corev1.TableReader) (value string) {
	pid := corev1Native.NewIDReader([]byte(t.Meta().Application()+"/"+t.Meta().Parent()), false)
	if pid != nil {
		return pid.ID() + t.Meta().Name()

	}
	return t.Meta().Name()
}
