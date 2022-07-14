package classfile

type ConstantUtf8Info struct {
	str string
}

func (cui *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	cui.str = decodeMUTF8(bytes)
}

//todo read sourcecode
func decodeMUTF8(bytes []byte) string {
	return string(bytes)
}
