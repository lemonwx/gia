/**
 *  author: lim
 *  data  : 18-8-7 下午10:10
 */

package main

import (
	"github.com/lemonwx/gia/tstruct/minStack"
	"github.com/lemonwx/log"
	"os"
)

func main() {

	log.NewDefaultLogger(os.Stdout)
	log.SetLevel(log.DEBUG)

	s := minStack.Constructor()

	log.Debug(s)
	s.Push(123)
	s.Push(-1)
	log.Debug(s)
	s.Pop()
	log.Debug(s)
}
