package interface_generator

import (
	"go.uber.org/zap"
	"io"
	"text/template"
	"github.com/nebtex/omnibuff/pkg/generators/golang"
	"bytes"
	"github.com/nebtex/omnibuff/pkg/io/omniql/corev1"
	"strings"
	"fmt"
)

type ShardGenerator struct {
	zap            *zap.Logger
	buffer         *bytes.Buffer
	wildcardBuffer *bytes.Buffer
	forwardBuffer  *bytes.Buffer
	table          corev1.TableReader
}

func NewShardGenerator(table corev1.TableReader, logger *zap.Logger) *ShardGenerator {
	zap := logger.With(zap.String("TableName", table.Metadata().Name()),
		zap.String("Type", "Reader Interface"),
		zap.String("Application", table.Metadata().Application()),
	)
	s := &ShardGenerator{table: table, zap: zap}
	s.buffer = bytes.NewBuffer(nil)
	s.wildcardBuffer = bytes.NewBuffer(nil)
	s.forwardBuffer = bytes.NewBuffer(nil)

	return s
}

func (s *ShardGenerator) Start() (err error) {
	tmpl, err := template.New("ShardGenerator::interface::StartStruct").Funcs(golang.DefaultTemplateFunctions).
		Parse(`
//{{TableName .}}Shard ...
type {{TableName .}}Shard interface {
`)
	if err != nil {
		return
	}
	err = tmpl.Execute(s.buffer, s.table)
	if err != nil {
		return
	}

	tmpl, err = template.New("ShardGenerator::interface::StartWildcard").Funcs(golang.DefaultTemplateFunctions).
		Parse(`
//{{TableName .}}WildcardShard ...
func {{TableName .}}WildcardShard(s {{TableName .}}Shard){
`)
	if err != nil {
		return
	}
	err = tmpl.Execute(s.wildcardBuffer, s.table)
	if err != nil {
		return
	}

	tmpl, err = template.New("ShardGenerator::interface::StartForward").Funcs(golang.DefaultTemplateFunctions).
		Parse(`
//{{TableName .}}ForwardShard ...
func {{TableName .}}ForwardShard(s {{TableName .}}Shard){
`)
	if err != nil {
		return
	}
	err = tmpl.Execute(s.forwardBuffer, s.table)

	return
}

func (s *ShardGenerator) StructAddField(fn string, ft string) (err error) {
	_, err = s.buffer.Write([]byte("\n    " + fn + "() " + ft))
	return
}

func (s *ShardGenerator) End() (err error) {
	_, err = s.buffer.Write([]byte("\n" + "}\n"))
	if err != nil {
		return
	}

	_, err = s.wildcardBuffer.Write([]byte("\n" + "}\n"))
	if err != nil {
		return
	}
	_, err = s.forwardBuffer.Write([]byte("\n" + "}\n"))

	return
}

func (s *ShardGenerator) VectorStringAccessor(freader corev1.FieldReader, fn uint16) (err error) {
	err = s.StructAddField(strings.Title(freader.Name()), "ShardHolderAndDisposer")
	if err != nil {
		return
	}
	_, err = s.wildcardBuffer.Write([]byte(fmt.Sprintf("\n    s.%s()", strings.Title(freader.Name()))))
	if err != nil {
		return
	}
	_, err = s.forwardBuffer.Write([]byte(fmt.Sprintf("\n    s.%s()", strings.Title(freader.Name()))))

	return
}

func (s *ShardGenerator) StringAccessor(freader corev1.FieldReader, fn uint16) (err error) {
	err = s.StructAddField(strings.Title(freader.Name()), "ShardHolderAndDisposer")
	if err != nil {
		return
	}
	_, err = s.wildcardBuffer.Write([]byte(fmt.Sprintf("\n    s.%s()", strings.Title(freader.Name()))))
	if err != nil {
		return
	}
	_, err = s.forwardBuffer.Write([]byte(fmt.Sprintf("\n    s.%s()", strings.Title(freader.Name()))))

	return
}
func (s *ShardGenerator) VectorTableAccessor(freader corev1.FieldReader, fn uint16, tableName string) (err error) {
	_, err = s.buffer.Write([]byte(fmt.Sprintf("\n    %s(func(%sShard)) ShardHolderAndDisposer", strings.Title(freader.Name()), tableName)))
	if err != nil {
		return
	}
	_, err = s.forwardBuffer.Write([]byte(fmt.Sprintf("\n    s.%s(%sForwardShard)", strings.Title(freader.Name()), tableName)))
	return
}
func (s *ShardGenerator) TableAccessor(freader corev1.FieldReader, fn uint16, tableName string) (err error) {
	_, err = s.buffer.Write([]byte(fmt.Sprintf("\n    %s(func(%sShard)) ShardHolderAndDisposer", strings.Title(freader.Name()), tableName)))
	if err != nil {
		return
	}
	_, err = s.forwardBuffer.Write([]byte(fmt.Sprintf("\n    s.%s(%sForwardShard)", strings.Title(freader.Name()), tableName)))

	return
}

func (s *ShardGenerator) EnumerationAccessor(freader corev1.FieldReader, fn uint16, enumName string) (err error) {
	err = s.StructAddField(strings.Title(freader.Name()), "ShardHolderAndDisposer")
	if err != nil {
		return
	}
	_, err = s.wildcardBuffer.Write([]byte(fmt.Sprintf("\n    s.%s()", strings.Title(freader.Name()))))
	if err != nil {
		return
	}
	_, err = s.forwardBuffer.Write([]byte(fmt.Sprintf("\n    s.%s()", strings.Title(freader.Name()))))

	return
}

func (s *ShardGenerator) Generate(wr io.Writer) (err error) {
	//generate struct
	err = s.Start()
	if err != nil {
		return
	}
	err = golang.CreateAccessors(s, s.table, 0)
	if err != nil {
		return
	}
	err = s.End()
	if err != nil {
		return
	}

	_, err = s.buffer.WriteTo(wr)

	if err != nil {
		return
	}

	//wildcard function
	_, err = s.wildcardBuffer.WriteTo(wr)

	if err != nil {
		return
	}
	//forward function
	_, err = s.forwardBuffer.WriteTo(wr)

	return
}
