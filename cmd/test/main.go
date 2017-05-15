package main

import (
	"github.com/nebtex/omnibuff/pkg/schema"
	"fmt"
	"bytes"
	"github.com/google/flatbuffers/go"
)

type MetaReader interface {
	Name() string
	Application() string
}
type MetaWriter interface {
	Name() string
	Application() string
	SetApplication(s string)
	SetName(s string) string
}

type ResourceReader interface {
	Meta() MetaReader
	Fields() string
}

type ResourceWriter interface {
	Meta() MetaWriter
	SetMeta(m MetaWriter)
	Fields() string
	SetFields(s string)

}

func beta(r MetaWriter){

}
func lala(r ResourceWriter){
	beta(r.Meta())
	fmt.Println(r)
}

type ResourceAsMutable struct {
	l map[ResourceWriter]MetaWriter


}

type NodeValue struct {
	Children map[int64]NodeType
	offset flatbuffers.UOffsetT
}

type NodeType struct {
	Children map[int64]NodeValue

}

func main() {
	app, err := schema.Load("/home/cristian/nebtex/go/src/github.com/nebtex/omnibuff/reflection/omniql")
	fmt.Println(err)
	buf := bytes.NewBufferString("namespace io.omniql.core.v1.fbs;\nunion UnionItems { Interface, Table, Resource }\n")


	//make tables
	for _, v := range app.Resources {
		err := v.ToStreamInterface(buf)
		if err!=nil{
			panic(err)
		}
	}


	fmt.Println(buf.String())
	//make unions

	//make resource

	//make root

}
