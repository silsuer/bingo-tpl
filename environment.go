package bingo_tpl

import (
	"os"
	"io"
)

type Environment struct {
	//RootPath string  // 根目录
	loader         *Loader                  // 加载器
	Ext            string                   // 文件后缀
	LDelimiter     string                   // 左定界符
	RDelimiter     string                   // 右定界符
	BufferFileSize int                      // 每次读取文件的长度
	Tpl            map[string]*LexicalChain // 模版map，模版名->词法链
}
type Loader struct {
	Path string // 模版根目录
}

// 初始化环境支持(添加支持的操作符、支持的运算符、 支持的标签、支持的内置函数等等)
func (e *Environment) Init() {
	e.Tpl = make(map[string]*LexicalChain)
	e.BufferFileSize = 100
	e.LDelimiter = "{{"
	e.RDelimiter = "}}"
}

// 创建一个环境对象
func NewEnv(loader *Loader) *Environment {
	env := &Environment{loader: loader}
	env.Init()
	return env
}

// 输入文件路径，转换为词法链
// 按照固定字节读取并遍历文件
// 使用类似词法分析的方式区分词法节点和文本节点
// 从文件中解析词法链
func (e *Environment) OpenLexicalChain(filePath string) {
	// 按固定字节读取文件
	//bufferSize := 100

	// 打开文件流
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buffer := make([]byte, e.BufferFileSize)

	chain := new(LexicalChain)

	e.Tpl[filePath] = chain
	// 初始化为0态
	// 遇到分隔符第一位转为1态，下一个字符如果是分隔符第二位转换为2态，以此类推
	// 当转换态等于左分隔符长度的时候，封闭节点，这个节点就是文本节点
	// 当转换态等于右分隔符长度的时候，封闭节点，这个节点就是词法节点
	stats := 0
	var tmpStats int
	var delimiter string
	//var tmpSlice []byte
	var n = new(LexicalNode)
	n.T = textNode // 先挂上一个初始节点
	chain.Nodes = append(chain.Nodes, n)

	for {
		readBytes, err := file.Read(buffer)
		if err != nil {
			if err != io.EOF {
				panic(err)
			}
			break
		}

		// 遍历每一个字符
		for _, value := range buffer[:readBytes] {

			//fmt.Print(string(value))
			// 先把字符压入临时切片中
			//tmpSlice = append(tmpSlice, value)
			n.Content = append(n.Content, value)
			// 根据字符，转换字符状态
			// 字符态应该是上一次遇见的左分隔符的对应索引+1 ，为0时是没遇见到
			// 得到目标索引
			delimiter = e.LDelimiter
			tmpStats = stats
			//dFlag := 0 // 代表现在在找左分隔符
			if stats >= len(e.LDelimiter) {
				delimiter = e.RDelimiter // 当前字符态对应的分隔符
				//dFlag = 1                         // 代表现在在找右分隔符
				tmpStats = stats % len(delimiter) // 当前字符态对应的索引位置
			}

			// 检测到了
			if value == delimiter[tmpStats] {
				stats++ // 变更转换态
				// 分为左分隔节点和右分割节点
				if stats == len(delimiter) { // 左节点  将当前的节点封闭，并创建一个新节点挂上

					// 老节点到此为止
					n.Content = n.Content[:len(n.Content)-len(delimiter)]
					// 创建一个节点
					n = new(LexicalNode)
					n.T = textNode
					chain.Nodes = append(chain.Nodes, n) // 放入词法链中
					// 左节点，不许控制转换态回转
				}
				if stats == len(e.LDelimiter)+len(e.RDelimiter) { // 右节点
					n.T = lexicalNode                                     // 更改节点类型
					n.Content = n.Content[:len(n.Content)-len(delimiter)] // 塞入内容，塞入之前要弹出右分隔符
					// 创建一个新节点
					n = new(LexicalNode)
					n.T = textNode
					chain.Nodes = append(chain.Nodes, n) // 放入词法链中
					// 因为已经到达闭合节点，需要回转词法状态
					stats = 0
				}
			} else {

				// 如果转换态是0，证明是文本节点，直接跳过
				// 如果大于0，并且小于分隔符长度，则转换态回滚
				if tmpStats > 0 && tmpStats < len(delimiter) {
					stats--
				}
			}

		}
	}
}
