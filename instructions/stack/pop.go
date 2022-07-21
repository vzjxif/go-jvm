package stack

import (
	"github.com/vzjxif/go-jvm/instructions/base"
	rtda "github.com/vzjxif/go-jvm/runtimedataarea"
)

type POP struct{ base.NoOperandsInstruction }

type POP2 struct{ base.NoOperandsInstruction }

// pop top of stack
// only pop int ...
func (pop *POP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	stack.PopSlot()
}

func (pop *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	stack.PopSlot()
	stack.PopSlot()
}
