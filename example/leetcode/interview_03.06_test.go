//动物收容所。有家动物收容所只收容狗与猫，且严格遵守“先进先出”的原则。在收养该收容所的动物时，收养人只能收养所有动物中“最老”（由其进入收容所的时间长短而定
//）的动物，或者可以挑选猫或狗（同时必须收养此类动物中“最老”的）。换言之，收养人不能自由挑选想收养的对象。请创建适用于这个系统的数据结构，实现各种操作方法，比如
//enqueue、dequeueAny、dequeueDog和dequeueCat。允许使用Java内置的LinkedList数据结构。
//
// enqueue方法有一个animal参数，animal[0]代表动物编号，animal[1]代表动物种类，其中 0 代表猫，1 代表狗。
//
// dequeue*方法返回一个列表[动物编号, 动物种类]，若没有可以收养的动物，则返回[-1,-1]。
//
// 示例1:
//
//  输入：
//["AnimalShelf", "enqueue", "enqueue", "dequeueCat", "dequeueDog", "dequeueAny"
//]
//[[], [[0, 0]], [[1, 0]], [], [], []]
// 输出：
//[null,null,null,[0,0],[-1,-1],[1,0]]
//
//
// 示例2:
//
//  输入：
//["AnimalShelf", "enqueue", "enqueue", "enqueue", "dequeueDog", "dequeueCat", "
//dequeueAny"]
//[[], [[0, 0]], [[1, 0]], [[2, 1]], [], [], []]
// 输出：
//[null,null,null,null,[2,1],[0,0],[1,0]]
//
//
// 说明:
//
//
// 收纳所的最大容量为20000
//
// Related Topics 设计
// 👍 24 👎 0

package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnimalShelf(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		s := ConstructorAnimalShelf()
		s.Enqueue([]int{0, 0})
		s.Enqueue([]int{1, 0})
		assert.Equal(t, defaultValue, s.DequeueDog())
		assert.Equal(t, []int{0, 0}, s.DequeueCat())
		assert.Equal(t, []int{1, 0}, s.DequeueAny())
	})

	t.Run("case 2", func(t *testing.T) {
		s := ConstructorAnimalShelf()
		s.Enqueue([]int{0, 0})
		s.Enqueue([]int{1, 0})
		s.Enqueue([]int{2, 1})
		assert.Equal(t, []int{2, 1}, s.DequeueDog())
		assert.Equal(t, []int{0, 0}, s.DequeueCat())
		assert.Equal(t, []int{1, 0}, s.DequeueAny())
	})

}

var defaultValue = []int{-1, -1}

type Animal = []int
type AnimalShelf struct {
	cat []Animal
	dog []Animal
}

func ConstructorAnimalShelf() AnimalShelf {
	return AnimalShelf{}
}

func (this *AnimalShelf) Enqueue(animal []int) {
	if animal[1] == 0 {
		this.cat = append(this.cat, animal)
	} else {
		this.dog = append(this.dog, animal)
	}
}

func (this *AnimalShelf) DequeueAny() []int {
	if len(this.cat) > 0 && len(this.dog) > 0 {
		cat := this.cat[0]
		dog := this.dog[0]
		if cat[0] < dog[0] {
			return this.DequeueCat()
		}
		return this.DequeueDog()
	} else {
		if len(this.cat) > 0 {
			return this.DequeueCat()
		} else if len(this.dog) > 0 {
			return this.DequeueDog()
		} else {
			return defaultValue
		}
	}
}

func (this *AnimalShelf) DequeueDog() []int {
	if len(this.dog) == 0 {
		return defaultValue
	}
	val := this.dog[0]
	this.dog = this.dog[1:]
	return val
}

func (this *AnimalShelf) DequeueCat() []int {
	if len(this.cat) == 0 {
		return defaultValue
	}
	val := this.cat[0]
	this.cat = this.cat[1:]
	return val
}
