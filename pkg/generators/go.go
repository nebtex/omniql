package generators

//GoGenerator generate the code for read omnibuff schemas in golang
type GoGenerator struct {

}

//Scalar implements AccessorsGenerator.Boolean
func (gc *GoGenerator) Scalar(mutationAllowed bool) string {
	return `func ({instance} *{object}){field.Name}(){GoType}{
		return {instance}._fbo.{field.Name}()
	}
}`}

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


struct Table {
	scalars = []byte
	nonScalars [][]byte
	children []interface{}
}

func (t *Table)Name(){
	return (childs[1])
}

func (t Table)AAA(){
	tables[1]
}