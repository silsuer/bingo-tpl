package bingo_tpl

// 语法解析器接口
type ParserInterface interface {
	GetTagName() string // 获取标记名
}

// 将词法链解析为token流

// 跳过文本节点，只解析词法节点，所有词法节点共同构成token流，对于尚未传递进来的变量，暂时略过
func (l *LexicalChain) TokenParse() {

}

// 定义变量类型
const (
	TypeEOS              = -1   // end of sentence 语句结尾
	TypeDigital          = iota // 数字
	TypeString                  // 字符串
	TypeStart                   // 标记开始符（左定界符）
	TypeEnd                     // 标记结束符（右定界符）
	TypeVariable                // 变量
	TypeAssignment              // 赋值
	TypeOperator                // 运算符 + - * / 等
	TypePunctuation             // 标点
	TypeLogic                   // 逻辑操作 （标签 函数等）
	TypeInit                    // 初始化类型
	TypeName                    // 名称类型
	TypeLeftParenthesis         // 左括号
	TypeRightParenthesis        // 右括号
)

// 每一个token代表一个词法节点中的数据，多个token组成表达式，每条语句中间使用 ; 分割，如果一条语句最终的结果是变量，则打印出来
// 否则不打印结果，只做逻辑操作
// 先计算能计算的 对于暂时存在未知变量的结果，保留表达式不做输出。
// 当有http请求的时候，克隆出一个token流，对其进行操作，避免污染全局变量
type Token struct {
	T     int    // token类型  数字、字符串、变量、赋值符、操作符、逻辑操作(标签，函数),语句结束符
	Value []byte // 值,使用的时候根据类型进行强制转换
	Content string  // 转换的字符串
}

type TokenStream struct {
	Tokens  []*Token // token流
	current int      // 当前指针
}

// 将传入的字符串解析成token流
// 就是将变量 操作符等按顺序切分成多个token
func parseContentToTokenStream(content []byte, e *Environment) *TokenStream {
	//ts := new(TokenStream)
	//t := new(Token)
	//ts.Tokens = append(ts.Tokens, t) // 将新建的token挂在stream上
	//
	//// 根据传入的e，判定栈中的数据类型
	//
	//var tmpSlice []byte // 用来存储临时数据
	//// 遍历字符串
	//for _, v := range content {
	//	tmpSlice = append(tmpSlice, v) // 压入栈中
	//	switch v {
	//	case 32: // 空格
	//		// 判断临时栈中的数据是什么
	//		setTokenType(tmpSlice, e, t)
	//	case 34: // 双引号
	//	case 39: // 单引号
	//	case 40: // 左括号
	//	case 41: // 右括号
	//	default:
	//		// 默认把字符压入操作数栈中
	//		break
	//	}
	//}
	return nil
}

func setTokenType(content []byte, e *Environment, t *Token) {
	//name := string(content)

	// 数字 字符串 变量 赋值运算符 运算符
	// 判断是否是数字
	//v, err := strconv.Atoi(name)
	//if err == nil { // 是数字
	//	t.T = TypeDigital
	//	t.Value = v
	//}

	// 判断是否是运算符

}

// 递归解析content
// 遍历content的所有字符
// 直到遇到 空格、引号、括号
// 判断之前的数据是什么格式 —— 变量 标签 函数 表达式
// 以t为父节点，构成
func parseToken(t *Token, content []byte, e *Environment) {
	//var tmpSlice []byte
	//pair := 0  // 当前未遇到成对的操作符

}

// 遇见空格
func parseTokenSpace(s []byte, e *Environment) {

}
