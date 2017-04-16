package generators

//GoGenerator generate the code for read omnibuff schemas in golang
type GoGenerator struct {
}

//Boolean implements AccessorsGenerator.Boolean
func (gc *GoGenerator) Boolean(mutationAllowed bool) string {
	return `
	`

}

type Tree struct {
	offset   int32
	byteStart int32
	byEnds    int32
	kind int32
	Children []*Tree
}

type Monster struct{
	_tableBuff []byte
	_nameBuff  []byte

}

fromEdge(Tree){
	this._tableBuff = rootBuff[this.offset]
	name = this.NamePos()
	_nameBuff = rootBuff[]
}

Monster.Name(
	string(this.NameBuff)
	return rootBuff[tree.byteStart: tree.byteEnd]
)
Monster.SetNmae()


all.sssss=function(){
	return 
}