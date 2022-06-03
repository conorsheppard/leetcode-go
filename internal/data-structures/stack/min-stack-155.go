package stack

type stackNode struct {
	val int
	min int
}

type MinStack struct {
	stackArray []stackNode
}

func Constructor() MinStack {
	return MinStack{}
}

const (
	MaxStackSize = 30000
	//MaxStackSize = 1 << 31 // 2^31 == 2147483648 == 10000000000000000000000000000000 (which means the largest index is MaxStackSize - 1)
)

func (ms *MinStack) Push(val int) {
	if len(ms.stackArray) == MaxStackSize {
		panic("stackoverflow")
	}
	var min = ms.GetMin()
	if val < min {
		min = val
	}
	ms.stackArray = append(ms.stackArray, stackNode{val: val, min: min})
}

func (ms *MinStack) Pop() {
	if len(ms.stackArray) == 0 {
		return
	}
	ms.stackArray = ms.stackArray[:len(ms.stackArray)-1]
}

func (ms *MinStack) Top() int {
	if len(ms.stackArray) == 0 {
		return MaxStackSize
	}
	return ms.stackArray[len(ms.stackArray)-1].val
}

func (ms *MinStack) GetMin() int {
	if len(ms.stackArray) == 0 {
		return MaxStackSize
	}
	return ms.stackArray[len(ms.stackArray)-1].min
}
