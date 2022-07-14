package classfile

// method or field name and descript
// field
// |field desc| field type| method desc| method    |
// | S		  | short     | ()V		   | void run()|
type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (cni *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	cni.nameIndex = reader.readUint16()
	cni.descriptorIndex = reader.readUint16()
}
