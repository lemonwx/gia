/**
 *  author: lim
 *  data  : 18-8-7 下午9:55
 */

package queue

type Q []int

func (q *Q) Push(item int) {
	*q = append(*q, item)
}

func (q *Q) Pop() int {
	if q.Empty() {
		return -1
	}
	item := (*q)[0]
	*q = (*q)[1:]
	return item
}

func (q *Q) Empty() bool {
	return len(*q) == 0
}
