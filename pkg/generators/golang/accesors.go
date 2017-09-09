package golang

import (
	"github.com/nebtex/omniql/pkg/io/omniql/corev1Native"
	"github.com/nebtex/omniql/pkg/io/omniql/corev1"
	"github.com/nebtex/omniql/pkg/utils"
)

type AccessorsGenerators interface {
	VectorStringAccessor(freader corev1.FieldReader, fn uint16) error
	StringAccessor(freader corev1.FieldReader, fn uint16) error
	VectorTableAccessor(freader corev1.FieldReader, fn uint16, tableName string) error
	TableAccessor(freader corev1.FieldReader, fn uint16, tableName string) error
	EnumerationAccessor(freader corev1.FieldReader, fn uint16, enumName string) error
}

func CreateAccessors(ag AccessorsGenerators, table corev1.TableReader, offset uint16) (err error) {
	var field corev1.FieldReader

	//create Accessors
	for i := 0; i < table.Fields().Len(); i++ {
		field, err = table.Fields().Get(i)
		fieldNumber := uint16(i) + offset
		if err != nil {
			return
		}
		switch field.Type() {
		case "String":
			err = ag.StringAccessor(field, fieldNumber)
			if err != nil {
				return
			}
		case "Vector":
			switch field.Items() {
			case "String":
				err = ag.VectorStringAccessor(field, fieldNumber)
				if err != nil {
					return
				}
			default:
				pid := corev1Native.NewIDReader([]byte(table.Metadata().Application()+"/"+field.Items()), false)
				if pid != nil {
					if pid.Kind() == "Table" {
						err = ag.VectorTableAccessor(field, fieldNumber, utils.TableNameFromID(pid))
						if err != nil {
							return
						}
					}
				}
			}
		default:
			pid := corev1Native.NewIDReader([]byte(table.Metadata().Application()+"/"+field.Type()), false)
			if pid != nil {
				if pid.Kind() == "Table" {
					err = ag.TableAccessor(field, fieldNumber, pid.ID())
					if err != nil {
						return
					}
				}
				if pid.Kind() == "EnumerationGroup" {
					err = ag.EnumerationAccessor(field, fieldNumber, pid.Parent().ID())

					if err != nil {
						return
					}

				}
			}

		}
	}
	return
}
