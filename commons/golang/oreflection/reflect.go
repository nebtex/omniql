package oreflection

import (
	"sync"
	"github.com/nebtex/hybrids/golang/hybrids"
)

//go:generate mockery -name=ReflectStore
type ReflectStore interface {
	//OReflect return the reflection, of a omniql type using the magic string Type and Items
	//This method should be thread safe for read
	OReflect(magicType string) (o OType, err error)
	//Upsert update the reflection definition
	//This method should be thread safe for write
	Upsert(magicType string, table interface{}) OType

	LookupResourceByID(application string, rid []byte) (ot OType, ok bool)
}

//go:generate mockery -name=LookupFields
type LookupFields interface {
	ByNumber(fn hybrids.FieldNumber) (fieldName string, t OType, ok bool)
	BySnakeCase(fieldName string) (fn hybrids.FieldNumber, t OType, ok bool)
	ByCamelCase(fieldName string) (fn hybrids.FieldNumber, t OType, ok bool)
}

//go:generate mockery -name=LookupEnumeration
type LookupEnumeration interface {
	ByUint8ToCamelCase(input uint8) (value string, ok bool)
	ByUint16ToCamelCase(input uint16) (value string, ok bool)
	ByUint32ToCamelCase(input uint32) (value string, ok bool)
	ByUint8ToSnakeCase(input uint8) (value string, ok bool)
	ByUint16ToSnakeCase(input uint16) (value string, ok bool)
	ByUint32ToSnakeCase(input uint32) (value string, ok bool)
	ByStringToUint8(input string) (value uint8, ok bool)
	ByStringToUint16(input string) (value uint16, ok bool)
	ByStringToUint32(input string) (value uint32, ok bool)
}

//go:generate mockery -name=LookupTableOnUnion
type LookupTableOnUnion interface {
	ByString(input string) (position hybrids.UnionKind, tableType OType, ok bool)
	ByPositionToCamelCase(input hybrids.UnionKind) (value string, tableType OType, ok bool)
	ByPositionToSnakeCase(input hybrids.UnionKind) (value string, tableType OType, ok bool)
}

//go:generate mockery -name=OType
type OType interface {
	Id() string
	Application() string
	Kind() uint16
	Name() string
	Items() OType
	HybridType() hybrids.Types
	FieldCount() hybrids.FieldNumber
	LookupFields() LookupFields
	LookupEnumeration() LookupEnumeration
	LookupTableOnUnion() LookupTableOnUnion
	Enumeration() hybrids.TableReader
	Union() hybrids.TableReader
	Table() hybrids.TableReader
	Resource() hybrids.TableReader
}

type reflectStore struct {
	store  map[string]*OType
	wrLock sync.RWMutex
}

/*
func (r *reflectStore) OReflect(mType string, oItems []byte) (ot *OType, err error) {
	var id []byte
	var ok bool

	if mType == nil {
		err = fmt.Errorf("oreflect: mType should not be nil")
		return
	}

	if oItems == nil {
		id = mType
	} else {
		id = append(mType, byte("["))
		id = append(id, oItems...)
		id = append(id, byte("]"))
	}

	r.wrLock.RLock()
	ot, ok = r.store[id]
	r.wrLock.RUnlock()

	if !ok {
		err = fmt.Errorf("oreflect: Type %s does not found in the store, probably you should update it", string(id))
		return
	}

	return
}

func (r *reflectStore) Upsert(mType string, oItems string, table interface{}) (ot *OType, err error) {
	if mType == "" {
		err = fmt.Errorf("oreflect: mType should not be nil")
		return
	}

	if mType == "Vector" {
		if oItems == "" {
			err = fmt.Errorf("oreflect: oItems should not be nil, when mtype is a vector")
			return
		}

		ht := hybrids.BasicTypeFromString(oItems)
		if ht.isScalar() {
			ot.HybridType = "Vector" + ht.String()
			ot.OmniqlType = "Vector"
			ot.OmniqlItems = ht.String()
			return
		}

	}

	result := bytes.Split(s, []byte("/"))
	if len(result) < 3 || len(result)%2 == 0 {
		return nil
	}
	var idObj *IDReader = nil
	app := result[0]
	for i := 1; i < len(result); i = i + 2 {
		idObj = &OType{id: string(result[i+1]),
			_type:         string(result[i]),
			parent:        idObj,
			application:   string(app),
			isLocal:       isLocal}
	}

	return idObj
}
*/
