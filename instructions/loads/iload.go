package loads

import (
	"github.com/vzjxif/go-jvm/instructions/base"
	rtda "github.com/vzjxif/go-jvm/runtimedataarea"
)

type ILOAD struct{ base.Index8Instruction }
type ILOAD_0 struct{ base.NoOperandsInstruction }
type ILOAD_1 struct{ base.NoOperandsInstruction }
type ILOAD_2 struct{ base.NoOperandsInstruction }
type ILOAD_3 struct{ base.NoOperandsInstruction }

func _iload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars.GetInt(index)
	frame.OperandStack.PushInt(val)
}

func (il *ILOAD) Execute(frame *rtda.Frame) {
	_iload(frame, il.Index)
}

func (il *ILOAD_1) Execute(frame *rtda.Frame) {
	_iload(frame, 1)
}
