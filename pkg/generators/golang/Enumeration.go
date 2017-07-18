package golang

import "github.com/nebtex/omnibuff/pkg/io/omniql/corev1"
import "github.com/nebtex/omnibuff/pkg/generators"

//EnumerationGenerator implements the  function that create the enumeration for golang
//it creates a package for each enumeration, and a file for each group
//because enumeration are a core part of omniql, and expected to have an strict predictable behaviour it don't expose any method for customize the output of the generator
type EnumerationGenerator struct {
}

//GenerateEnumeration  implement enumeration and enumeration group in golang
func (e *EnumerationGenerator) GenerateEnumeration(corev1.ApexMaker, corev1.EnumerationReader, generators.LibraryWriter, corev1.ErrorWriter) error {
	//create a new file under the
	return nil
}

func (e *EnumerationGenerator)(corev1.EnumerationReader, generators.LibraryWriter, corev1.ErrorWriter) error {
	//create a new file under the
	return nil
}


func SetFrament(FragmentBuilder) Executer, error{
	ForWard(FragmentBuilder.NodeEnumerationItem())
	Forward(FragmentBuilder.NodeEnumerationItem().Groups())
}

type QueryPartialStreamer(){

}

type QueryStreamer(){

}

type Executer(){

}



Query.Execute(func(QeuryResult, err:=dxd){

})


func generateGroups(am corev1.ApexMaker, er corev1.EnumerationReader, lw generators.LibraryWriter, ew corev1.ErrorWriter) (err error) {
	apex := am.FromID(er.OqlID())
	groups := apex.EnumerationGroupApex().MatchAny()


	for v:=nil; v==nil{
		v =: er.Items().Next()

	}

	groups.Apply(func(egr corev1.EnumerationGroupReader, applyErr error) (stop bool, funcErr error) {
		stop = true
		if applyErr != nil {
			err = applyErr
			return
		}


		stop = false
		return
	})

	//create a new file
	key, err := generators.StringToKey(gr.Name() + ".go")
	wr, err := lw.NewKey(gr.Name() + ".go")
	if err != nil {
		return
	}

	return nil
}
