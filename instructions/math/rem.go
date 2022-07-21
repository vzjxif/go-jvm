package math

import (
	"math"

	"github.com/vzjxif/go-jvm/instructions/base"
	rtda "github.com/vzjxif/go-jvm/runtimedataarea"
)

type DREM struct{ base.NoOperandsInstruction }

type FREM struct{ base.NoOperandsInstruction }

type IREM struct{ base.NoOperandsInstruction }

type LREM struct{ base.NoOperandsInstruction }

func (irem *IREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v2 == 0 {
		panic("java.lang.ArithmeticException:/ by zero")
	}
	result := v1 % v2
	stack.PushInt(result)
}

func (drem *DREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	resutl := math.Mod(v1, v2)
	stack.PushDouble(resutl)
}
