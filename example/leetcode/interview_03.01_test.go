package leetcode

type TripleInOne struct {
	list      [][]int
	stackSize int
}

func Constructor0301(stackSize int) TripleInOne {
	return TripleInOne{
		list:      make([][]int, 3),
		stackSize: stackSize,
	}
}

func (this *TripleInOne) Push(stackNum int, value int) {
	if len(this.list[stackNum]) == this.stackSize {
		return
	}

	this.list[stackNum] = append(this.list[stackNum], value)
}

func (this *TripleInOne) Pop(stackNum int) int {
	n := len(this.list[stackNum])
	v := this.Peek(stackNum)
	if n > 0 {
		this.list[stackNum] = this.list[stackNum][:n-1]
	}
	return v
}

func (this *TripleInOne) Peek(stackNum int) int {
	n := len(this.list[stackNum])
	if n > 0 {
		return this.list[stackNum][n-1]
	}
	return -1
}

func (this *TripleInOne) IsEmpty(stackNum int) bool {
	return len(this.list[stackNum]) == 0
}
