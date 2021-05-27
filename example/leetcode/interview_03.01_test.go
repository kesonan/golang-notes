//三合一。描述如何只用一个数组来实现三个栈。
//
// 你应该实现push(stackNum, value)、pop(stackNum)、isEmpty(stackNum)、peek(stackNum)方法。s
//tackNum表示栈下标，value表示压入的值。
//
// 构造函数会传入一个stackSize参数，代表每个栈的大小。
//
// 示例1:
//
//  输入：
//["TripleInOne", "push", "push", "pop", "pop", "pop", "isEmpty"]
//[[1], [0, 1], [0, 2], [0], [0], [0], [0]]
// 输出：
//[null, null, null, 1, -1, -1, true]
//说明：当栈为空时`pop, peek`返回-1，当栈满时`push`不压入元素。
//
//
// 示例2:
//
//  输入：
//["TripleInOne", "push", "push", "push", "pop", "pop", "pop", "peek"]
//[[2], [0, 1], [0, 2], [0, 3], [0], [0], [0], [0]]
// 输出：
//[null, null, null, null, 2, 1, -1, -1]
//
// Related Topics 设计
// 👍 34 👎 0

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
