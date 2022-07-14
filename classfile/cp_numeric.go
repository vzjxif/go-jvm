package classfile

import "math"

// java int
type ConstantIntegerInfo struct {
	val int32
}

func (cii *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	cii.val = int32(bytes)
}

type ConstantFloatInfo struct {
	val float32
}

func (cii *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	cii.val = math.Float32frombits(bytes)
}

//CONSTANT_Long_info {
//	u1 tag;
//  u4 high_bytes;
//  u4 low_bytes;
//}
// 8 byte
type ConstantLongInfo struct {
	val int64
}

func (cii *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	cii.val = int64(bytes)
}

//CONSTANT_Double_info {
//	u1 tag;
//  u4 high_bytes;
//  u4 low_bytes;
//}
// 8 byte
type ConstantDoubleInfo struct {
	val float64
}

func (cii *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	cii.val = math.Float64frombits(bytes)
}
