package stack

import (
	"github.com/vzjxif/go-jvm/instructions/base"
	"github.com/vzjxif/go-jvm/runtimedataarea"
)

type SWAP struct{ base.NoOperandsInstruction }

func (swap *SWAP) Execute(frame *runtimedataarea.Frame) {
	stack := frame.OperandStack
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}
