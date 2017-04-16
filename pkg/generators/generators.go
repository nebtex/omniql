package generators

//AccessorsGenerator should implements all the method necessary for generate the code
//that can handle an omnibuff schema
type AccessorsGenerator interface {
	//Should return a function for access to a boolean
	Boolean(mutationAllowed bool) string
	UnsignedByte(mutationAllowed bool) string
	Short(mutationAllowed bool) string
	UnsignedShort(mutationAllowed bool) string
	Integer(mutationAllowed bool) string
	UnsignedInteger(mutationAllowed bool) string
	Long(mutationAllowed bool) string
	UnsignedLong(mutationAllowed bool) string
	Float(mutationAllowed bool) string
	Double(mutationAllowed bool) string
	String(mutationAllowed bool) string
	Table(mutationAllowed bool) string
	BasicVector(mutationAllowed bool) string
	StringVector(mutationAllowed bool) string
	TableVector(mutationAllowed bool) string
	UnionVector(mutationAllowed bool) string
	ResourceVector(mutationAllowed bool) string
}
// Get the value of a table's scalar.
static void GetScalarFieldOfTable(const StructDef &struct_def,
                                  const FieldDef &field,
                                  std::string *code_ptr) {
  std::string &code = *code_ptr;
  std::string getter = GenGetter(field.value.type);
  GenReceiver(struct_def, code_ptr);
  code += " " + MakeCamel(field.name);
  code += "() " + TypeName(field) + " ";
  code += OffsetPrefix(field) + "\t\treturn " + getter;
  code += "(o + rcv._tab.Pos)\n\t}\n";
  code += "\treturn " + field.value.constant + "\n";
  code += "}\n\n";
}

func (rcv *Monster) Mana() int16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt16(o + rcv._tab.Pos)
	}
	return 150
}



type Monster struct{
	buff = []
}