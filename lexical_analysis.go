package bingo_tpl

import (
	"fmt"
	"strconv"
)

const textNode = 1    // 文本节点
const lexicalNode = 2 // 词法节点



// 词法分析器接口
type LexerInterface interface {

}






// 词法分析链包含大量节点
type LexicalChain struct {
	Nodes       []*LexicalNode
	current     int                    // 当前指针
	Params      map[string]interface{} // 变量名->变量值
	TokenStream *TokenStream           // token流，这是通过节点解析出来的
}

type LexicalNode struct {
	T       int          // 类型（词法节点还是文本节点）
	Content []byte       // 内容，生成模版的时候要使用内容进行输出
	stream  *TokenStream // token流
	root    *Token       // 抽象语法树跟节点
}

// 打印节点值
func (n *LexicalNode) Print() {
	// 打印当前指针

	switch n.T {
	case textNode:
		fmt.Println("[node type]: TEXT") // 文本节点
	case lexicalNode:
		fmt.Println("[node type]: LEXICAL") // 词法节点
	default:
		fmt.Println("[node type]: UNKNOWN TYPE") // 未知类型
		break
	}

	fmt.Println("[content]: " + string(n.Content))
}

// 对节点进行解析
// 遍历节点的字符，先创建2个栈，操作符栈和操作数栈，操作数不断入栈，
// 遇到 ( ,判断前一个字符是否是空格，如果不是，证明是个函数，构建函数token
// 如果是空格或者操作符，证明是提升优先级的操作，正常通过递归下降计算
// 遇到空格进行一次比较，如果不是注册过的标记，则是变量
// 遇到（）
// 遇到 xxx()
// 遇到 + - * / ，正常通过递归下降进行计算
//func (n *LexicalNode) TokenParse(env *Environment, chain *LexicalChain, params map[string]interface{}) {
//
//	for _, v := range n.Content {
//		// 遍历内容，生成token
//	}
//}

func (l *LexicalChain) Print() {
	// 打印当前节点
	//fmt.Println("[current index]: " + string(l.current))

	l.Iterator(func(node *LexicalNode) {
		fmt.Println("====================")
		fmt.Println("[index]: " + strconv.Itoa(l.current))
		node.Print()
	})
}

// 将指针移动到下一个位置，返回当前位置，如果已经到最后一个，则返回 -1
func (l *LexicalChain) Next() int {
	if l.current+1 < len(l.Nodes) {
		l.current++
		return l.current
	} else {
		return -1
	}
}

func (l *LexicalChain) Current() *LexicalNode {
	return l.Nodes[l.current]
}

func (l *LexicalChain) Iterator(call func(node *LexicalNode)) {
	// 对于链中的每个节点，执行传入的方法
	call(l.Current())
	for {
		if l.Next() != -1 {
			call(l.Current())
		} else {
			break
		}
	}
}
