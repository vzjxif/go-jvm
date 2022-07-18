package runtimedataarea

import (
	"testing"
)

func TestLocalVars(t *testing.T) {
	frame := NewFrame(100, 100)
	vars := frame.LocalVars
	vars.SetInt(0, 100)
	vars.SetInt(1, -100)
	println(vars.GetInt(0))
	println(vars.GetInt(1))
}
