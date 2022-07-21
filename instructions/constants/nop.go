package constants

import (
	rtda "github.com/vzjxif/go-jvm/runtimedataarea"

	"github.com/vzjxif/go-jvm/instructions/base"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (nop *NOP) Execute(frame *rtda.Frame) {
	// do nothing	
}
