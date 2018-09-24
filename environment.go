package bingo_tpl

import (
	"os"
	"io"
)

//
//import (
//	"os"
//	"io"
//	"github.com/pkg/errors"
//	"reflect"
//	"strconv"
//)
//
//type Environment struct {
//	//RootPath string  // 根目录
//	loader         *Loader                  // 加载器
//	Ext            string                   // 文件后缀
//	LDelimiter     string                   // 左定界符
//	RDelimiter     string                   // 右定界符
//	BufferFileSize int                      // 每次读取文件的长度
//	Tpl            map[string]*LexicalChain // 模版map，模版名->词法链
//	TokenStream    map[string]*TokenStream  // token map, 模版名->token流
//	Operators      map[string]*Operator     // 操作符
//	Tags           map[string]*Tag          // 标签
//}

// 加载器接口
type LoaderInterface interface {
	GetPath() string
}

type Loader struct {
	Path string // 模版根目录
}

func (l *Loader) GetPath() string {
	return l.Path
}

//
//// 初始化环境支持(添加支持的操作符、支持的运算符、 支持的标签、支持的内置函数等等)
//func (e *Environment) Init() {
//	e.Tpl = make(map[string]*LexicalChain)
//	e.BufferFileSize = 100
//	e.LDelimiter = "{{"
//	e.RDelimiter = "}}"
//
//	// 注册操作符，注册基本函数
//	envAddOperators(e)
//}
//
//const (
//	PrecedenceHigh   = 3 // 高
//	ProcedenceMedium = 2 // 中
//	ProcedenceLow    = 1 // 低
//

//)
//
//// 向环境中添加操作符
//func envAddOperators(env *Environment) {
//	// + - * / ^
//	env.AddOperator("+", func(left interface{}, right interface{}) (interface{}, error) {
//		// 如果是+ ，判定左侧是字符串还是数字
//		leftType := reflect.TypeOf(left).String()
//		rightType := reflect.TypeOf(right).String()
//		if leftType == "string" || rightType == "string" {
//			if leftType == "int" {
//				left = strconv.Itoa(left.(int))
//			}
//			if rightType == "int" {
//				right = strconv.Itoa(right.(int))
//			}
//
//			return left.(string) + right.(string), nil
//		}
//
//		// 如果都是int的话，执行相加操作
//		if leftType == "int" && rightType == "int" {
//			return left.(int) + right.(int), nil
//		}
//
//		return nil, errors.New("unknown type with operator '+'")
//	}, ProcedenceLow, AssociativityLeft)
//
//	env.AddOperator("-", func(left interface{}, right interface{}) (interface{}, error) {
//		leftType := reflect.TypeOf(left).String()
//		rightType := reflect.TypeOf(right).String()
//		if leftType == "int" && rightType == "int" {
//			return left.(int) - right.(int), nil
//		}
//		return nil, errors.New("wrong args type with operator '-'")
//	}, ProcedenceLow, AssociativityLeft)
//
//	env.AddOperator("*", func(left interface{}, right interface{}) (interface{}, error) {
//		leftType := reflect.TypeOf(left).String()
//		rightType := reflect.TypeOf(right).String()
//		if leftType == "int" && rightType == "int" {
//			return left.(int) * right.(int), nil
//		}
//		return nil, errors.New("wrong args type with operator '*'")
//	}, ProcedenceMedium, AssociativityLeft)
//
//	env.AddOperator("/", func(left interface{}, right interface{}) (interface{}, error) {
//		leftType := reflect.TypeOf(left).String()
//		rightType := reflect.TypeOf(right).String()
//		if leftType == "int" && rightType == "int" {
//			return left.(int) / right.(int), nil
//		}
//		return nil, errors.New("wrong args type with operator '/'")
//	}, ProcedenceMedium, AssociativityLeft)
//
//	env.AddOperator("^", func(left interface{}, right interface{}) (interface{}, error) {
//		leftType := reflect.TypeOf(left).String()
//		rightType := reflect.TypeOf(right).String()
//		if leftType == "int" && rightType == "int" {
//			return left.(int) ^ right.(int), nil
//		}
//		return nil, errors.New("wrong args type with operator '^'")
//	}, PrecedenceHigh, AssociativityRight)
//
//}
//
//// 操作符
//type Operator struct {
//	Name          string                                                         // 名称
//	callback      func(left interface{}, right interface{}) (interface{}, error) // 回调
//	precedence    int                                                            // 优先级
//	associativity int                                                            // 结合性
//}
//
//// 为环境附加操作符
//// name:  操作符的值
//// callback: 遇到该操作符所执行的回调
//// precedence 优先级
//// associativity 结合性
//// return 返回一个随意类型的值
//func (e *Environment) AddOperator(name string, callback func(left interface{}, right interface{}) (interface{}, error), precedence int, associativity int) bool {
//	operator := new(Operator)
//	operator.Name = name
//	operator.callback = callback
//	operator.precedence = precedence
//	operator.associativity = associativity
//	e.Operators[name] = operator
//	return true
//}
//
//// 创建一个环境对象
func NewEnv(loader *Loader, options map[string]string) *Environment {
	env := &Environment{loader: loader}
	env.Init(loader, options)
	return env
}

//
//// 将词法链转换为token流，返回一个字符串
//// 传入模版路径，变量map
//// 遍历词法链，如果是文本节点，直接压入结果
//func (e *Environment) Render(path string, params map[string]interface{}) (string, error) {
//	//var strem *TokenStream
//	var res []byte
//	if v, ok := e.Tpl[path]; ok {
//		v.Iterator(func(node *LexicalNode) {
//			if node.T == textNode {
//				// 如果是文本节点，就压入结果字符串中
//				for _, char := range node.Content {
//					res = append(res, char)
//				}
//			}
//
//			if node.T == lexicalNode {
//				// 解析抽象语法树，得到最终结果，默认是空字符串，每一个词法节点都有一棵抽象语法树
//
//			}
//
//		})
//
//	} else {
//		return "", errors.New("no such template path exist in tpl map :" + path)
//	}
//}
//
// 输入文件路径，转换为词法链
// 按照固定字节读取并遍历文件
// 使用类似词法分析的方式区分词法节点和文本节点
// 从文件中解析词法链
func (e *Environment) OpenLexicalChain(filePath string) *LexicalChain {
	// 按固定字节读取文件
	// 打开文件流
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buffer := make([]byte, e.BufferFileSize)

	chain := new(LexicalChain)

	tmp := new(Template)
	tmp.chain = chain

	//e.Tpl[filePath] = chain
	e.loadedTemplates[filePath] = tmp

	//e.Tpl[filePath] = chain
	// 初始化为0态
	// 遇到分隔符第一位转为1态，下一个字符如果是分隔符第二位转换为2态，以此类推
	// 当转换态等于左分隔符长度的时候，封闭节点，这个节点就是文本节点
	// 当转换态等于右分隔符长度的时候，封闭节点，这个节点就是词法节点
	stats := 0
	var tmpStats int
	var delimiter string
	lineNum := 0
	//var tmpSlice []byte
	var n = new(LexicalNode)
	n.T = textNode      // 先挂上一个初始节点
	n.root = new(Token) // 挂上语法树根节点
	n.lineNum = lineNum
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

			// 遇到换行符
			if value == 10 {
				lineNum++
			}

			// 先把字符压入临时切片中
			n.Content = append(n.Content, value)
			// 根据字符，转换字符状态
			// 字符态应该是上一次遇见的左分隔符的对应索引+1 ，为0时是没遇见到
			// 得到目标索引
			delimiter = e.LDelimiter
			tmpStats = stats
			if stats >= len(e.LDelimiter) {
				delimiter = e.RDelimiter // 当前字符态对应的分隔符

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
					n.lineNum = lineNum
					chain.Nodes = append(chain.Nodes, n) // 放入词法链中
					// 左节点，不许控制转换态回转
				}
				if stats == len(e.LDelimiter)+len(e.RDelimiter) { // 右节点
					n.T = lexicalNode                                     // 更改节点类型
					n.Content = n.Content[:len(n.Content)-len(delimiter)] // 塞入内容，塞入之前要弹出右分隔符
					n.lineNum = lineNum
					// 将content解析称token流，然后解析token流成抽象语法树
					//n.stream = parseContentToTokenStream(n.Content, e)

					// 解析节点，生成抽象语法树 传入根节点，和要解析的文本，e中包括当前可使用的操作符、标签、变量等数据
					parseToken(n.root, n.Content, e)

					// 创建一个新节点
					n = new(LexicalNode)
					n.T = textNode
					n.lineNum = lineNum
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

	return chain
}

const version = "1.0"

type Environment struct {
	//charset             string                 // 编码
	loader          LoaderInterface      // 加载器
	debug           bool                 // 是否是调试模式
	lexer           *Lexer               // 词法分析器
	parser          ParserInterface      // 语法解析器
	compiler        CompilerInterface    // 编译器（最后进行变量映射）
	templateExt     string               // 模版文件名后缀
	loadedTemplates map[string]*Template // 已经加载好了的模版
	extensionSet    *ExtensionSet        // 扩展集合
	LDelimiter      string               // 左定界符
	RDelimiter      string               // 右定界符
	BufferFileSize  int                  // 每次读取的文件流大小
}

// 初始化结构体  相当于构造函数
func (e *Environment) Init(loader LoaderInterface, options map[string]string) {

	// 设置加载器
	e.setLoader(loader)

	// 初始化配置
	if v, ok := options["ext"]; ok {
		e.templateExt = v
	} else {
		e.templateExt = ".html"
	}

	if v, ok := options["leftDelimiter"]; ok {
		e.LDelimiter = v
	} else {
		e.LDelimiter = "{{"
	}

	if v, ok := options["rightDelimiter"]; ok {
		e.RDelimiter = v
	} else {
		e.RDelimiter = "}}"
	}

	// 初始化模版map
	e.loadedTemplates = make(map[string]*Template)
	// 是否是调试模式
	// 设置字符集
	// 设置是否自动重载
	// 是否以严格模式运行
	// 是否缓存

	e.BufferFileSize = 100

	// 初始化扩展集合
	e.extensionSet = new(ExtensionSet)

	// 添加扩展
	//e.AddExtension(new(extensions.Ext_Core))
}

// 向环境中添加扩展
func (e *Environment) AddExtension(ext ExtensionInterface) {
	e.extensionSet.AddExtension(ext)

	// 将扩展中的数据挂载到环境中


	// 此处更新了optionsHash值
}


func (e *Environment) setLoader(loader LoaderInterface) {
	e.loader = loader
}

// 将模版解析成token流
func (e *Environment) tokenize(chain *LexicalChain) *TokenStream {
	if e.lexer == nil {
		e.lexer = new(Lexer)
		e.lexer.LDelimiter = e.LDelimiter
		e.lexer.RDelimiter = e.RDelimiter

		// 单纯词法分析，将文本解析成token流，不涉及到语法分析，所以无需传入扩展集合
	}
	chain.Iterator(func(node *LexicalNode) {
		node.Tokenize(e.lexer)
	})

	//for _, v := range chain.Nodes {
	//	v.Tokenize(e.lexer)
	//}

	return nil
	//return e.lexer.tokenize(source)
}

// 加载模版，应当将模版解析成token流，并初步解析成语法树
func (e *Environment) LoadTemplate(tmp string) *Template {
	path := e.loader.GetPath() + "/" + tmp + e.templateExt
	// 判断文件是否存在或者是否已经加载
	if v, ok := e.loadedTemplates[path]; ok {
		return v
	}
	// 不存在，去查找文件
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		panic("no such file :  " + path)
	}

	// 解析成词法链
	chain := e.OpenLexicalChain(path)

	// 将词法链解析成token流
	e.tokenize(chain)

	//for _, v := range e.lexer.tokens {
	//	fmt.Println(string(v.Value))
	//	//fmt.Println(v.T)
	//}

	return e.loadedTemplates[path]
}
