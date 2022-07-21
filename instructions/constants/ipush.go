package constants

import (
	"github.com/vzjxif/go-jvm/instructions/base"
	rtda "github.com/vzjxif/go-jvm/runtimedataarea"
)

type BIPUSH struct {
	val int8
}
type SIPUSH struct {
	val int16
}

func (bi *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	bi.val = reader.ReadInt8()
}

func (bi *BIPUSH) Execute(frame *rtda.Frame) {
	i := int32(bi.val)
	frame.OperandStack.PushInt(i)
}
