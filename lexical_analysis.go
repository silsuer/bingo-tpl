package bingo_tpl

import (
	"fmt"
	"strconv"
	"bytes"
	"sync"
)

const textNode = 1    // 文本节点
const lexicalNode = 2 // 词法节点

// 词法分析器接口
type LexerInterface interface {
}

type Lexer struct {
	tokens     []*Token      // token
	chain      *LexicalChain // 词法链
	source     []byte        // 词法分析器中的资源
	cursor     int           // 光标
	lineNum    int           // 当前行数
	stack      []int         // 临时栈，用来记忆if、for等操作
	tmpSlice   []byte        // 临时字符串的栈
	status     int           // 词法分析时转换态状态码
	LDelimiter string        // 左定界符
	RDelimiter string        // 右定界符
	sync.RWMutex
}

func (le *Lexer) tokenize(source []byte) *TokenStream {
	return nil
}

func (le *Lexer) pushToken(t *Token) {
	le.tokens = append(le.tokens, t)
}

// 词法分析链包含大量节点
type LexicalChain struct {
	Nodes       []*LexicalNode
	current     int                    // 当前指针
	Params      map[string]interface{} // 变量名->变量值
	TokenStream *TokenStream           // token流，这是通过节点解析出来的
}

type LexicalNode struct {
	T       int      // 类型（词法节点还是文本节点）
	Content []byte   // 内容，生成模版的时候要使用内容进行输出
	tokens  []*Token // token流
	root    *Token   // 抽象语法树跟节点
	lineNum int      // 行数
	stack   []int    // 符栈，用来记录成对的操作符
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

	fmt.Println("[line number]: " + strconv.Itoa(n.lineNum))
	fmt.Println("[content]: " + string(n.Content))

	// 词法节点，要打印token
	//if n.T == lexicalNode {
	//	n.TokenPrint()
	//}

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

// 将词法链转换为token流,传入一个词法分析器
func (n *LexicalNode) Tokenize(lexer *Lexer) {

	// 只有词法节点才会使用
	if n.T == lexicalNode {

		//fmt.Println(string(n.Content))
		for _, v := range n.Content {

			// 如果是空格...
			//if v == 32 {
			//  lexer.lexSpace()
			//}
			// 如果是运算符... 操作符包括一元运算符和二元运算符

			// 空格 . ' " 操作符 5种

			// 如果是空格
			switch v {
			case 32:
				lexer.lexSpace()
			case 46: // .
				lexer.lexPunctuation(v)
			case 39, 34: // ' "
				lexer.lexQuotation(v)
				//case 43, 45, 42, 47, 94: // + - * / ^
				//	lexer.lexOperator(v)

			default:
				lexer.tmpSlice = append(lexer.tmpSlice, v)
			}

		}
		n.tokens = lexer.tokens[:] // 赋值给节点的token流
		lexer.tokens = []*Token{}  // 清空token
		//for _, vv := range n.tokens {
		//	//fmt.Println(vv.T)
		//	fmt.Println(string(vv.Value))
		//}

	}

}

// 解析空格字符
func (le *Lexer) lexSpace() {
	if len(le.tmpSlice) > 0 { // 临时切片中有数据，将切片中数据作为token保存
		t := new(Token)
		t.Value = le.tmpSlice[:]

		// 检测类型
		if le.lookStackTop() == 0 { // 栈空,证明目前不是在引号内
			_, err := strconv.Atoi(string(le.tmpSlice))
			if err == nil { // 可以强转
				t.T = TypeDigital
			} else {
				t.T = TypeName
			}
		}

		//le.tmpSlice = le.tmpSlice[:0] // 不能使用这个方法清空，此时清空了，但是指针还在，再度添加数据之后，会影响之前的数据
		le.tmpSlice = []byte{}
		le.pushToken(t)
	}
}

// 解析标点 .
func (le *Lexer) lexPunctuation(b byte) {
	if le.lookStackTop() == 0 {
		le.lexSpace() // 将当前的临时切片组成token
		t := new(Token)
		t.T = TypePunctuation
		t.Value = append(t.Value, b)
		le.pushToken(t) // 挂载到token上
	} else {
		le.tmpSlice = append(le.tmpSlice, b)
	}
}

// 解析引号
func (le *Lexer) lexQuotation(b byte) {
	// 左引号

	if le.lookStackTop() == 0 && !le.preIsBackslash() {
		le.pushStack(int(b))
	} else if le.lookStackTop() == int(b) && le.preIsBackslash() {
		// 转码后的 '  ,去掉 \ 并加入引号
		le.tmpSlice = append(le.tmpSlice[:len(le.tmpSlice)-1], b)
	} else if le.lookStackTop() == int(b) && !le.preIsBackslash() {
		t := new(Token)
		t.T = TypeString // 字符串类型
		t.Value = le.tmpSlice[:]
		le.tmpSlice = []byte{}
		le.pushToken(t)
		le.popStack()
	}
}

// 当前上一个字符是否是反斜杠
func (le *Lexer) preIsBackslash() bool {
	if len(le.tmpSlice) > 0 {
		if le.tmpSlice[len(le.tmpSlice)-1] != 92 {
			return false
		} else {
			return true
		}
	} else {
		return false
	}
}

// 解析操作符（加减乘除这种）
func (le *Lexer) lexOperator(b byte) {
	if le.lookStackTop() == 0 {
		// 先把临时栈内的数据拼成token
		t := new(Token)
		t.Value = le.tmpSlice
		_, err := strconv.Atoi(string(le.tmpSlice))
		if err == nil {
			t.T = TypeDigital
		} else {
			t.T = TypeName
		}
		le.pushToken(t)
		le.tmpSlice = []byte{}
		t = new(Token)
		t.Value = append(t.Value, b)
		t.T = TypeOperator

		le.pushToken(t)
	} else {
		le.tmpSlice = append(le.tmpSlice, b) // 如果不是在引号内，则直接压入栈中
	}
}

// 弹出slice顶层数据后的数组
func popSlice(s []byte) []byte {
	if len(s) > 0 {
		return s[0 : len(s)-1]
	}
	return s
}

// 去掉左右的空格
func trimSlice(s []byte) []byte {
	return bytes.Trim(s, " ")
}

// 将传入的数据组成token
func (n *LexicalNode) lexData(s []byte) {
	t := new(Token)
	// 判断能否强转成字符串
	t.T = TypeName
	t.Value = s
	_, err := strconv.Atoi(string(s))
	if err == nil {
		t.T = TypeDigital // 数字token
	}
	n.pushToken(t)
}

// 将数字压入栈中
func (le *Lexer) pushStack(i int) {
	le.stack = append(le.stack, i) // 放入栈
}

// 从栈顶弹出数字
func (le *Lexer) popStack() int {
	if len(le.stack) > 0 {
		res := le.stack[len(le.stack)-1]
		le.stack = le.stack[0 : len(le.stack)-1]
		return res
	} else {
		return 0
	}

}

// 查看站顶数据
func (le *Lexer) lookStackTop() int {
	if len(le.stack) > 0 {
		return le.stack[len(le.stack)-1]
	} else {
		return 0
	}
}

func (n *LexicalNode) pushToken(t *Token) {
	n.tokens = append(n.tokens, t)
}

// 打印节点中的token流
func (n *LexicalNode) TokenPrint() {
	for k, v := range n.tokens {
		fmt.Println("   [token index]: " + strconv.Itoa(k))
		fmt.Println("   [token type]: " + strconv.Itoa(v.T))
		fmt.Println("   [token value]: " + string(v.Value))
		fmt.Println("   ---------------")
	}
}
