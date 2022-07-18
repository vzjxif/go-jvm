package runtimedataarea

type Thread struct {
	PC    int
	Stack *Stack
}

func NewThread() *Thread {
	return &Thread{
		Stack: newStask(1024),
	}
}

func (t *Thread) PushFrame(frame *Frame) {
	t.Stack.push(frame)
}

func (t *Thread) PopFrame() *Frame {
	return t.Stack.pop()
}
func (t *Thread) CurrentFrame() *Frame {
	return t.Stack.top()
}
