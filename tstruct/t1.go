/**
 *  author: lim
 *  data  : 18-8-7 下午9:26
 */

package main

import (
	"github.com/lemonwx/log"
	"os"
)

type node struct {
}

func (n *node) foo() {
	log.Debugf("this is node's func foo, %v", n)
}

func (n node) foo2() {
	log.Debugf("this is node's func foo2, %v", n)
}

func main() {
	log.NewDefaultLogger(os.Stdout)
	log.SetLevel(log.DEBUG)

	// nil pointer 可以调用指针接受者函数 (可以正常进入函数体内) 但访问成员的变量会panic
	// nil pointer 不可以调用 值接收者函数 如下, n.foo2() 会转化为 (*n).foo2() , 访问 *n 会panic
	var n *node
	n.foo() // this is foo, <nil>
	log.Debug(n, *n)
	n.foo2() // panic
}
