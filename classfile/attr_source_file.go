package classfile

// inner classfile
// attribute_lenght must be 2
// fileIndex pointer to cp utf8
//SourceFile_attribute {    u2 attribute_name_index;    u4 attribute_length;    u2 sourcefile_index; }
type SourceFileAttribute struct {
	cp              ConstantPool
	sourceFileIndex uint16
}

func (sfa *SourceFileAttribute) readInfo(reader *ClassReader) {
	sfa.sourceFileIndex = reader.readUint16()
}

func (sfa *SourceFileAttribute) FileName() string {
	return sfa.cp.getUtf8(sfa.sourceFileIndex)
}
