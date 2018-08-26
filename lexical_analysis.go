package bingo_tpl

import (
	"fmt"
	"strconv"
)

const textNode = 1    // 文本节点
const lexicalNode = 2 // 词法节点

// 词法分析

// 词法分析链包含大量节点
type LexicalChain struct {
	Nodes   []*LexicalNode
	current int // 当前指针
}

type LexicalNode struct {
	T       int    // 类型（词法节点还是文本节点）
	Content []byte // 内容，生成模版的时候要使用内容进行输出
}

// 打印节点值
func (n *LexicalNode) Print() {
	// 打印当前指针
	//fmt.Print("[Current Index])

	switch n.T {
	case textNode:
		fmt.Println("[node type]: " + "TEXT") // 文本节点
	case lexicalNode:
		fmt.Println("[node type]: " + "LEXICAL") // 词法节点
	default:
		fmt.Println("[node type]: " + "UNKNOWN TYPE") // 未知类型
		break
	}

	fmt.Println("[content]: " + string(n.Content))
}

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
