// The following directive is necessary to make the package coherent:

// +build ignore

// This program generates scalar-decoder.gen.go It can be invoked by running
// go generate

package main

import "os"
import "text/template"
import "log"

import (
	"io"
	"github.com/nebtex/hybrids/golang/hybrids"
)

func RenderScalarFakers(f io.Writer) {
	var err error
	var tmp *template.Template

	tmp, err = template.New("ScalarFakers").Parse(`
func (j *Json) get{{.scalar.String}}() (r {{.jsonType}}, err error) {
	var v int

	v, err = randutil.IntRange({{.min}}, {{.max}})

	if err != nil {
		return
	}

	r = {{.jsTransform}}

	return
}

func (j *Json) fake{{.scalar.String}}(path string, out map[string]interface{}, fieldName string) (err error) {
	var v {{.jsonType}}

	v, err = j.get{{.scalar.String}}()

	if err != nil {
		return
	}

	out[fieldName] = v

	return
}
	
func (j *Json) fakeVector{{.scalar.String}}(path string, out map[string]interface{}, fieldName string) (err error) {
	var choice randutil.Choice
	var vLen int
	var v {{.jsonType}}

	// generate nil or vector
	nullVector := randutil.Choice{10, int(0)}
	newVector := randutil.Choice{100, int(1)}

	choice, err = randutil.WeightedChoice([]randutil.Choice{nullVector, newVector})

	if err != nil {
		return err
	}

	item, _ := choice.Item.(int)

	if item == 0 {
		out[fieldName] = nil
		return
	}

	//generate vector len
	vLen, err = randutil.IntRange(1, 127)

	if err != nil {
		return err
	}

	r := make([]interface{}, 0, vLen)

	for i := 0; i < vLen; i++ {
		v, err =  j.get{{.scalar.String}}()
		if err != nil {
	        return err
	    }
		r[i] = v
	}

	out[fieldName] = r

	return
}

`)
	die(err)

	scalars := []hybrids.Types{
		hybrids.Boolean,
		hybrids.Int8,
		hybrids.Uint8,
		hybrids.Int16,
		hybrids.Uint16,
		hybrids.Int32,
		hybrids.Uint32,
		hybrids.Int64,
		hybrids.Uint64,
		hybrids.Float32,
		hybrids.Float64,
	}
	jsonFunction := []string{
		"iToBool(v)",
		"float64(v)",
		"float64(v)",
		"float64(v)",
		"float64(v)",
		"float64(v)",
		"float64(v)",
		"strconv.FormatInt(int64(v), 10)",
		"strconv.FormatUint(uint64(v), 10)",
		"float64(v)",
		"float64(v)",
	}

	jsonType := []string{
		"bool",
		"float64",
		"float64",
		"float64",
		"float64",
		"float64",
		"float64",
		"string",
		"string",
		"float64",
		"float64",
	}

	min := []string{
		"hybrids.MinBoolean",
		"hybrids.MinInt8",
		"hybrids.MinUint8",
		"hybrids.MinInt16",
		"hybrids.MinUint16",
		"hybrids.MinInt32",
		"hybrids.MinUint32",
		"-200000000",
		"0",
		"-158",
		"-3500",
	}

	max := []string{
		"hybrids.MaxBoolean",
		"hybrids.MaxInt8",
		"hybrids.MaxUint8",
		"hybrids.MaxInt16",
		"hybrids.MaxUint16",
		"hybrids.MaxInt32",
		"hybrids.MaxUint32",
		"-200000000",
		"0",
		"-158",
		"-3500",
	}

	for index, v := range scalars {
		err = tmp.Execute(f, map[string]interface{}{
			"scalar":      v,
			"jsTransform": jsonFunction[index],
			"jsonType":    jsonType[index],
			"max":         max[index],
			"min":         min[index],
		})
		die(err)
	}

	decodeScalarTemplate, err := template.New("decodeScalarTemplate").Parse(`

func (j *Json) fakeScalar(path string, out map[string]interface{}, fieldName string, fieldType hybrids.Types)(err error){
	switch fieldType{
{{ range $index, $value := .scalars }}
	case hybrids.{{$value.String}}:
		err = j.fake{{$value.String}}(path, out, fieldName)

{{ end }}
    }
	return
}


func (j *Json) fakeVectorScalar(path string, out map[string]interface{}, fieldName string, fieldType hybrids.Types)(err error){
	switch fieldType{
{{ range $index, $value := .scalars }}
	case hybrids.Vector{{$value.String}}:
		err = j.fakeVector{{$value.String}}(path, out, fieldName)

{{ end }}
	}
	return
}


`)
	die(err)

	err = decodeScalarTemplate.Execute(f, map[string]interface{}{
		"scalars": scalars,
	})
	die(err)

}

func main() {
	f, err := os.Create("scalar-encoder.gen.go")
	die(err)
	defer f.Close()
	_, err = f.Write([]byte("// Code generated, DO NOT EDIT.\n"))
	die(err)
	_, err = f.Write([]byte(`package faker
`))
	die(err)
	_, err = f.Write([]byte(`
import (
	"github.com/nebtex/hybrids/golang/hybrids"
	"strconv"
	"github.com/jmcvetta/randutil"
)
`))
	die(err)
	RenderScalarFakers(f)
	/*RenderVectorScalarDecoders(f)

	test, err := os.Create("scalar-decoder_test.go")
	die(err)
	defer test.Close()
	_, err = test.Write([]byte("// Code generated. DO NOT EDIT.\n"))
	die(err)
	_, err = test.Write([]byte(`package omarshaller
`))
	die(err)
	_, err = test.Write([]byte(`
import (
	"testing"
	"github.com/nebtex/hybrids/golang/hybrids/mocks"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/nebtex/hybrids/golang/hybrids"
	"fmt"
)
`))
	die(err)
	RenderTestScalarDecoders(test)
	RenderVectorScalar_TEST_Decoders(test)*/
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
