package base

import rtda "github.com/vzjxif/go-jvm/runtimedataarea"

type Instruction interface {
	// 从字节码提取操作数
	FetchOperands(reader *BytecodeReader)
	Execute(frame *rtda.Frame)
}

type NoOperandsInstruction struct {
}

func (nop *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	// do nothig
}

// jump op
type BranchInstruction struct {
	Offset int
}

func (bi *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	bi.Offset = int(reader.ReadInt16())
}

type Index8Instruction struct {
	Index uint
}

func (ii *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	ii.Index = uint(reader.ReadUint8())
}

type Index16Instruction struct {
	Index uint
}

func (ii *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	ii.Index = uint(reader.ReadUint16())
}
