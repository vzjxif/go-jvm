package stack

import (
	"github.com/vzjxif/go-jvm/instructions/base"
	rtda "github.com/vzjxif/go-jvm/runtimedataarea"
)

type DUP struct{ base.NoOperandsInstruction }

type DUP_X1 struct{ base.NoOperandsInstruction }

type DUP_X2 struct{ base.NoOperandsInstruction }

type DUP2 struct{ base.NoOperandsInstruction }
type DUP2_X1 struct{ base.NoOperandsInstruction }
type DUP2_X2 struct{ base.NoOperandsInstruction }

// copy top of stack

func (dup *DUP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	slot := stack.PopSlot()
	stack.PushSlot(slot)
	stack.PushSlot(slot)
}

// todo read source code
