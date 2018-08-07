/**
 *  author: lim
 *  data  : 18-8-7 下午9:57
 */

package main

import (
	"github.com/lemonwx/gia/tstruct/queue"
	"github.com/lemonwx/log"
	"os"
)

func main() {
	log.NewDefaultLogger(os.Stdout)
	log.SetLevel(log.DEBUG)

	q := queue.Q{}
	log.Debug(q)

	q.Push(1)
	q.Push(2)

	x := q.Pop()
	log.Debug(x)
	x = q.Pop()
	log.Debug(x)
	log.Debug(q.Pop())

}
