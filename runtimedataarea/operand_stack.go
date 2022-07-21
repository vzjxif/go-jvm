package runtimedataarea

import "math"

type OperandStack struct {
	// 记录栈顶
	size  uint
	slots []Slot
}

func newOperandStack(maxStack uint) *OperandStack {
	if maxStack > 0 {
		return &OperandStack{
			slots: make([]Slot, maxStack),
		}
	}
	return nil
}

func (op *OperandStack) PushInt(val int32) {
	op.slots[op.size].num = val
	op.size++
}
func (op *OperandStack) PopInt() int32 {
	op.size--
	return op.slots[op.size].num
}

func (op *OperandStack) PushFloat(val float32) {
	bits := math.Float32bits(val)
	op.slots[op.size].num = int32(bits)
	op.size++
}
func (op *OperandStack) PopFloat() float32 {
	op.size--
	bits := uint32(op.slots[op.size].num)
	return math.Float32frombits(bits)
}

func (op *OperandStack) PushLong(val int64) {
	op.slots[op.size].num = int32(val)
	op.slots[op.size+1].num = int32(val >> 32)
	op.size += 2
}
func (op *OperandStack) PopLong() int64 {
	op.size -= 2
	low := uint32(op.slots[op.size].num)
	high := uint32(op.slots[op.size+1].num)
	return int64(high)<<32 | int64(low)
}

func (op *OperandStack) PushDouble(val float64) {
	bits := math.Float64bits(val)
	op.PushLong(int64(bits))
}
func (op *OperandStack) PopDouble() float64 {
	bits := uint64(op.PopLong())
	return math.Float64frombits(bits)
}

func (op *OperandStack) PushRef(ref *Object) {
	op.slots[op.size].ref = ref
	op.size++
}
func (op *OperandStack) PopRef() *Object {
	op.size--
	ref := op.slots[op.size].ref
	// help gc
	op.slots[op.size].ref = nil
	return ref
}

func (op *OperandStack) PushSlot(slot Slot) {
	op.slots[op.size] = slot
	op.size++
}

func (op *OperandStack) PopSlot() Slot {
	op.size--
	return op.slots[op.size]
}
