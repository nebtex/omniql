package main

import (
	"github.com/nebtex/omnibuff/pkg/schema"
	"fmt"
)

func main() {
	app, err := schema.LoadResource("/home/cristian/nebtex/go/src/github.com/nebtex/omnibuff/examples/menshend/AdminService")
	fmt.Println(err)
	fbs := ""
	//make tables
	for _, v := range app.Tables {
		res, err := v.ToFlatBuffer()
		if err!=nil{
			panic(err)
		}
		fbs += res
		fmt.Println(fbs)

	}
	//make unions

	//make resource

	//make root

}
