/**
 *  author: lim
 *  data  : 18-8-7 下午10:08
 */

package minStack

type MinStack struct {
	vals []int
	min  int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{[]int{}, 1 << 31}
}

func (this *MinStack) Push(x int) {
	this.vals = append(this.vals, x)
	if x < this.min {
		this.min = x
	}
}

func (this *MinStack) Pop() {
	if len(this.vals) == 0 {
		return
	}

	size := len(this.vals)
	item := this.vals[size-1]
	this.vals = this.vals[:size-1]

	if item != this.min {
		return
	}
	this.min = 1 << 31
	for _, val := range this.vals {
		if val < this.min {
			this.min = val
		}
	}
}

func (this *MinStack) Top() int {
	size := len(this.vals)
	if size == 0 {
		return 0
	}
	return this.vals[size-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
