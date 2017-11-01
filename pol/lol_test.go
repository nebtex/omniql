package main

import (
	"testing"
	"github.com/nebtex/delta/go"
	"sort"
	"github.com/Pallinder/go-randomdata"
	"fmt"
)

var values []string
var mapp map[string]bool

const max = delta.MaxUint16/16

func init() {
	values = make([]string, 0, max)
	mapp = map[string]bool{}

	for i := 0; i < max; i++ {
		s := randomdata.SillyName()
		for{
			_, ok := mapp[s]
			if !ok{
				break
			}
			s = randomdata.SillyName()

		}

		values = append(values, s)
		mapp[s] = true
	}
	fmt.Println(len(mapp))
	sort.Strings(values)

}
func BenchmarkSample(b *testing.B) {

	for i := 0; i < b.N; i++ {
		//_ = mapp[values[i%max]]
		BinarySearch(values, values[i%max])
	}

	//	fmt.Println(BinarySearch(values, 5))

}
