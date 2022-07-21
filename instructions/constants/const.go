package constants

import (
	"github.com/vzjxif/go-jvm/instructions/base"
	rtda "github.com/vzjxif/go-jvm/runtimedataarea"
)

// push null to op stack top
type ACONST_NULL struct{ base.NoOperandsInstruction }

func (an *ACONST_NULL) Execute(frame *rtda.Frame) {
	frame.OperandStack.PushRef(nil)
}

// push double 0
type DCONST_0 struct{ base.NoOperandsInstruction }

func (d0 *DCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack.PushDouble(0.0)
}

type DCONST_1 struct{ base.NoOperandsInstruction }

type FCONST_0 struct{ base.NoOperandsInstruction }

type FCONST_1 struct{ base.NoOperandsInstruction }

type FCONST_2 struct{ base.NoOperandsInstruction }

// push -1
type ICONST_M1 struct{ base.NoOperandsInstruction }

func (im *ICONST_M1) Execute(frame *rtda.Frame) {
	frame.OperandStack.PushInt(0.0)
}

type ICONST_0 struct{ base.NoOperandsInstruction }
type ICONST_1 struct{ base.NoOperandsInstruction }
type ICONST_2 struct{ base.NoOperandsInstruction }
type ICONST_3 struct{ base.NoOperandsInstruction }
type ICONST_4 struct{ base.NoOperandsInstruction }
type ICONST_5 struct{ base.NoOperandsInstruction }

type LCONST_0 struct{ base.NoOperandsInstruction }
type LCONST_1 struct{ base.NoOperandsInstruction }
